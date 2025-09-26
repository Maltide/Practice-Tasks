# Project Coding Rules (Non-Obvious Only)

- **Package naming**: ALL source files use `package problems` regardless of directory name
- **Import aliases**: Sprint tests use descriptive aliases like `findsmallestmap`, `basicsslice` when importing from subdirectories  
- **Empty input handling**: Empty maps/slices return 0 (not error), following project convention in existing solutions
- **Use `math.MaxInt`**: For finding minimums in non-empty collections (not zero initialization)
- **Import paths**: Use full module path `example.com/practice-tasks/sprints/sprint-XX-name/subdir`