# Strings and UTF-8

## What I Learned

A Go string is a read-only sequence of bytes.

```text
string → bytes → UTF-8 encoded text
```

---

## Strings Are Bytes

```go
text := "Go"
```

A string is stored as bytes.

The `len()` function returns the number of **bytes** in a string.

```go
len(text)
```

It does **not always return the number of visible characters**.

---

## UTF-8

UTF-8 is an encoding system that represents characters as bytes.

Simple characters like:

```text
A
```

can use one byte.

Unicode characters and emojis like:

```text
🔥
🇳🇬
```

can use multiple bytes.

---

## String Indexing

When indexing a string:

```go
text[0]
```

Go returns a `byte`.

Example:

```go
text := "Go"
fmt.Println(text[0])
```

This prints the numeric byte value of `G`, not the character itself.

To print the character:

```go
fmt.Printf("%c\n", text[0])
```

---

## Important Mental Model

> A Go string is a sequence of bytes.

> A `rune` represents a Unicode code point.

> `len()` measures bytes.

> `range` iterates through a string as runes.

---

## What I Observed

I tested:

```go
text := "Go 🔥"
```

I printed:

- The string length
- Every byte
- Every rune

The number of bytes was different from the number of visible characters because Unicode characters can require multiple bytes in UTF-8.

---

## Why This Matters

This is important when building:

- Chat applications
- Text editors
- Search systems
- APIs
- Text validation
- User-generated content systems

---

## Senior Engineering Insight

Do not assume:

> `1 character = 1 byte`

When working with Unicode text, byte length and character count can be different.

For example, this:

```go
len(text)
```

is not always correct for counting visible characters.

---

## Questions

- What does `len()` return for a string?
- What does string indexing return?
- Why can one visible character use multiple bytes?
- What does UTF-8 do?
- Why is `range` useful when working with Unicode strings?
- What is the difference between a byte and a rune?
