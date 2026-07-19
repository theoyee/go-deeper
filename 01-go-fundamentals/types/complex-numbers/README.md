# Complex Numbers

## What I Learned

Go supports complex numbers.

A complex number has two parts:

```text
real part + imaginary part
```

Example:

```go
2 + 3i
```

In this example:

```text
2  → real part
3i → imaginary part
```

---

## Complex Types

Go provides two complex number types:

```go
complex64
complex128
```

### `complex64`

Uses `float32` precision for the real and imaginary parts.

### `complex128`

Uses `float64` precision for the real and imaginary parts.

---

## `real()` and `imag()`

Go provides built-in functions for accessing the two parts of a complex number.

```go
real(number)
```

Returns the real part.

```go
imag(number)
```

Returns the imaginary part.

---

## Example

```go
number := 2 + 3i

fmt.Println(real(number))
fmt.Println(imag(number))
```

---

## Mental Model

> A complex number is made up of a real part and an imaginary part.

---

## Why Complex Numbers Matter

Complex numbers are used in areas such as:

- Signal processing
- Electrical engineering
- Physics
- Mathematics
- Scientific computing

---

## Senior Engineering Insight

Complex number types are not commonly used in typical backend APIs.

However, Go provides domain-specific numeric types for specialized engineering and scientific problems.

The correct type should match the problem domain.

---

## What I Practiced

I created complex numbers and:

- Printed them
- Extracted their real parts
- Extracted their imaginary parts
- Performed arithmetic with complex numbers

---

## Questions

- What are the two parts of a complex number?
- What is the difference between `complex64` and `complex128`?
- What does `real()` do?
- What does `imag()` do?
- Where might complex numbers be useful?
