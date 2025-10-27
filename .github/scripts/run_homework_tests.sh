#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

EXIT_CODE=0
[[ "${DEBUG:-}" == "1" ]] && set -x

echo "────────────────────────────────────────────"
echo "🔍 Searching for homework directories with tests..."

DIRS=$(find . -type d -path '*/homework*' | sort)

if [ -z "$DIRS" ]; then
  echo "⚠️  No homework directories found."
  exit 0
fi

for d in $DIRS; do
  # skip dirs that have no test files
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
  else
    echo "❌ FAIL $d"
    echo "▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼"
    echo "$OUTPUT"
    echo "▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲"
    EXIT_CODE=$STATUS
  fi
done

echo "────────────────────────────────────────────"
if [ $EXIT_CODE -eq 0 ]; then
  echo "🏁 All homework directories passed ✅"
else
  echo "❗ Some homework directories failed ❌"
fi

exit $EXIT_CODE