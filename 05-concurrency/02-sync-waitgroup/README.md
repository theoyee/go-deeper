# sync.WaitGroup

## What I Learned

A `sync.WaitGroup` allows the main goroutine to wait until one or more goroutines finish their work.

Without a `WaitGroup`, the program may exit before background goroutines complete.

---

## Basic Workflow

1. Create a `WaitGroup`.
2. Call `Add()` with the number of goroutines.
3. Start the goroutines.
4. Call `Done()` when each goroutine finishes.
5. Call `Wait()` to block until all work is complete.

---

## Example

```go
var wg sync.WaitGroup

wg.Add(2)

go worker("A", &wg)
go worker("B", &wg)

wg.Wait()
```

---

## Important Methods

### `Add(n)`

Increases the number of goroutines to wait for.

```go
wg.Add(2)
```

---

### `Done()`

Signals that one goroutine has finished.

```go
defer wg.Done()
```

This decreases the counter by one.

---

### `Wait()`

Blocks until the counter reaches zero.

```go
wg.Wait()
```

---

## Important Mental Model

Think of a `WaitGroup` as a countdown timer.

```text
Add(3)

3
↓
2
↓
1
↓
0

Wait() unblocks
```

---

## Best Practice

Always write:

```go
defer wg.Done()
```

at the beginning of the goroutine.

This guarantees the counter is decremented even if the function exits early.

---

## Senior Engineering Insight

A `WaitGroup` synchronizes goroutines.

It does **not** transfer data between them.

For communication, Go uses **channels**.

---

## What I Practiced

I launched multiple goroutines, waited for all of them to finish using a `WaitGroup`, and observed that the program only exited after every worker completed.

---

## Questions

- What problem does `sync.WaitGroup` solve?
- What does `Add()` do?
- What does `Done()` do?
- What does `Wait()` do?
- Why is `defer wg.Done()` considered best practice?
- What is the difference between synchronization and communication?
