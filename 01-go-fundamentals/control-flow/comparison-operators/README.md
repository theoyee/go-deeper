# Comparison Operators

## What I Learned

Comparison operators are used to compare values.

The result of a comparison is always a `bool`.

```text id="u6g8c9"
true
false
```

---

## Comparison Operators

| Operator | Meaning                  |
| -------- | ------------------------ |
| `==`     | Equal to                 |
| `!=`     | Not equal to             |
| `>`      | Greater than             |
| `<`      | Less than                |
| `>=`     | Greater than or equal to |
| `<=`     | Less than or equal to    |

---

## Example

```go id="o6x9cq"
a := 10
b := 20

fmt.Println(a < b)
```

The result is:

```text id="2j7k1v"
true
```

Because `10` is less than `20`.

---

## Important Difference: `=` vs `==`

```go id="4j0s1n"
=
```

is the assignment operator.

It assigns a value.

```go id="y3n5s6"
==
```

is the equality comparison operator.

It checks whether two values are equal.

---

## Mental Model

```text id="u8q3z4"
=  → assign
== → compare
```

---

## What I Practiced

I compared integer values using the different comparison operators.

I also used `==` to check whether a password matched an expected value.

---

## Questions

- What does `==` do?
- What does `!=` do?
- What is the difference between `=` and `==`?
- What type does a comparison return?
- What does `>=` mean?
- What does `<=` mean?
