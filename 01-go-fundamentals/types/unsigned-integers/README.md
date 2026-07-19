# Unsigned Integers

## What I Learned

Unsigned integers are integer types that can only store **zero and positive numbers**.

The `u` in `uint` means **unsigned**.

## Signed vs Unsigned

### Signed integers

Signed integers can store:

```text
negative numbers
zero
positive numbers
```

Examples:

```go
int
int8
int16
int32
int64
```

### Unsigned integers

Unsigned integers can store:

```text
zero
positive numbers
```

Examples:

```go
uint
uint8
uint16
uint32
uint64
```

## Example

```go
var positiveNumber uint = 100
var smallPositive uint8 = 255
```

## Important Mental Model

> Unsigned integers cannot represent negative values.

For example:

```go
var number uint = 0
```

The value cannot become `-1`.

## What I Observed

I tried assigning a negative number to an unsigned integer.

```go
var number uint = -1
```

Go rejected this because `uint` cannot represent negative values.

I also tested subtracting from a `uint` with a value of `0`.

## Senior Engineering Insight

`uint` should not automatically be used just because a value is positive.

For example, a user age or ID could technically be positive, but `uint` is not always the best choice.

Unsigned integers can introduce underflow problems and can make calculations and APIs more complicated.

The type should be chosen based on the **domain and requirements of the data**.

## Questions

- What does the `u` in `uint` mean?
- What is the difference between `int` and `uint`?
- Why can't `uint` store `-1`?
- Why is `uint` not automatically the best type for a user ID?
- What happens when a `uint` with a value of `0` is decremented?
