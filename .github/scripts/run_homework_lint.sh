#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

EXIT_CODE=0
PASSED_DIRS=""
FAILED_DIRS=""

[[ "${DEBUG:-}" == "1" ]] && set -x

echo "────────────────────────────────────────────"
echo "🔍 Searching for homework directories to lint..."

# Find dirs that look like homework, sort for deterministic output
DIRS=$(find . -type d -path '*/homework*' | sort)

if [ -z "$DIRS" ]; then
  echo "⚠️  No homework directories found."
  exit 0
fi

# ensure golangci-lint exists
if ! command -v golangci-lint >/dev/null 2>&1; then
  echo "Installing golangci-lint..."
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b "$(go env GOPATH)/bin" latest
  export PATH="$(go env GOPATH)/bin:$PATH"
fi

for d in $DIRS; do
  # skip dirs without .go files directly in them
  if ! ls "$d"/*.go >/dev/null 2>&1; then
    continue
  fi

  echo "────────────────────────────────────────────"
  echo "📂 Linting directory: $d"

  {
    set +e
    OUTPUT=$(cd "$d" && golangci-lint run --timeout=5m --out-format=colored-line-number . 2>&1)
    STATUS=$?
    set -e
  }

  if [ $STATUS -eq 0 ]; then
    echo "✅ PASS $d"
    PASSED_DIRS="${PASSED_DIRS}\n  • ${d}"
  else
    echo "❌ FAIL $d"
    echo "▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼"
    echo "$OUTPUT"
    echo "▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲"
    FAILED_DIRS="${FAILED_DIRS}\n  • ${d}"
    EXIT_CODE=$STATUS
  fi
done

echo "────────────────────────────────────────────"
echo "📊 Lint summary"

if [ -n "$PASSED_DIRS" ]; then
  echo "✔ Passed:"
  # echo -e is bashism but we're already running bash
  echo -e "$PASSED_DIRS"
else
  echo "✔ Passed:"
  echo "  • (none)"
fi

if [ -n "$FAILED_DIRS" ]; then
  echo "✘ Failed:"
  echo -e "$FAILED_DIRS"
else
  echo "✘ Failed:"
  echo "  • (none)"
fi

echo "────────────────────────────────────────────"
if [ $EXIT_CODE -eq 0 ]; then
  echo "🏁 All homework directories passed linting ✅"
else
  echo "❗ Some homework directories failed linting ❌"
fi

exit $EXIT_CODE