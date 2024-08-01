package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var from string
	var to string
	var out string
	if len(os.Args) > 2 {
		from = os.Args[1]
		to = os.Args[2]
		out = os.Args[3]
	} else {
		panic("Expected 3 args")
	}

	log.Printf("Applying properties of %s to %s\n", from, to)

	fromBytes, err := os.ReadFile(from)
	if err != nil { panic(err) }

	toBytes, err := os.ReadFile(to)
	if err != nil { panic(err) }

	fromMap := map[string]interface{}{}
	toMap := map[string]interface{}{}

	err = json.Unmarshal(fromBytes,&fromMap)
	if err != nil { panic(err) }
	err = json.Unmarshal(toBytes, &toMap)
	if err != nil { panic(err) }

	merge(fromMap, toMap)

	outFile, err := os.OpenFile(out, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil { panic(err) }

	encoder := json.NewEncoder(outFile)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(toMap)
	if err != nil { panic(err) }

	err = outFile.Close()
	if err != nil { panic(err) }

	log.Printf("Merged JSON saved to %s\n", out)
}

func merge(a, b map[string]interface{}) {
	for k, v := range a {
		// if v is JSON primitive, or if b does not have this key, simply assign it to b
		_, bHasK := b[k]
		mapA, isMap := v.(map[string]interface{})
		sliceA, isSlice := v.([]interface{})

		if !bHasK || (!isMap && !isSlice) {
			b[k] = v
		} else if isMap {
			mapB := b[k].(map[string]interface{})
			merge(mapA, mapB)
		} else {
			sliceB := b[k].([]interface{})
			b[k] = mergeSlice(sliceA, sliceB)
		}
	}
}

func mergeSlice(a, b []interface{}) []interface{} {
	keep := []int{}
	for i, v := range a {
		// special case, if a is false and b is not a bool, prune that element from b
		boolA, aIsBool := v.(bool)
		_, bIsBool := b[i].(bool)
		if aIsBool && !boolA && !bIsBool {
			continue;
		}
		keep = append(keep, i)

		mapA, isMap := v.(map[string]interface{})
		sliceA, isSlice := v.([]interface{})
		// assume the slice is not mixed-type
		if isMap {
			mapB := b[i].(map[string]interface{})
			merge(mapA, mapB)
		} else if isSlice {
			sliceB := b[i].([]interface{})
			b = mergeSlice(sliceA, sliceB)
		} else {
			// add unique primitives to b
			if !contains(b, v) {
				b = append(b, v)
			}
		}
	}

	// if all elements aren't being kept, push the ones we are keeping to a new array
	if len(keep) != len(b) {
		prunedB := []interface{}{}
		for i := 0; i < len(keep); i++ {
			prunedB = append(prunedB, b[keep[i]])
		}
		b = prunedB
	}

	return b
}

func contains(slice []interface{}, value interface{}) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}
	return false
}
