# Go Practice Tasks

## 📊 Общий прогресс
- **Всего задач**: 15
- **Решено**: 15 (100%)
- **Текущий спринт**: Sprint 02

### Прогресс по спринтам
| Спринт | Задач | Решено | Прогресс |
|--------|-------|--------|----------|
| [Sprint 01: Основы](sprints/sprint-01-basics/) | 4 | 4 | ✅ 100% |
| [Sprint 02: Алгоритмы](sprints/sprint-02-algorithms/) | 7 | 7 | ✅ 100% |
| [Sprint 03: Продвинутые](sprints/sprint-03-advanced/) | 4 | 0 | ⏳ 0% |

## 🎯 Быстрый старт
1. Начните с [Sprint 01](sprints/sprint-01-basics/)
2. Решите все задачи спринта
3. Запустите `go run scripts/progress.go` для обновления статистики
4. Переходите к следующему спринту

## 📚 Структура обучения
- **Sprint 01**: Основы работы со строками, слайсами, мапами (поиск минимума)
- **Sprint 02**: Базовые алгоритмы и структуры данных (анаграммы, баланс скобок, частоты, реверс, удаление, второй максимум, Фибоначчи)
- **Sprint 03**: Продвинутые техники и оптимизация

## Repository Layout
- `sprints/` — задачи, сгруппированные по спринтам.
- `archive/` — старые задачи, не вошедшие в текущую структуру.

## Working on Tasks
1. Read the task description and comments inside the starter function.
2. Implement or review the solution, paying attention to edge cases like empty inputs.
3. Run the tests for that task (`go test ./sprints/sprint-01-basics/maps-find-smallest/ -run TestFindSmallestIntInMap` as an example) until they pass.
4. Use `go test ./...` from the repository root to make sure every exercise still succeeds.

## Authoring New Exercises
- Keep exercises self-contained inside their subdirectory with both solution and `*_test.go` files.
- Maintain the explanatory header comments that describe the problem, hints, and sample data.
- Prefer table-driven tests with `t.Run` and cover empty inputs, negative numbers, and single-item cases where relevant.
- Stick to clear, idiomatic Go without introducing external dependencies for these fundamentals.
- Update this root README and the topic-specific READMEs when you add a new set of tasks so learners know where to begin and which tests to run.

## Next Steps
В разделах `slices/find-smallest` и `maps/find-smallest` лежат готовые решения поиска минимума и соответствующие тесты. Остальные задачи пока содержат только формулировки — реализуй их и добавь проверку. Расширяя упражнения, создавай новые подпапки и дополняй документацию.
