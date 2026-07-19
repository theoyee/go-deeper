# If Statements

## What I Learned

An `if` statement allows a Go program to make decisions.

The program evaluates a condition.

```text
true  → execute the code
false → skip the code
```

---

## Basic Syntax

```go
if condition {
	// code
}
```

Example:

```go
age := 22

if age >= 18 {
	fmt.Println("You are an adult")
}
```

The condition is:

```go
age >= 18
```

Go evaluates whether the condition is `true` or `false`.

---

## Important Mental Model

> `if` allows a program to make a decision based on a condition.

---

## Boolean Conditions

An `if` condition must evaluate to a `bool`.

```go
isLoggedIn := true

if isLoggedIn {
	fmt.Println("Welcome")
}
```

Go does not treat numbers as booleans.

For example:

```go
if 1 {
}
```

is invalid in Go.

---

## Important Go Syntax

Go does not require parentheses around an `if` condition.

Correct:

```go
if age >= 18 {
}
```

The opening `{` must also be on the same line as the condition.

---

## What I Practiced

I created an age check that only prints a message when the age is 18 or above.

I tested different ages and observed that the code inside the `if` block is skipped when the condition is false.

---

## Questions

- What does an `if` statement do?
- What happens when an `if` condition is false?
- What type must an `if` condition return?
- Does Go require parentheses around `if` conditions?
- Why must the opening `{` stay on the same line?
