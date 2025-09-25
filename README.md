# Practice Tasks

This repository collects small Go practice exercises focused on fundamental algorithms and data structures. Each topic lives in its own directory and may be further split into subdirectories by theme. Every section ships with problem statements, solutions, tests, and guidance so that new tasks can be added later.

## Progress Checklist
### `strings/`
- [x] `IsPalindrome(s string) bool`
- [x] `AreAnagrams(a, b string) bool`
- [ ] `IsBracketSequenceBalanced(s string) bool`

### `slices/find-smallest/`
- [x] `FindSmallestInt(nums []int) int`
- [x] `FindSmallestTime(times []time.Time) time.Time`

### `slices/basics/`
- [ ] `SumAndAverage(nums []int) (int, float64)`
- [ ] `ReverseSlice(nums []int) []int`
- [ ] `RemoveAtIndex(nums []int, i int) []int`
- [ ] `SecondLargest(nums []int) int`

### `maps/find-smallest/`
- [x] `FindSmallestIntInMap(m map[int]int) int`
- [x] `FindSmallestIntInMapMap(m map[int]map[int]int) int`

### `maps/basics/`
- [ ] `CountFrequencies(items []string) map[string]int`
- [ ] `UniqueValues(nums []int) []int`

### `misc/`
- [ ] `Fibonacci(n int) []int`

## Getting Started
- Install Go 1.21+ (or your team's current toolchain).
- Clone this repository and explore the topic directories.
- Open the README in the topic you want to study (for example `slices/README.md`) to see the recommended order of tasks and test commands.

## Repository Layout
- `strings/` — формулировки задач на палиндром, анаграммы и баланс скобок без готовых решений.
- `slices/` — поддиректории с задачами на слайсы (`find-smallest/`, `basics/`).
- `maps/` — поддиректории с задачами на отображения (`find-smallest/`, `basics/`).
- `misc/` — числовые задачи; сейчас содержит генератор чисел Фибоначчи.

Add new topics by creating additional directories at the repo root that follow the same pattern: exported functions under the `problems` package, an explanatory README, and matching tests. Внутри тем группируйте упражнения по подпапкам, чтобы не захламлять каталог.

## Working on Tasks
1. Read the task description and comments inside the starter function.
2. Implement or review the solution, paying attention to edge cases like empty inputs.
3. Run the tests for that task (`go test ./slices/find-smallest -run TestFindSmallestInt` as an example) until they pass.
4. Use `go test ./...` from the repository root to make sure every exercise still succeeds.

## Authoring New Exercises
- Keep exercises self-contained inside their subdirectory with both solution and `*_test.go` files.
- Maintain the explanatory header comments that describe the problem, hints, and sample data.
- Prefer table-driven tests with `t.Run` and cover empty inputs, negative numbers, and single-item cases where relevant.
- Stick to clear, idiomatic Go without introducing external dependencies for these fundamentals.
- Update this root README and the topic-specific READMEs when you add a new set of tasks so learners know where to begin and which tests to run.

## Next Steps
В разделах `slices/find-smallest` и `maps/find-smallest` лежат готовые решения поиска минимума и соответствующие тесты. Остальные задачи пока содержат только формулировки — реализуй их и добавь проверку. Расширяя упражнения, создавай новые подпапки и дополняй документацию.
