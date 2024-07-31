#!/bin/sh
cd "$(dirname "$0")"

TMP=./tmp

BASE=highlights.scm

# download URLs
JS_URL=https://raw.githubusercontent.com/tree-sitter/tree-sitter-javascript/master/queries/
JS_PARAMS=highlights-params.scm
JS_JSX=highlights-jsx.scm

JS_BASE_URL=$JS_URL$BASE
JS_PARAMS_URL=$JS_URL$JS_PARAMS
JS_JSX_URL=$JS_URL$JS_JSX

HTML_URL=https://raw.githubusercontent.com/zed-industries/zed/main/extensions/html/languages/html/
HTML_BASE_URL=$HTML_URL$BASE

XML_URL=https://raw.githubusercontent.com/sweetppro/zed-xml/main/languages/xml/
XML_BASE_URL=$XML_URL$BASE

# override files
CV_JS=highlights-js.scm
CV_HTML=highlights-html.scm
CV_XML=highlights-xml.scm

# fetch the highlight files
mkdir -p $TMP/javascript
mkdir -p $TMP/html
mkdir -p $TMP/xml

wget -q -O $TMP/javascript/$BASE $JS_BASE_URL &
wget -q -O $TMP/javascript/$JS_PARAMS $JS_PARAMS_URL &
wget -q -O $TMP/javascript/$JS_JSX $JS_JSX_URL &
wget -q -O $TMP/html/$BASE $HTML_BASE_URL &
wget -q -O $TMP/xml/$BASE $XML_BASE_URL &

wait

# concatenate files with custom overrides at the bottom of the file
echo >> $TMP/javascript/$BASE && cat $TMP/javascript/$JS_PARAMS >> $TMP/javascript/$BASE
echo >> $TMP/javascript/$BASE && cat $TMP/javascript/$JS_JSX >> $TMP/javascript/$BASE
echo >> $TMP/javascript/$BASE && cat $CV_JS >> $TMP/javascript/$BASE

echo >> $TMP/html/$BASE && cat $CV_HTML >> $TMP/html/$BASE

echo >> $TMP/xml/$BASE && cat $CV_XML >> $TMP/xml/$BASE

# copy final highlights.scm files to languages directory
cp $TMP/javascript/$BASE ../languages/javascript
cp $TMP/html/$BASE ../languages/html
cp $TMP/xml/$BASE ../languages/xml

# cleanup
rm -Rf $TMP

# fetch other scm files
wget -q -O ../languages/javascript/injections.scm $JS_URL/injections.scm &
wget -q -O ../languages/javascript/locals.scm   $JS_URL/locals.scm &
wget -q -O ../languages/javascript/tags.scm     $JS_URL/tags.scm &

wget -q -O ../languages/html/brackets.scm   $HTML_URL/brackets.scm &
wget -q -O ../languages/html/indents.scm    $HTML_URL/indents.scm &
wget -q -O ../languages/html/injections.scm $HTML_URL/injections.scm &
wget -q -O ../languages/html/outline.scm    $HTML_URL/outline.scm &
wget -q -O ../languages/html/overrides.scm  $HTML_URL/overrides.scm &

wget -q -O ../languages/xml/indents.scm $XML_URL/indents.scm &

wait