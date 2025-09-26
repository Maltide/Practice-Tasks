# Project Architecture Rules (Non-Obvious Only)

- **Package namespace**: All exercises use `package problems` despite different directories - creates namespace collision resolved by import aliases
- **Centralized testing**: Sprint-level test files (`sprint_test.go`) import from subdirectories using aliases like `findsmallestmap` - not typical Go pattern
- **Import architecture**: Full module paths required (`example.com/practice-tasks/sprints/sprint-XX-name/subdir`) - no relative imports
- **Error handling convention**: Empty inputs return zero values (not errors) - architectural decision for learning exercises
- **Sprint modularity**: Each sprint is self-contained module with own test suite, but shares package namespace across all exercises
- **Mixed-language docs**: Russian function specs + English code comments - intentional bilingual architecture for learning