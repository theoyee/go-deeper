# Zero Values

## What I Learned

In Go, every variable has a zero value when it is declared without an initial value.

## Zero Values

| Type    | Zero Value |
| ------- | ---------- |
| string  | `""`       |
| int     | `0`        |
| float64 | `0`        |
| bool    | `false`    |

## Mental Model

Go does not leave variables with random or undefined values.

Every variable starts with a predictable zero value based on its type.

## Why It Matters

Zero values make Go code safer and reduce the need for unnecessary initialization.

## Questions

- What is the zero value of a string?
- What is the zero value of a bool?
- Why does Go use zero values?
