#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

EXIT_CODE=0
[[ "${DEBUG:-}" == "1" ]] && set -x

echo "────────────────────────────────────────────"
echo "🔍 Searching for homework directories to lint..."

# Find candidate dirs. You can tune this pattern if you rename later.
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
  # Check for any .go files directly in this directory
  if ! ls "$d"/*.go >/dev/null 2>&1; then
    # nothing to lint in this dir (maybe only has subdirs)
    continue
  fi

  echo "────────────────────────────────────────────"
  echo "📂 Linting directory: $d"

  {
    set +e
    # We cd into the dir so golangci-lint will respect its local go.mod if it has one
    OUTPUT=$(cd "$d" && golangci-lint run --timeout=5m --out-format=colored-line-number . 2>&1)
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
  echo "🏁 All homework directories passed linting ✅"
else
  echo "❗ Some homework directories failed linting ❌"
fi

exit $EXIT_CODE