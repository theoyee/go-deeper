# Zero Values

## What I Learned

Every Go type has a **zero value**.

When a variable is declared without an assigned value, Go automatically gives it the zero value for its type.

---

## Examples

```go id="pxynb5"
var number int
```

The zero value is:

```text id="14t85y"
0
```

```go id="5f8j1s"
var text string
```

The zero value is:

```text id="9j4e8h"
""
```

---

## Common Zero Values

| Type      | Zero Value |
| --------- | ---------- |
| `int`     | `0`        |
| `float64` | `0`        |
| `string`  | `""`       |
| `bool`    | `false`    |
| pointer   | `nil`      |
| slice     | `nil`      |
| map       | `nil`      |

---

## Important Mental Model

> Go never leaves a declared variable with an unpredictable value.

The variable receives the zero value for its type.

---

## Structs and Zero Values

A struct also has a zero value.

```go id="x3jsdf"
type User struct {
	Name string
	Age  int
}
```

When creating:

```go id="a1qyqf"
var user User
```

The fields automatically receive their zero values:

```text id="tr0a3p"
Name: ""
Age: 0
```

---

## Senior Engineering Insight

The zero value is an important part of idiomatic Go.

A well-designed Go type should ideally be useful in its zero state.

> The zero value should be useful.

This makes Go code more predictable and reduces unnecessary initialization.

---

## What I Practiced

I declared variables without assigning values and observed their zero values.

I also created a zero-value struct and observed that each field received the zero value of its own type.

---

## Questions

- What is a zero value?
- What is the zero value of an `int`?
- What is the zero value of a `string`?
- What is the zero value of a `bool`?
- What happens when a struct is declared without initialization?
- Why is the zero value important in Go?
