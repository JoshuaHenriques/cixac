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
$ make run or $ ./bin/cixac
Cixac Version: 0.1-alpha (Aug 20 2024)
Type "quit()" to exit the REPL
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
+ [For In Loop](#for-in-loop)
+ [While Loop](#while-loop)
+ [Functions and Closures](#functions-and-closures)
+ [Recursion](#recursion)
+ [Strings](#strings)
+ [Arrays](#arrays)
+ [Objects](#objects)
+ [Binary and Unary Operators](#binary-and-unary-operators)
+ [Builtin Functions](#builtin-functions)
+ [Array Builtin Functions](#array-builtin-functions)
+ [Object Builtin Functions](#object-builtin-functions)
+ [String Builtin Functions](#string-builtin-functions)

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
// 7
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
      "20 is greater" // return keyword is optional
    }
}

print(greaterThan(a))
// 20 is greater
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
// 1
// 3
// 5
```

### For In Loop

```
for (i, ele in [1, 2, 3]) {
  print("i:ele = " + i + ":" + ele)
}
// i:ele = 0:1
// i:ele = 1:2
// i:ele = 2:3

for (key, val in {0: 1, 1: 2, 2: 3}) {
  print("key:val = " + key + ":" + val)
}
// key:val = 0:1
// key:val = 1:2
// key:val = 2:3

for (i, ch in "string") {
  print("i:ch = " + i + ":" + ch)
}
// i:ch = 0:s
// i:ch = 1:t
// i:ch = 2:r
// i:ch = 3:i
// i:ch = 4:n
// i:ch = 5:g
```

### While Loop

```
let i = 0

while (i < 5) {
  i += 1
}

print(i)
// 5
```

### Functions and Closures

Using the return keyword is optional when returning an expression.
```
fn multiply(x, y) { 
  x * y 
}
print(multiply(40 / 2, 5))
// 100 

print(fn(x) { x }(5))
// 5

// closure
const newAdder = fn(x) { return fn(y) { x + y } }
const addTwo = newAdder(2)
print(addTwo(3))
// 5

// higher-order function
const sub = fn(x, y) { x - y }
const applyFn = fn(x, y, func) { return func(x, y) }
print(applyFn(2, 4, sub))
// -2
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
// 610
```

### Strings

```
const name = "Joshua"
print(name)
// Joshua

print(name[3] + name[5])
// ha

let makeGreeter = fn(greeting) { fn(name) { greeting + " " + name + "!" } }
let hey = makeGreeter("Hey")
print(hey("Joshua"))
// Hey Joshua!
```

### Arrays

```
let arr = ["string", true, 29, fn(x) { x * x }]
print(arr[0])
// string

print(arr[4 - 2])
// 29

print(arr[3](2))
// 4
```

### Objects

```
let obj = {"name": "Alex", "age": 42, "title": "CEO", true: "boolean key", 50: "integer key"}
print(obj["name"])
// Alex

print(obj["age"])
// 42

print(obj[true])
// boolean key

print(obj[50])
// integer key
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

### Array Builtin Functions

| Function | Signature | Description | 
|----------|-----------|-------------| 
| `first` | `ARRAY.first() -> ANY \| NULL` | Returns the first element of the array or NULL if empty | 
| `last` | `ARRAY.last() -> ANY \| NULL` | Returns the last element of the array or NULL if empty | 
| `rest` | `ARRAY.rest() -> ARRAY` | Returns new array with the first element removed | 
| `push` | `ARRAY.push(value: EXPRESSION) -> ARRAY` | Mutates the array by adding the value to the end. Returns the mutated array. |
| `pushleft` | `ARRAY.pushleft(value: EXPRESSION) -> ARRAY` | Mutates the array by adding the value to the beginning. Returns the mutated array. |
| `pop` | `ARRAY.pop() -> ANY` | Mutates the array by removing the last element. Returns the popped value. | 
| `popleft` | `ARRAY.popleft() -> ANY` | Mutates the array by removing the first element. Returns the popped value. | 
| `slice` | `ARRAY.slice(idx1: EXPRESSION, idx2?: EXPRESSION) -> ARRAY` | Returns selected elements in an array as a new array. It selects from a given start, up to a (not inclusive) given end. |
| `contains` | `ARRAY.contains(ele: ANY) -> BOOLEAN` | Return true if the given value is inside the array and false if not. | 
| `index` | `ARRAY.index(ele: ANY) -> INTEGER` | Returns the index of the first element with the specified value and -1 if it's not in the array. | 

### Object Builtin Functions

| Function | Signature | Description | 
|----------|-----------|-------------| 
| `clear` | `HASHMAP.clear() -> VOID` | Clears all the entries of the hashmap. | 
| `keys` | `HASHMAP.keys() -> ARRAY` | Returns an array of all the keys in the hashmap. | 
| `values` | `HASHMAP.values() -> ARRAY` | Returns an array of all the values in the hashmap | 
| `delete` | `HASHMAP.delete(key: HASHABLE) -> VOID` | Deletes the key:value pair at the given key. | 
| `get` | `HASHMAP.get(key: HASHABLE) -> ANY` | Returns the value at the given key. | 
| `set` | `HASHMAP.set(key: HASHABLE, value: ANY) -> VOID` | Sets the given value at the given key. | 
| `contains` | `HASHMAP.contains(key: HASHABLE) -> BOOLEAN` | Returns true if the given key is inside the hashmap and false if not. |

### String Builtin Functions

| Function | Signature | Description | 
|----------|-----------|-------------| 
| `split` | `STRING.split(delim: STRING) -> ARRAY` | Returns an array of the split string at the given delimiter. | 
| `capitalize` | `STRING.capitalize() -> STRING` | Mutates the string by capitalizing the first letter. Returns the string. | 
| `lower` | `STRING.lower() -> STRING` | Mutates the string by making every character lowercase. Returns the string. | 
| `upper` | `STRING.upper() -> STRING` | Mutates the string by making every character uppercase. Returns the string. | 
