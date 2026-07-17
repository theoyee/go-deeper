# Variable Scope & Shadowing Challenge

## Challenge

Predict the output of the following Go program **before running it**:

```go
package main

import "fmt"

var name = "Global"

func main() {
	name := "Function"

	if true {
		name := "Block"
		fmt.Println(name)
	}

	fmt.Println(name)
}
```

---

## My Prediction

```text
Block
Function
```

---

## Actual Output

```text
Block
Function
```

---

## What I Learned

Go resolves variables based on their scope.

The `name` variable declared inside the `if` block is the closest variable to the `fmt.Println` call inside that block.

Therefore:

```go
name := "Block"
```

is used and `"Block"` is printed.

After the `if` block ends, the block-scoped variable no longer exists.

The `name` variable declared inside `main` is then used:

```go
name := "Function"
```

This is why `"Function"` is printed.

---

## Mental Model

> The closest variable in the current scope wins.

The scopes in this example are:

```text
Package Scope
└── Global

Function Scope
└── Function

Block Scope
└── Block
```

The inner variable shadows the outer variable.

---

## Important Insight

Variable shadowing does **not** change or overwrite the outer variable.

It creates a new variable in a more specific scope.

---

## Questions

- What is variable shadowing?
- Why is `"Block"` printed first?
- Why does the `name` variable inside the `if` block not affect the `name` variable in `main`?
- What happens to the block-scoped variable after the `if` block ends?
