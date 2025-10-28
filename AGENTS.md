# AGENTS.md

Guidance for agents working with the Go Practice Tasks repository.

## 📋 Project Overview

This repository is a comprehensive Go learning resource organized into progressive sprints, covering fundamental data structures, algorithms, and concurrent programming patterns. Each sprint builds on previous knowledge with increasingly complex challenges.

### Sprint Structure
- **Sprint 01: Основы (Basics)** — Fundamental operations with maps and slices (finding minimum values)
- **Sprint 02: Алгоритмы (Algorithms)** — Basic algorithms and data structures (strings, slices, maps, Fibonacci)
- **Sprint 03: Продвинутые (Advanced)** — Complex algorithms and optimization techniques (merging, encoding, two-pointer patterns)
- **Sprint 04: Go Concurrency Course** — Goroutines, channels, and concurrent programming patterns (parallel processing, worker pools, rate limiting, timeouts)

## 🏗️ Directory Structure

```
sprints/
├── sprint-01-find-smallest/
│   ├── README.md
│   └── homework/
│       ├── sprint_test.go
│       ├── map/
│       │   ├── int.go
│       │   └── nested.go
│       └── slice/
│           ├── int.go
│           └── time.go
├── sprint-02-basics/
│   ├── README.md
│   └── homework/
│       ├── sprint_test.go
│       ├── basics-map/
│       ├── basics-slice/
│       ├── basics-strings/
│       ├── misc/
│       └── the-one-with-the-star/
├── sprint-03-advanced/
│   ├── README.md
│   └── homework/
│       ├── sprint_test.go
│       ├── map/
│       ├── slice/
│       ├── string/
│       └── mixed/
└── sprint-04-goroutines/
    ├── README.md
    ├── Makefile
    ├── demos/
    ├── channels/
    └── homework/
        ├── tasks/
        │   ├── task1_parallel_sum.go
        │   ├── task2_http_fetch.go
        │   ├── task3_pipeline.go
        │   ├── task4_worker_pool.go
        │   ├── task5_rate_limiter.go
        │   ├── task6_fan_out_in.go
        │   ├── task7_timeout.go
        │   ├── task8_semaphore.go
        │   └── tasks_test.go
        └── errors/
```

## 🧪 Testing Commands

### Sprints 01-03 (Standard Go Testing)
```bash
# Run all tests in repository
go test ./...

# Run tests for specific sprint
go test ./sprints/sprint-01-find-smallest/...
go test ./sprints/sprint-02-basics/...
go test ./sprints/sprint-03-advanced/...

# Run specific test
go test ./sprints/sprint-01-find-smallest/ -run TestFindSmallestIntInMap

# Run with verbose output
go test ./sprints/sprint-02-basics/ -v
```

### Sprint 04 (Makefile-based Testing)
```bash
# Navigate to sprint-04-goroutines directory
cd sprints/sprint-04-goroutines

# Run all tests
make test

# Run tests with verbose output
make test-verbose

# Run specific task tests
make test-task1
make test-task2
make test-task3
make test-task4

# Run benchmarks
make bench
make bench-task1
make bench-task4

# Run with race condition detection
make race

# Generate coverage report
make coverage

# Run demo examples
make run-demos
make run-channels
```

## 📝 Project-Specific Conventions

### Package Naming
- **ALL source files use `package problems`** regardless of directory name or sprint
- This ensures consistent imports across the entire project

### Import Aliases
- Sprint tests use descriptive aliases when importing from subdirectories
- Examples: `findsmallestmap`, `basicsslice`, `advancedslice`
- Helps clarify which sprint/category the imported functions come from

### Empty Input Handling
- Empty maps/slices return `0` (not error), following project convention
- This applies consistently across all sprints
- Simplifies testing and reduces error handling complexity

### Test Organization
- **Sprints 01-03**: Tests are at sprint level in `sprint_test.go`, importing from all subdirectories
- **Sprint 04**: Tests are in `tasks_test.go` within the `homework/tasks/` directory
- Individual exercise files do NOT have their own test files

### Language Mix
- Function descriptions and algorithm explanations are in **Russian**
- Implementation comments and code logic are in **English**
- This bilingual approach supports both Russian-speaking learners and international developers

## 💻 Code Style

### Testing Patterns
- **Table-driven tests** with `t.Run()` for organizing test cases
- Comprehensive edge case coverage:
  - Empty inputs (empty slices, maps, strings)
  - Negative numbers and boundary values
  - Single-item inputs
  - Large datasets for performance testing
- **Benchmarks** included for performance-critical functions

### Function Documentation
- Detailed Russian comments explaining:
  - Algorithm approach and complexity
  - Input/output specifications
  - Examples with expected results
  - Edge cases and constraints
- Clear English inline comments for implementation details

### Algorithm Optimization
- Use `math.MaxInt` for finding minimums in non-empty collections
- Two-pointer techniques for sorted array operations
- Goroutines and channels for concurrent processing (Sprint 04)
- Proper synchronization with `sync.WaitGroup` and context cancellation

## 🚀 Working with Different Sprints

### Sprints 01-03 (Algorithm Challenges)
1. Read the function's Russian documentation comments
2. Understand the problem requirements and examples
3. Implement the solution in the corresponding `.go` file
4. Run tests: `go test ./sprints/sprint-XX-name/ -v`
5. Verify all edge cases pass
6. Check overall project: `go test ./...`

### Sprint 04 (Concurrency Course)
1. Study the demos: `make run-demos && make run-channels`
2. Read the task file comments for requirements
3. Implement the function in `homework/tasks/taskN_*.go`
4. Run task-specific tests: `make test-taskN`
5. Check for race conditions: `make race`
6. Benchmark performance: `make bench-taskN`
7. Review coverage: `make coverage`

## 📚 Key Concepts by Sprint

### Sprint 01: Basics
- Map operations and iteration
- Slice operations and indexing
- Finding minimum values in collections
- Handling empty inputs

### Sprint 02: Algorithms
- String manipulation (palindromes, anagrams)
- Slice operations (reverse, remove, find second max)
- Map operations (frequencies, unique values)
- Bracket balancing
- Fibonacci sequence generation

### Sprint 03: Advanced
- Slice intersection and merging
- Duplicate removal from sorted slices
- Array rotation
- Maximum sum subarrays
- String encoding/decoding (RLE)
- Longest common prefix
- Two-sum problem with indices

### Sprint 04: Concurrency
- Goroutines and parallel processing
- Channels (buffered, unbuffered, directional)
- WaitGroup synchronization
- Worker pools and job distribution
- Rate limiting and throttling
- Fan-out/fan-in patterns
- Timeout handling
- Semaphore patterns
- HTTP concurrent requests

## 🔍 Debugging Tips

- Use `go test -v` for verbose output showing each test case
- Use `make race` in Sprint 04 to detect race conditions
- Check function documentation comments for expected behavior
- Review test cases to understand edge cases
- Use `go run` to execute demo files for learning patterns

## 📖 Additional Resources

- Each sprint has its own `README.md` with specific task descriptions
- `HINTS.md` contains hints for Sprint 03 (use as last resort)
- Sprint 04 includes `channels/README.md` with detailed channel patterns
- Makefile in Sprint 04 documents all available commands