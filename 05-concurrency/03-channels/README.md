# Channels

## What I Learned

Channels allow goroutines to communicate safely.

Instead of sharing variables directly, one goroutine sends data through a channel while another receives it.

---

## Creating a Channel

```go
ch := make(chan string)
```

This creates a channel that carries `string` values.

Other examples:

```go
make(chan int)
make(chan bool)
make(chan float64)
```

---

## Sending Data

```go
ch <- "Hello"
```

The value is sent into the channel.

---

## Receiving Data

```go
message := <-ch
```

The value is received from the channel.

---

## Blocking Behavior

Channels are blocking by default.

- Sending waits until another goroutine receives.
- Receiving waits until another goroutine sends.

This synchronization happens automatically.

---

## Important Mental Model

Think of a channel as a pipe connecting goroutines.

```text
Goroutine A
      │
      ▼
   Channel
      │
      ▼
Goroutine B
```

Data flows through the channel from one goroutine to another.

---

## Why Channels Matter

Channels help avoid many problems caused by sharing memory between goroutines.

Go encourages communication through channels rather than direct shared state.

> Don't communicate by sharing memory; share memory by communicating.

---

## Senior Engineering Insight

Channels are widely used for:

- Worker pools
- Task queues
- Event processing
- Pipelines
- Concurrent API requests
- Background processing

They are one of Go's defining concurrency features.

---

## What I Practiced

I created a channel, launched a goroutine, sent data through the channel, received the value in the main goroutine, and observed that both goroutines synchronized automatically.

---

## Questions

- What is a channel?
- How do you create a channel?
- How do you send data through a channel?
- How do you receive data from a channel?
- What does it mean that channels are blocking?
- Why are channels preferred over sharing memory directly?
