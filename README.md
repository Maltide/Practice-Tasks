# Go Practice Tasks

## 📊 Общий прогресс
- **Всего задач**: 46
- **Решено**: 23 (50%)
- **Текущий спринт**: Sprint 03

### Прогресс по спринтам
| Спринт | Задач | Решено | Прогресс |
|--------|-------|--------|----------|
| [Sprint 01: Основы](sprints/sprint-01-find-smallest/) | 4 | 4 | ✅ 100% |
| [Sprint 02: Алгоритмы](sprints/sprint-02-basics/) | 9 | 9 | ✅ 100% |
| [Sprint 03: Продвинутые](sprints/sprint-03-advanced/) | 11 | 10 | 🔄 91% |
| [Sprint 04: Конкурентность](sprints/sprint-04-goroutines/) | 10 | 0 | ⏳ 0% |
| [Sprint 05: Структуры данных](sprints/sprint-05-data-structures/) | 12 | 0 | ⏳ 0% |

## 🎯 Быстрый старт
1. Начните с [Sprint 01](sprints/sprint-01-find-smallest/)
2. Решите все задачи спринта
3. Запустите `go run scripts/progress.go` для обновления статистики
4. Переходите к следующему спринту

## 📚 Структура обучения
- **Sprint 01**: Основы работы со слайсами и мапами (поиск минимальных значений)
- **Sprint 02**: Базовые алгоритмы и структуры данных (строки, слайсы, мапы, Фибоначчи)
- **Sprint 03**: Продвинутые техники и оптимизация (слияние, кодирование, two-pointer)
- **Sprint 04**: Конкурентность в Go (горутины, каналы, паттерны параллелизма)
- **Sprint 05**: Структуры данных (стек, очередь, связный список)

## Repository Layout
- `sprints/` — задачи, сгруппированные по спринтам.
- `archive/` — старые задачи, не вошедшие в текущую структуру.

## Working on Tasks
1. Read the task description and comments inside the starter function.
2. Implement or review the solution, paying attention to edge cases like empty inputs.
3. Run the tests for that task (`go test ./sprints/sprint-01-find-smallest/ -run TestFindSmallestIntInMap` as an example) until they pass.
4. Use `go test ./...` from the repository root to make sure every exercise still succeeds.

## Authoring New Exercises
- Keep exercises self-contained inside their subdirectory with both solution and `*_test.go` files.
- Maintain the explanatory header comments that describe the problem, hints, and sample data.
- Prefer table-driven tests with `t.Run` and cover empty inputs, negative numbers, and single-item cases where relevant.
- Stick to clear, idiomatic Go without introducing external dependencies for these fundamentals.
- Update this root README and the topic-specific READMEs when you add a new set of tasks so learners know where to begin and which tests to run.

## Sprint 03 — Advanced (easiest → hardest)

Hints live in **HINTS.md** and are intended as a **last resort**. Attempt the problems unaided first.

- [x] 1) Intersection of Two Slices (slice)
- [x] 2) Remove Duplicates from Sorted Slice (slice)
- [x] 3) Rotate Slice Right by k (slice)
- [x] 4) Maximum Sum Subarray, fixed k (slice)
- [x] 5) Merge Two Sorted Slices (slice)
- [x] 6) Longest Common Prefix (string)
- [x] 7) First Non-Repeating Character (string)
- [x] 8) Run-Length Encoding (string)
- [x] 9) Run-Length Decoding (string)
- [x] 10) Invert Map with Unique Values (map)
- [ ] 11) Two Sum — return indices (mixed)

## Sprint 04 — Go Concurrency

Comprehensive course on goroutines and channels with practical exercises:

- [ ] Task 1: Parallel Sum
- [ ] Task 2: HTTP Fetcher
- [ ] Task 3: Pipeline Pattern
- [ ] Task 4: Worker Pool
- [ ] Task 5: Rate Limiter
- [ ] Task 6: Fan-Out/Fan-In
- [ ] Task 7: Timeout Pattern
- [ ] Task 8: Semaphore
- [ ] Errors: Custom error types
- [ ] Demos: Basic goroutine examples

## Sprint 05 — Data Structures

Implementation of fundamental data structures with LeetCode-style problems:

### Stack
- [ ] Basic Stack Implementation
- [ ] Valid Parentheses (LeetCode 20)
- [ ] Min Stack (LeetCode 155)
- [ ] Evaluate RPN (LeetCode 150)

### Queue
- [ ] Basic Queue Implementation
- [ ] Circular Queue (LeetCode 622)
- [ ] Stack using Queues (LeetCode 225)
- [ ] Sliding Window Maximum (LeetCode 239)

### Linked List
- [ ] Basic Linked List Implementation
- [ ] Reverse Linked List (LeetCode 206)
- [ ] Detect Cycle (LeetCode 141)
- [ ] Merge Two Sorted Lists (LeetCode 21)

## Next Steps
В разделах `slices/find-smallest` и `maps/find-smallest` лежат готовые решения поиска минимума и соответствующие тесты. Остальные задачи пока содержат только формулировки — реализуй их и добавь проверку. Расширяя упражнения, создавай новые подпапки и дополняй документацию.
