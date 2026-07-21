# Type Conversion

## What I Learned

Go requires **explicit type conversion** when converting between different types.

Go does not automatically convert values just because the types are related.

---

## Type Conversion Syntax

```go
TargetType(value)
```

Example:

```go
var number int = 42
var decimal float64 = float64(number)
```

The `int` value is explicitly converted to a `float64`.

---

## Important Mental Model

> Go makes type conversions explicit.

This means I must clearly tell Go that I want to convert a value from one type to another.

---

## Conversion Can Lose Data

```go
var decimal float64 = 10.99
number := int(decimal)
```

The result is:

```text
10
```

The decimal part is discarded.

The value is **not rounded**.

---

## Important Insight

Type conversion can change how a value is represented.

Converting between numeric types can cause:

- Data loss
- Precision loss
- Overflow problems

The conversion should be intentional.

---

## Senior Engineering Insight

Go's explicit conversion system helps prevent accidental type conversions and hidden bugs.

The developer must consciously decide to convert a value.

For example:

```go
float64(age)
```

clearly communicates that the `age` value is being converted into a `float64`.

---

## What I Practiced

I converted:

- `int` to `float64`
- `float64` to `int`

I observed that converting a decimal value to an integer discards the decimal portion.

---

## Questions

- What is explicit type conversion?
- What is the syntax for converting a type in Go?
- What happens to the decimal part when converting `10.99` to `int`?
- Why does Go require explicit conversion?
- What problems can happen during type conversion?
