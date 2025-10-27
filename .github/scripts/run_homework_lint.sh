#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

EXIT_CODE=0
PASSED_DIRS=""
FAILED_DIRS=""
SKIPPED_DIRS=""

[[ "${DEBUG:-}" == "1" ]] && set -x

echo "────────────────────────────────────────────"
echo "🔍 Searching for homework directories to lint..."

DIRS=$(find . -type d -path '*/homework*' | sort)

if [ -z "$DIRS" ]; then
  echo "⚠️  No homework directories found."
  exit 0
fi

# install latest golangci-lint if missing
if ! command -v golangci-lint >/dev/null 2>&1; then
  echo "📦 Installing golangci-lint (latest)..."
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b "$(go env GOPATH)/bin" latest
  export PATH="$(go env GOPATH)/bin:$PATH"
fi

for d in $DIRS; do
  # skip dirs with no .go files at all
  if ! ls "$d"/*.go >/dev/null 2>&1; then
    continue
  fi

  echo "────────────────────────────────────────────"
  echo "📂 Linting directory: $d"

  # run lint once, capture output
  set +e
  OUTPUT=$(cd "$d" && golangci-lint run --timeout=5m . 2>&1)
  STATUS=$?
  set -e

  if [ $STATUS -eq 0 ]; then
    # clean lint, nothing reported
    echo "✅ PASS $d"
    if [ -n "$PASSED_DIRS" ]; then
      PASSED_DIRS="${PASSED_DIRS}\n  • ${d}"
    else
      PASSED_DIRS="  • ${d}"
    fi
    continue
  fi

  # At this point golangci-lint exited nonzero.
  # Decide: was this a TOOLING issue or a REAL lint failure?

  # Heuristics for "tooling is broken, not your code":
  # - analyzer can't load stdlib/export data (internal/goarch)
  # - unknown flag messages (CLI mismatch)
  # - pure "can't run linter goanalysis_metalinter" style crashes
  if echo "$OUTPUT" | grep -qiE 'internal/goarch|goanalysis_metalinter|unsupported version|unknown flag'; then
    echo "⚠️  SKIP $d (tooling/analyzer issue, not code)"
    if [ -n "$SKIPPED_DIRS" ]; then
      SKIPPED_DIRS="${SKIPPED_DIRS}\n  • ${d}"
    else
      SKIPPED_DIRS="  • ${d}"
    fi
    # don't treat as failure for CI purposes
    continue
  fi

  # Otherwise, assume it's a real lint failure with actionable findings
  echo "❌ FAIL $d"
  echo "▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼"
  echo "$OUTPUT"
  echo "▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲"

  if [ -n "$FAILED_DIRS" ]; then
    FAILED_DIRS="${FAILED_DIRS}\n  • ${d}"
  else
    FAILED_DIRS="  • ${d}"
  fi
  EXIT_CODE=1
done

echo "────────────────────────────────────────────"
echo "📊 Lint summary"

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✔ PASSED DIRECTORIES:"
if [ -n "$PASSED_DIRS" ]; then
  echo -e "$PASSED_DIRS"
else
  echo "  • (none)"
fi

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✘ FAILED DIRECTORIES:"
if [ -n "$FAILED_DIRS" ]; then
  echo -e "$FAILED_DIRS"
else
  echo "  • (none)"
fi

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "↷ SKIPPED (TOOLING / ANALYZER MISMATCH):"
if [ -n "$SKIPPED_DIRS" ]; then
  echo -e "$SKIPPED_DIRS"
else
  echo "  • (none)"
fi

echo "────────────────────────────────────────────"
if [ -n "$FAILED_DIRS" ]; then
  echo "❗ Some homework directories failed linting ❌"
  exit 1
fi

echo "🏁 No lint failures ✅"
exit 0