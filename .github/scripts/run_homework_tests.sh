#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

EXIT_CODE=0
PASSED_DIRS=""
FAILED_DIRS=""

[[ "${DEBUG:-}" == "1" ]] && set -x

echo "────────────────────────────────────────────"
echo "🔍 Searching for homework directories with tests..."

DIRS=$(find . -type d -path '*/homework*' | sort)

if [ -z "$DIRS" ]; then
  echo "⚠️  No homework directories found."
  exit 0
fi

for d in $DIRS; do
  # skip dirs with no test files
  ls "$d"/*_test.go >/dev/null 2>&1 || continue

  echo "────────────────────────────────────────────"
  echo "📁 Directory: $d"

  {
    set +e
    OUTPUT=$(cd "$d" && go test -v . 2>&1)
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
echo "📊 Test summary"

if [ -n "$PASSED_DIRS" ]; then
  echo "✔ Passed:"
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
  echo "🏁 All homework directories passed tests ✅"
else
  echo "❗ Some homework directories failed tests ❌"
fi

exit $EXIT_CODE