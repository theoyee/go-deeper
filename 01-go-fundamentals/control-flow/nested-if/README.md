# Nested If Statements

## What I Learned

A nested `if` statement is an `if` statement inside another `if` statement.

---

## Example

```go id="j5v9x2"
if age >= 18 {
	if hasID {
		fmt.Println("Access granted")
	}
}
```

The inner `if` is only evaluated if the outer condition is true.

---

## Important Mental Model

> Nested conditions work like multiple gates.

The program must pass the outer condition before reaching the inner condition.

---

## Program Flow

```text id="x8q2m4"
Outer condition
      ↓
Inner condition
      ↓
    Result
```

If the outer condition is false, the inner condition is never evaluated.

---

## Nested If vs Logical Operators

Nested conditions can sometimes be written using logical operators.

Nested version:

```go id="r7k3p9"
if age >= 18 {
	if hasID {
		fmt.Println("Access granted")
	}
}
```

Logical operator version:

```go id="h2m6v8"
if age >= 18 && hasID {
	fmt.Println("Access granted")
}
```

Both can produce the same result.

---

## Senior Engineering Insight

Deeply nested code can become difficult to read and maintain.

Too many nested conditions can be a sign that the code should be refactored.

Later, early returns will help reduce unnecessary nesting.

---

## What I Practiced

I created a nested access check.

The program first checked whether the user passed the outer condition.

Only then did it evaluate the inner condition.

---

## Questions

- What is a nested `if` statement?
- When is the inner `if` evaluated?
- What happens when the outer condition is false?
- How can logical operators replace some nested conditions?
- Why can excessive nesting be a problem?
