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
echo "📊 Lint summary"

# Passed block
if [ -n "$PASSED_DIRS" ]; then
  # green header
  echo -e "\e[32m✔ Passed:\e[0m"
  echo -e "$PASSED_DIRS"
else
  echo -e "\e[32m✔ Passed:\e[0m"
  echo "  • (none)"
fi

echo    # blank line for readability between sections

# Failed block
if [ -n "$FAILED_DIRS" ]; then
  # red header
  echo -e "\e[31m✘ Failed:\e[0m"
  echo -e "$FAILED_DIRS"
else
  echo -e "\e[31m✘ Failed:\e[0m"
  echo "  • (none)"
fi

echo "────────────────────────────────────────────"
if [ $EXIT_CODE -eq 0 ]; then
  echo -e "🏁 \e[32mAll homework directories passed tests ✅\e[0m"
else
  echo -e "❗ \e[31mSome homework directories failed tests ❌\e[0m"
fi



exit $EXIT_CODE