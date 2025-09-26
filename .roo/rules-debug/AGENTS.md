# Project Debug Rules (Non-Obvious Only)

- **Test isolation**: Each sprint has single test file (`sprint_test.go`) importing from all subdirectories - debug at sprint level
- **Empty input edge case**: Functions return 0 for empty inputs (not error) - this is intentional, not a bug
- **Package naming confusion**: All `.go` files use `package problems` even though they're in different directories
- **Import path debugging**: Full module paths required: `example.com/practice-tasks/sprints/sprint-XX-name/subdir`
- **Mixed language comments**: Russian function descriptions are intentional, not encoding issues