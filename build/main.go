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

	err := mergeJson(from, to, out)
	if err != nil { panic(err) }
}

func mergeJson(from, to, out string) error {
	log.Printf("Applying properties of %s to %s\n", from, to)

	fromMap, err := unmarshal(from)
	if err != nil { return err }
	toMap, err := unmarshal(to)
	if err != nil { return err }

	outMap := mergeMap(fromMap, toMap)

	err = marshal(out, outMap)
	if err != nil { return err }

	log.Printf("Merged JSON saved to %s\n", out)
	return err
}

func unmarshal(file string) (map[string]interface{}, error) {
	bytes, err := os.ReadFile(file)
	if err != nil { return nil, err }

	v := map[string]interface{}{}
	return v, json.Unmarshal(bytes,&v)
}

func marshal(file string, v any) error {
	outFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil { return err }

	encoder := json.NewEncoder(outFile)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(v)
	if err != nil { return err }

	return outFile.Close()
}

func mergeMap(a, b map[string]interface{}) map[string]interface{} {
	for k, v := range a {
		_, bHasK := b[k]
		mapA, isMap := v.(map[string]interface{})
		sliceA, isSlice := v.([]interface{})

		// if v is JSON primitive, or if b does not have this key, simply assign it to b
		if !bHasK || (!isMap && !isSlice) {
			b[k] = v
		} else if isMap {
			mapB := b[k].(map[string]interface{})
			mergeMap(mapA, mapB)
		} else if isSlice {
			sliceB := b[k].([]interface{})
			b[k] = mergeSlice(sliceA, sliceB)
		}
	}

	return b
}

func mergeSlice(a, b []interface{}) []interface{} {
	keep := []int{}
	for i, v := range a {
		// special case, if v is false and b[i] is not a bool, prune that element from b
		boolA, aIsBool := v.(bool)
		_, bIsBool := b[i].(bool)
		if aIsBool && !boolA && !bIsBool {
			continue
		}
		keep = append(keep, i)

		mapA, isMap := v.(map[string]interface{})
		sliceA, isSlice := v.([]interface{})
		// assume the slice is not mixed-type
		if isMap {
			mapB := b[i].(map[string]interface{})
			mergeMap(mapA, mapB)
		} else if isSlice {
			sliceB := b[i].([]interface{})
			b = mergeSlice(sliceA, sliceB)
		} else {
			// push unique primitives to b
			if !contains(b, v) {
				b = append(b, v)
			}
		}
	}

	// if all elements aren't being kept, prune the slice to only specified indices
	if len(keep) != len(b) {
		b = pruneSlice(b, keep)
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

func pruneSlice(slice []interface{}, keepIndices []int) []interface{} {
	pruned := []interface{}{}
	for i := 0; i < len(keepIndices); i++ {
		pruned = append(pruned, slice[keepIndices[i]])
	}
	return pruned
}
