# The Cixac Programming Language

The Cixac programming language — pronounced 'SIGH-zack' — is a passion project that I decided to pursue with inspiration and guidance from Thorsten Ball's [book](https://interpreterbook.com/). Creating this interpreter is the first (out of three) implementation of my object-oriented dynamic programming language.

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

## Table of Contents
+ [Supported Types](#supported-types)
+ [Variable Bindings](#variable-bindings)
+ [Arithmetic Expressions](#arithmetic-expressions)
+ [Single and Multi-Line Comments](#single-and-multi-line-comments)
+ [Conditional Expressions](#conditional-expressions)
+ [For Loop](#for-loop)
+ [While Loop](#while-loop)
+ [Functions and Closures](#functions-and-closures)
+ [Recursion](#recursion)
+ [Strings](#strings)
+ [Arrays](#arrays)
+ [Objects](#objects)
+ [Binary and Unary Operators](#binary-and-unary-operators)
+ [Builtin Functions](#builtin-functions)

## Summary

Cixac is an object-oriented dynamic programming language where the syntax is similar to Go and Python. 

- Variable bindings
- Integers
- Floats
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
| float | ``0.23 9.33 51.22`` |
| string | ``"" "hello"`` |
| null | ``null`` |
| array | ``[] [1, 10] ["food", 49, true, {"foo": "bar"}]`` |
| objects/hashmap | ``{"arr": [1, 2], 5: "five"} `` |

### Variable Bindings

```
let int = 105329                         // integer
let float = 38.221                       // float 
let str = "This is a string"             // string
let nil = null                           // null
let obj = { "x": 0, "y": 0, "z": 0 }     // object/hashmap
let arr = [1, 2, 3, 4, 5, 6, [0, 0, 0]]  // array 
let sub = fn(x, y) { x - y }             // function

// const ensures that the variable's value cannot be changed after its initial assignment
const int = 105329                          // const integer
const fun = fn(a, b) { (a - b) * (a + b) }  // const function
```

### Arithmetic Expressions

```
let x = 5
let y = x * 3
print((x + y) / 2 - 3)
# 7
```

### Single and Multi-Line Comments

```
/*
function adds two variables together

@param - int x
@param - int y
@returns - int
*/
fn add2(x, y) {
	// add x + y
	return x + y
}
```

### Conditional Expressions

```
let a = 18

let greaterThan = fn(x) {
	if (x > 20) {
		return "x is greater"
	} else if (x == 20) {
    return "x equals 20"
  } else {
      "20 is greater" # return keyword is optional
    }
}

print(greaterThan(a))
# 20 is greater
```

### For Loop

```
for (let i = 0; i < 10; i++) {
  if (i % 2 == 0) {
    continue
  }

  print(i)

  if (i >= 5) {
    break
  }
}
# 1
# 3
# 5
```

### While Loop

```
let i = 0

while (i < 5) {
  i += 1
}

print(i)
# 5
```

### Functions and Closures

Using the return keyword is optional when returning an expression.
```
fn multiply(x, y) { 
  x * y 
}
print(multiply(40 / 2, 5))
# 100 

print(fn(x) { x }(5))
# 5

# closure
const newAdder = fn(x) { return fn(y) { x + y } }
const addTwo = newAdder(2)
print(addTwo(3))
# 5

# higher-order function
const sub = fn(x, y) { x - y }
const applyFn = fn(x, y, func) { return func(x, y) }
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
const name = "Joshua"
print(name)
# Joshua

print(name[3] + name[5])
# ha

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

### Objects

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
| ```++ --``` | Increment & Decrement |
| ```+= -= *= /=``` | Compound Assignment |
| ```* / %``` | Multiplication, Division, Modulo |
| ```+ -``` | Addition, Subtraction |
| ```< > <= >=``` | Comparison |
| ```== !=``` | Equality |
| <code>&#124;&#124;</code> | Logical or |
| ```&&``` | Logical and |
| ```!``` | Logical not |

### Builtin Functions

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
