# Project Documentation Rules (Non-Obvious Only)

- **Sprint organization**: `sprints/sprint-XX-name/` structure, but README progress table uses different naming (e.g., "Sprint 01: Основы" vs actual dir "sprint-01-find-smallest")
- **Test location**: Tests are centralized at sprint level (`sprint_test.go`), not alongside individual exercise files
- **Mixed documentation**: Function descriptions in Russian, code comments in English - both are canonical
- **Progress tracking**: `go run scripts/progress.go` updates README statistics (script referenced but not visible in file structure)
- **Package structure**: All exercises use `package problems` regardless of subdirectory - counterintuitive but intentional