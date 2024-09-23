(method_definition
  name: (property_identifier) @function.method)

(pair
key: (property_identifier) @function.method
value: [(function_expression) (arrow_function)])

(assignment_expression
left: (member_expression
  property: (property_identifier) @function.method)
right: [(function_expression) (arrow_function)])

(call_expression
  function: (identifier) @function.call)

(call_expression
  function: (member_expression
    property: (property_identifier) @function.call))

((identifier) @type
 (#match? @type "^[A-Z]"))

((identifier) @variable.builtin
 (#match? @variable.builtin "^[A-Z]")
 (#is-not? local)
 (property_identifier))

([
    (identifier)
    (shorthand_property_identifier)
    (shorthand_property_identifier_pattern)
 ] @constant
 (#match? @constant "^[A-Z_][A-Z\\d_]+$"))

((identifier) @variable.builtin
 (#match? @variable.builtin "^(module|console|window|document|JSON|Date|history|localStorage)$")
 (#is-not? local))

((identifier) @function.builtin
(#match? @function.builtin "^(fetch|parseInt|parseFloat|URL|atob|bota|encodeURI|encodeURIComponent|decodeURI|decodeURIComponent|isNaN|isPrototypeOf|scrollTo|Boolean|Number|String|Map|Set|Promise)$")
(#is-not? local))

((identifier) @variable.parameter
(#eq? @variable.parameter "arguments")
(#is-not? local))

((identifier) @function
 (#eq? @function "require")
 (#is-not? local))

((identifier) @keyword
(#eq? @keyword "$")
(#is-not? local))

(this) @variable.this

((comment) @comment.doc
 (#match? @comment.doc "@[A-Za-z]+"))

[
  "..."
] @operator

[
  "as"
  "async"
  "await"
  "break"
  "case"
  "catch"
  "continue"
  "debugger"
  "default"
  "delete"
  "do"
  "else"
  "export"
  "extends"
  "finally"
  "for"
  "from"
  "get"
  "if"
  "import"
  "in"
  "instanceof"
  "new"
  "of"
  "return"
  "set"
  "static"
  "switch"
  "target"
  "throw"
  "try"
  "typeof"
  "void"
  "while"
  "with"
  "yield"
] @keyword

[
  "class"
  "const"
  "function"
  "let"
  "var"
] @keyword.subtle
