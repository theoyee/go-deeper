# Else If Statements

## What I Learned

`else if` allows a program to check multiple conditions.

---

## Basic Syntax

```go id="c1x3l6"
if conditionA {
	// code
} else if conditionB {
	// code
} else {
	// fallback
}
```

---

## How It Works

Go evaluates conditions from **top to bottom**.

The first condition that evaluates to `true` runs.

After that, Go skips the remaining branches.

---

## Example

```go id="9y7qbc"
score := 75

if score >= 90 {
	fmt.Println("Grade: A")
} else if score >= 80 {
	fmt.Println("Grade: B")
} else if score >= 70 {
	fmt.Println("Grade: C")
} else {
	fmt.Println("Grade: F")
}
```

The program checks:

```text id="6z3g75"
75 >= 90 → false
75 >= 80 → false
75 >= 70 → true
```

The result is:

```text id="o0r4k7"
Grade: C
```

---

## Important Mental Model

> Go evaluates conditions from top to bottom and stops at the first match.

---

## Why Order Matters

The conditions must be ordered carefully.

A broad condition placed before a specific condition can prevent the specific condition from ever running.

---

## What I Practiced

I created a grading system using multiple `else if` conditions.

I learned that only the first matching branch executes.

---

## Questions

- What does `else if` do?
- In what order does Go evaluate conditions?
- What happens after the first condition evaluates to `true`?
- Why does the order of conditions matter?
- What is the purpose of the final `else`?
