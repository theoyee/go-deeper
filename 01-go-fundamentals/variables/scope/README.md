# Variable Scope

## What I Learned

Scope determines where a variable can be accessed.

## Types of Scope

- Package scope
- Function scope
- Block scope

## Variable Shadowing

A variable in an inner scope can have the same name as a variable in an outer scope.

The inner variable shadows the outer variable.

## Mental Model

The closer scope wins.

## Questions

- What is package scope?
- What is function scope?
- What is block scope?
- What is variable shadowing?

````id="h4p4a5"

## Challenge 🔥

Create this:

```go
var name = "Global"

func main() {
	name := "Function"

	if true {
		name := "Block"
		fmt.Println(name)
	}

	fmt.Println(name)
}
````
