# Logical Operators

## What I Learned

Logical operators allow multiple conditions to be combined.

Go has three main logical operators:

```text id="h1v4p8"
&&  AND
||  OR
!   NOT
```

---

## `&&` — AND

Both conditions must be `true`.

```go id="u7q2k5"
age >= 18 && hasID
```

If either condition is false, the entire expression is false.

### Mental Model

> This **AND** that must be true.

---

## `||` — OR

At least one condition must be `true`.

```go id="a6r9w3"
isAdmin || isOwner
```

If either condition is true, the entire expression is true.

### Mental Model

> This **OR** that.

---

## `!` — NOT

The `!` operator reverses a boolean value.

```go id="p4n8x2"
!false
```

Becomes:

```text id="z5m7q1"
true
```

### Mental Model

> `!` means the opposite.

---

## Combining Operators

```go id="g2c6v9"
if age >= 18 && hasID && !isBanned {
	fmt.Println("Access granted")
}
```

This means:

> The user must be at least 18, have an ID, and not be banned.

---

## What I Practiced

I used logical operators to combine multiple conditions.

I built a login check that required both a correct username and a correct password.

---

## Questions

- What does `&&` mean?
- What does `||` mean?
- What does `!` do?
- What happens when one condition in an `&&` expression is false?
- What is the difference between `&&` and `||`?
- How can logical operators be used in a login system?
