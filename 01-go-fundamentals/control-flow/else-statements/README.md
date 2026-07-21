# Else Statements

## What I Learned

The `else` statement tells Go what to do when an `if` condition is `false`.

---

## Basic Syntax

```go id="w7x6s8"
if condition {
	// runs when true
} else {
	// runs when false
}
```

---

## Example

```go id="5x2l5h"
age := 17

if age >= 18 {
	fmt.Println("You can vote")
} else {
	fmt.Println("You cannot vote yet")
}
```

The condition is false because `17` is not greater than or equal to `18`.

Therefore, the `else` block runs.

---

## Important Mental Model

> If the condition is true, execute the `if` block.

> Otherwise, execute the `else` block.

---

## Program Flow

The program evaluates the `if` condition.

Only **one** of the two branches runs.

```text id="l99cb7"
true  → if block
false → else block
```

---

## Important Go Syntax

The `else` statement must be placed on the same line as the closing brace of the `if` block.

Correct:

```go id="oyq7gj"
} else {
```

---

## What I Practiced

I created a login check using an `if` and `else` statement.

I tested both `true` and `false` values and observed that the program executes a different branch depending on the condition.

---

## Questions

- When does the `else` block run?
- Can both the `if` and `else` blocks run?
- What happens if the `if` condition is false?
- Why must `else` be written as `} else {` in Go?
