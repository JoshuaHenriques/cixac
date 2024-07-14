# The Cixac Programming Language (WIP)

The Cixac programming language is a passion project that I decided to pursue with inspiration and guidance from Thorsten Ball's [book](https://interpreterbook.com/). Creating this interpreter is the first (out of three) implementation of my object-oriented dynamic programming language.

**Note:** The documentation for the Cixac programming language is currently under construction.

## Quickstart with REPL

**Requires**: Go Version >= 1.13

**Make Commands**:

```make test``` - Runs all tests

```make build``` - Builds the project

```make run``` - Builds and runs REPL

**Start REPL**:

```
$ make run
Hello joshua! This is the Cixac programming language!
Type commands
>> 
```

# Documentation

## Summary

Cixac is an object-oriented dynamic programming language where the syntax is similar to Go and Python. 

- Variable bindings
- Integers
- Booleans
- Strings
- Arrays
- Object/Hashmap
- Arithmetic Expressions
- Built-In Functions
- First-Class and Higher-Order Functions 
- Closures

### Supported Types

| Type | Syntax |
| ----- | ----- |
| bool | ``true false`` |
| int | ``0 33 7559`` |
| str | ``"" "hello"`` |
| null | ``null`` |
| array | ``[] [1, 10] ["food", 49, true, {"foo": "bar"}]`` |
| objects/hashmap | ``{"arr": [1, 2], 5: "five"} `` |

### Variable bindings

```
let int = 105329                         // integer
let str = "This is a string"             // string
let nil = null                           // null
let obj = { "x": 0, "y": 0, "z": 0 }     // object/hashmap
let arr = [1, 2, 3, 4, 5, 6, [0, 0, 0]]  // array 
let fun = fn(a, b){ (a - b) * (a + b) }  // function
```

### Arithmetic Expressions

```
let x = 5
let y = x * 3
print((x + y) / 2 - 3)
# 7
```

### Conditional Expressions

```
let a = 18

let greaterThan = fn(x) {
	if (x > 20) {
		return "x is greater"
	} else {
		"20 is greater" # return keyword is optional
	}
}

print(greaterThan(a))
# 20 is greater
```

### Functions and Closures

```
let multiply = fn(x, y) { x * y }
print(multiply(40 / 2, 5))
# 100 

print(fn(x) { x }(5))
# 5

# closure
let newAdder = fn(x) { return fn(y) { x + y } }
let addTwo = newAdder(2)
print(addTwo(3))
# 5

# higher-order function
let sub = fn(x, y) { x - y }
let applyFn = fn(x, y, func) { func(x, y) }
print(applyFn(2, 4, sub))
# -2
```

### Recursion

```
let fib = fn(x) {
	if (x == 0) {
		return 0
	} 
	if (x == 1) {
		return 1
	}
	return fib(x - 1) + fib(x - 2)
}

print(fib(15))
# 610
```

### Strings

```
let makeGreeter = fn(greeting) { fn(name) { greeting + " " + name + "!" } }
let hey = makeGreeter("Hey")
print(hey("Joshua"))
# Hey Joshua!
```

### Arrays

```
let arr = ["string", true, 29, fn(x) { x * x }]
print(arr[0])
# string

print(arr[4 - 2])
# 29

print(arr[3](2))
# 4
```

### Objects/Hashes

```
let obj = {"name": "Alex", "age": 42, "title": "CEO", true: "boolean key", 50: "integer key"}
print(obj["name"])
# Alex

print(obj["age"])
# 42

print(obj[true])
# boolean key

print(obj[50])
# integer key
```

### Binary and Unary Operators

| Operators | Description |
| --------- | ----------- |
| ```[]``` | Subscript |
| ```-``` | Unary minus |
| ```* /``` | Multiplication, Division |
| ```+ -``` | Addition, Subtraction |
| ```< > <= >=``` | Comparison |
| ```== !=``` | Equality |
| || | Logical or |
| && | Logical and |
| ```!``` | Logical not |

### Builtin functions

| Function | Signature | Description | 
|----------|-----------|-------------| 
| `len` | `len(arg: STRING \| ARRAY \| HASH) -> INTEGER` | Returns length of strings, arrays, and hashmaps | 
| `print` | `print(arg: EXPRESSION) -> NULL` | Prints the value(s) to standard output and returns NULL | 
| `first` | `first(arg: ARRAY) -> ANY \| NULL` | Returns the first element of the array or NULL if empty | 
| `last` | `last(arg: ARRAY) -> ANY \| NULL` | Returns the last element of the array or NULL if empty | 
| `rest` | `rest(arg: ARRAY) -> ARRAY` | Returns new array with the first element removed | 
| `push` | `push(arr: ARRAY, value: EXPRESSION) -> ARRAY` | Mutates the array by adding the value to the end. Returns the mutated array. |
| `pushleft` | `pushleft(arr: ARRAY, value: EXPRESSION) -> ARRAY` | Mutates the array by adding the value to the beginning. Returns the mutated array. |
| `pop` | `pop(arr: ARRAY) -> ANY` | Mutates the array by removing the last element. Returns the popped value. | 
| `popleft` | `popleft(arr: ARRAY) -> ANY` | Mutates the array by removing the first element. Returns the popped value. | 
