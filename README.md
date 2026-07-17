<div align="center">

<img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go Logo" width="200"/>

# Go Deeper

### A first-principles laboratory for mastering Go and backend engineering

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](CONTRIBUTING.md)
[![Made with Discipline](https://img.shields.io/badge/Made%20with-Discipline-red?style=for-the-badge)](#learning-principles)

**This is not a tutorial repo. It's a mind, on fire, taking Go apart bolt by bolt.**

</div>

---

## 🧭 Philosophy

> I am not collecting tutorials. I am building mental models.

Most repositories that claim to "learn Go" are graveyards of copy-pasted examples nobody re-derived. This one is different by design.

**Go Deeper** exists to answer one question, over and over, at increasing depth:

<<<<<<< HEAD
> *"Do I actually understand this — or do I just recognize it?"*
=======
> _"Do I actually understand this — or do I just recognize it?"_
>>>>>>> d005337 (learn: variables and scope)

Every concept here is earned: written from scratch, broken on purpose, benchmarked under pressure, and documented — mistakes included — until it becomes intuition instead of memory.

---

## ⚙️ Learning Principles

These aren't guidelines. They're the operating system this repo runs on.

<<<<<<< HEAD
| Principle | What it means in practice |
|---|---|
| 🧠 **Understand before abstracting** | No interface is written until the concrete version is fully understood |
| ✍️ **Write the code from scratch** | Copy-pasting is banned. Typing builds memory; reading doesn't |
| 💥 **Break things intentionally** | Every module ships with a "how I broke it" section |
| 📊 **Benchmark when performance matters** | Claims are backed by `go test -bench`, not vibes |
| 📖 **Read the standard library** | The best Go code ever written is already on your machine |
| 🐛 **Document mistakes** | Bugs and wrong turns are logged, not hidden |
| 🏗️ **Build production-oriented projects** | Toy examples graduate into real, deployable systems |
=======
| Principle                                 | What it means in practice                                              |
| ----------------------------------------- | ---------------------------------------------------------------------- |
| 🧠 **Understand before abstracting**      | No interface is written until the concrete version is fully understood |
| ✍️ **Write the code from scratch**        | Copy-pasting is banned. Typing builds memory; reading doesn't          |
| 💥 **Break things intentionally**         | Every module ships with a "how I broke it" section                     |
| 📊 **Benchmark when performance matters** | Claims are backed by `go test -bench`, not vibes                       |
| 📖 **Read the standard library**          | The best Go code ever written is already on your machine               |
| 🐛 **Document mistakes**                  | Bugs and wrong turns are logged, not hidden                            |
| 🏗️ **Build production-oriented projects** | Toy examples graduate into real, deployable systems                    |
>>>>>>> d005337 (learn: variables and scope)

---

## 🎯 Core Focus

<table>
<tr>
<td width="50%" valign="top">

### Language & Foundations
<<<<<<< HEAD
=======

>>>>>>> d005337 (learn: variables and scope)
- 🔤 Go fundamentals
- 🧩 Data structures
- ⚠️ Error handling
- 🔌 Interfaces & composition
- 🧬 Go internals (runtime, scheduler, GC)

</td>
<td width="50%" valign="top">

### Systems & Scale
<<<<<<< HEAD
=======

>>>>>>> d005337 (learn: variables and scope)
- 🔀 Concurrency (goroutines, channels, sync)
- 🌐 Networking
- 🏢 Backend engineering
- ⚡ Performance & benchmarking

</td>
</tr>
</table>

---

## 📂 Repository Structure

```text
go-deeper/
├── 01-fundamentals/       # syntax, types, memory model
├── 02-data-structures/    # hand-rolled, no shortcuts
├── 03-error-handling/     # idiomatic patterns, wrapped errors
├── 04-interfaces/         # composition over inheritance
├── 05-concurrency/        # goroutines, channels, race conditions
├── 06-networking/         # TCP/HTTP from first principles
├── 07-internals/          # scheduler, GC, escape analysis
├── 08-backend-projects/   # production-oriented builds
└── 09-benchmarks/         # profiling & performance notes
```

Each directory follows the same contract:

```text
├── README.md        → what this concept is, why it matters
├── notes.md          → mental model, mistakes, "aha" moments
├── *.go               → the actual, from-scratch implementation
└── *_test.go          → tests + benchmarks
```

---

## 🚀 Getting Started

```bash
git clone https://github.com/<your-username>/go-deeper.git
cd go-deeper
go run ./01-fundamentals/...
```

Run the benchmarks:

```bash
go test -bench=. -benchmem ./...
```

Run the race detector, because concurrency bugs hide until they don't:

```bash
go test -race ./...
```

---

## 🗺️ Roadmap

- [x] Go fundamentals & memory model
- [x] Data structures from scratch
- [ ] Concurrency deep dive (channels, select, context)
- [ ] Go runtime internals (GMP scheduler)
- [ ] Networking layer from raw TCP up
- [ ] Production backend project #1
- [ ] Continuous benchmarking pipeline

---

## 🤝 Contributing

This is primarily a personal learning lab, but if you're on the same path — going deeper instead of going wide — issues, discussions, and PRs that challenge an assumption or sharpen an explanation are genuinely welcome.

---

## 👤 Author

<div align="center">

**Oyetunji Olagoke** (King Faithful)
Full-stack engineer going deeper into Go, one internal at a time.

[![Portfolio](https://img.shields.io/badge/Portfolio-theoyee.vercel.app-000000?style=for-the-badge&logo=vercel&logoColor=white)](https://theoyee.vercel.app)
[![X](https://img.shields.io/badge/X-@_theoyee-000000?style=for-the-badge&logo=x&logoColor=white)](https://x.com/_theoyee)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Oyetunji%20Olagoke-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/oyetunji-olagoke/)

</div>

---

<div align="center">

### Built with obsession, one `go build` at a time.

<img src="https://go.dev/images/gophers/motorcycle.svg" alt="Gopher" width="120"/>

**⭐ Star this repo if you believe depth beats breadth.**

</div>
