# Iota

## What I Learned

`iota` is a counter used inside constant blocks.

It starts at `0` and increases by `1` for each constant declaration.

## Mental Model

`iota` is useful for representing ordered constant values such as states or categories.

## Example

```go
const (
	Pending = iota
	Processing
	Completed
)
```

---

## Challenge 🔥

Create a **task status system**:

```text
Todo
InProgress
Done
Cancelled
```
