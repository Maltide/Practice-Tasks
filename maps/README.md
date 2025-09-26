# Задачи на мапы

## Структура
- `find-smallest/` — поиск минимума в простых и вложенных мапах.
- `basics/` — частоты элементов и выделение уникальных значений.

## Задачи и тесты
### `find-smallest/`
1. `FindSmallestIntInMap(m map[int]int) int`
   - Минимум среди значений мапы.
   - Тест: `go test ./maps/find-smallest -run TestFindSmallestIntInMap`.
2. `FindSmallestIntInMapMap(m map[int]map[int]int) int`
   - Минимум в вложенных мапах.
   - Тест: `go test ./maps/find-smallest -run TestFindSmallestIntInMapMap`.

### `basics/`
3. `CountFrequencies(items []string) map[string]int`
   - Построение частот элементов.
4. `UniqueValues(nums []int) []int`
   - Слайс уникальных значений с сохранением порядка.

## Как работать
- Реализуй функции в нужной поддиректории вместо `panic("not implemented")`.
- Используй table-driven тесты и запускай их через `go test ./maps/<поддиректория> -run TestName` или общий `go test ./maps/...`.
- Проверь крайние случаи: пустые мапы, отрицательные значения, пустые внутренние структуры.

## Что дальше?
Новые упражнения раскладывай по отдельным поддиректориям и фиксируй их в README вместе с командами тестирования.
