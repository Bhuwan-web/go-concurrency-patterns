# Go Concurrency Learning Project

A comprehensive exploration of Go's concurrency patterns and mechanisms. This project demonstrates various concurrency concepts through practical examples and experiments.

## 🎯 Project Overview

This repository serves as a hands-on learning resource for understanding Go's concurrency features. Each module focuses on a specific aspect of concurrent programming, with examples that highlight both the power and potential pitfalls of concurrent code.

## 📁 Project Structure

```
learn-concurrency/
├── main.go                 # Entry point and experiment runner
├── race/                   # Race condition demonstrations
│   └── race.go
├── mem_access_sync/        # Memory synchronization examples
│   └── sync.go
└── README.md              # This file
```

## 🚀 Current Features

### 1. Race Conditions (`race/`)

-   **Purpose**: Demonstrates how race conditions occur in concurrent programs
-   **Key Concepts**:
    -   Shared memory access without synchronization
    -   Non-deterministic behavior in concurrent execution
    -   Data races and their unpredictable outcomes

**What it does:**

-   Runs the same function 2000 times to show different results
-   Tracks unique values generated due to race conditions
-   Illustrates why synchronization is crucial

### 2. Memory Access Synchronization (`mem_access_sync/`)

-   **Purpose**: Shows how to properly synchronize access to shared memory
-   **Key Concepts**:
    -   Mutex locks for thread-safe operations
    -   WaitGroups for goroutine coordination
    -   Atomic operations vs non-atomic operations

**What it does:**

-   Demonstrates safe concurrent counter incrementation
-   Uses `sync.Mutex` to prevent race conditions
-   Shows the difference between synchronized and unsynchronized access

## 🏃‍♂️ Running the Examples

```bash
# Clone the repository
git clone <repository-url>
cd learn-concurrency

# Run specific examples by uncommenting them in main.go
go run main.go
```

### Available Examples:

```go
// Uncomment in main.go to run:
race.DisplayRace()              // Shows race condition effects
mem_sync.DisplayMemorySync()    // Demonstrates mutex synchronization
```

## 📚 Learning Outcomes

After exploring this project, you'll understand:

-   ✅ How race conditions occur and why they're problematic
-   ✅ The importance of synchronization in concurrent programs
-   ✅ How to use mutexes to protect shared resources
-   ✅ WaitGroups for coordinating multiple goroutines
-   ✅ The difference between atomic and non-atomic operations

## 🔮 Planned Additions

This project will be incrementally updated with more concurrency patterns:

-   [ ] **Channels**: Communication between goroutines
-   [ ] **Select Statements**: Non-blocking channel operations
-   [ ] **Worker Pools**: Managing concurrent task execution
-   [ ] **Context Package**: Cancellation and timeouts
-   [ ] **Atomic Operations**: Lock-free programming
-   [ ] **Pipeline Patterns**: Data processing pipelines
-   [ ] **Fan-in/Fan-out**: Distributing and collecting work
-   [ ] **Rate Limiting**: Controlling execution frequency
-   [ ] **Deadlock Prevention**: Avoiding common pitfalls

## 🛠️ Prerequisites

-   Go 1.21+ (uses `for range` syntax)
-   Basic understanding of Go syntax
-   Familiarity with goroutines concept

## 📖 Resources

-   [Go Concurrency Patterns](https://go.dev/blog/pipelines)
-   [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
-   [The Go Memory Model](https://go.dev/ref/mem)

## 🤝 Contributing

Feel free to add more concurrency examples or improve existing ones. Each new pattern should:

1. Have its own package/directory
2. Include clear documentation
3. Demonstrate both correct and incorrect usage where applicable
4. Update this README with the new feature

---

_Happy learning! 🎉_
