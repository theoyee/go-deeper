# Byte and Rune

## What I Learned

Go handles characters and text using **bytes and runes**.

They are not the same thing.

---

## `byte`

In Go, `byte` is an alias for `uint8`.

```go
var letter byte = 65
```

A byte stores **8 bits of data**.

`byte` is commonly used for:

- Raw data
- Files
- Network data
- ASCII-oriented data

### Mental Model

> A byte is a small unit of raw data.

---

## `rune`

A `rune` is an alias for `int32`.

```go
var letter rune = 'A'
```

A rune represents a **Unicode code point**.

This allows Go to work with characters from different writing systems and symbols.

Examples:

```go
var character rune = '你'
var emoji rune = '🔥'
```

---

## Important Difference

```text
byte = uint8
rune = int32
```

### `byte`

Used for raw bytes and byte-level data.

### `rune`

Used to represent Unicode characters.

---

## Strings and Bytes

A Go string is a sequence of bytes.

This means:

```go
len(text)
```

returns the **number of bytes**, not necessarily the number of visible characters.

For example:

```go
text := "🔥"
```

The emoji is one visible character, but it takes multiple bytes in UTF-8.

---

## `range` and Runes

When using `range` on a string:

```go
for _, character := range text {
	fmt.Printf("%c\n", character)
}
```

Go iterates through the string as Unicode runes.

---

## Mental Model

> A string is bytes. A rune represents a Unicode character.

`len()` measures bytes.

`range` can decode the string into runes.

---

## What I Observed

I tested:

```go
text := "Go 🔥 Nigeria 🇳🇬"
```

I compared the result of `len(text)` with the number of visible characters.

The values were different because some Unicode characters require multiple bytes to be represented in UTF-8.

---

## Senior Engineering Insight

Assuming that:

> `1 character = 1 byte`

can cause bugs when working with international text.

This matters in:

- Chat applications
- Text editors
- Search systems
- APIs
- User-generated content

---

## Questions

- What is a `byte`?
- What is a `rune`?
- What is the difference between `byte` and `rune`?
- Why does `len()` return bytes instead of visible characters?
- What does `range` do when iterating over a string?
- Why is Unicode important in backend systems?
