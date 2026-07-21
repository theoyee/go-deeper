# Goroutines

## What I Learned

A goroutine is a lightweight concurrent function in Go.

It allows a function to run independently from the current execution flow.

---

## Basic Syntax

```go
go functionName()
```

The `go` keyword starts a new goroutine.

---

## Example

```go
go say("Hello from goroutine")
say("Hello from main")
```

The first call runs concurrently.

The second call runs in the main goroutine.

---

## Important Mental Model

> A goroutine allows a function to run concurrently with other functions.

The main goroutine continues executing immediately after starting a new goroutine.

---

## Why It Matters

Goroutines are useful for backend systems that need to handle many tasks at the same time, such as:

- HTTP requests
- Database queries
- External API calls
- Background jobs
- File processing
- Email sending

---

## Important Warning

If the `main` function exits, the program stops.

Any running goroutines may be terminated before they finish.

---

## Temporary Waiting

For simple examples, `time.Sleep` can keep the program alive long enough for goroutines to run.

In real applications, a better solution is to use a `WaitGroup`.

---

## What I Practiced

I started a function using the `go` keyword.

I observed that the output order can change because goroutines run concurrently.

I also learned that the main function can exit before a goroutine finishes.

---

## Questions

- What is a goroutine?
- What does the `go` keyword do?
- What is the main goroutine?
- Why can the output order change?
- What happens if `main` exits too early?
- Why is `time.Sleep` only a temporary solution?
