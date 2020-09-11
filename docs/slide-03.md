## Language Design & Syntaxes
* Use `var` when declaring variables, there is also a short way of declaration with `:=`
* Compiler error when something not used
* Function as first-class citizen
* Package based orginization instead of file/class based
* Basic data types such as `string`,`float32`,`float64`,`int`,`int[8|16|32|64]`,`byte`,etc.
* Can perform type alias via keyword `type` (`byte` is example of built-in type alias of `int8`)
* Six kinds of source-code elements which can be declared such as 
  - package imports
  - defined types and type alias
  - named constants
  - variables
  - functions
  - labels 
* Support local-scoped and closure

## Live Code
* Via https://play.golang.org