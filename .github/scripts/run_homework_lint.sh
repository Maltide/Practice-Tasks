#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

EXIT_CODE=0
PASSED_DIRS=""
FAILED_DIRS=""

[[ "${DEBUG:-}" == "1" ]] && set -x

echo "────────────────────────────────────────────"
echo "🔍 Searching for homework directories to lint..."

# Find dirs that look like homework, sorted for stable output
DIRS=$(find . -type d -path '*/homework*' | sort)

if [ -z "$DIRS" ]; then
  echo "⚠️  No homework directories found."
  exit 0
fi

# ensure golangci-lint exists (pin version so CI doesn't break on us)
if ! command -v golangci-lint >/dev/null 2>&1; then
  echo "Installing golangci-lint..."
  # you can bump this later intentionally, not accidentally
  GOLANGCI_LINT_VERSION="v1.58.2"
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b "$(go env GOPATH)/bin" "${GOLANGCI_LINT_VERSION}"
  export PATH="$(go env GOPATH)/bin:$PATH"
fi

for d in $DIRS; do
  # skip dirs that have no .go files directly in them
  if ! ls "$d"/*.go >/dev/null 2>&1; then
    continue
  fi

  echo "────────────────────────────────────────────"
  echo "📂 Linting directory: $d"

  # We'll run golangci-lint from within that dir so it can use a local go.mod
  set +e
  OUTPUT=$(cd "$d" && golangci-lint run --timeout=5m . 2>&1)
  STATUS=$?
  set -e

  # If it failed only because of config version, try again with a no-config fallback.
  # We detect that by grepping the error message.
  if [ $STATUS -ne 0 ] && echo "$OUTPUT" | grep -qi "can't load config"; then
    echo "⚠️  Falling back to minimal linters in $d (no valid config)"
    set +e
    OUTPUT=$(cd "$d" && golangci-lint run \
      --timeout=5m \
      --disable-all \
      --enable=govet \
      --enable=staticcheck \
      --enable=ineffassign \
      --enable=gofmt \
      . 2>&1)
    STATUS=$?
    set -e
  fi

  if [ $STATUS -eq 0 ]; then
    echo "✅ PASS $d"
    if [ -n "$PASSED_DIRS" ]; then
      PASSED_DIRS="${PASSED_DIRS}\n  • ${d}"
    else
      PASSED_DIRS="  • ${d}"
    fi
  else
    echo "❌ FAIL $d"
    echo "▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼"
    echo "$OUTPUT"
    echo "▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲"
    if [ -n "$FAILED_DIRS" ]; then
      FAILED_DIRS="${FAILED_DIRS}\n  • ${d}"
    else
      FAILED_DIRS="  • ${d}"
    fi
    EXIT_CODE=$STATUS
  fi
done

echo "────────────────────────────────────────────"
echo "📊 Lint summary"

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
  echo "🏁 All homework directories passed linting ✅"
else
  echo "❗ Some homework directories failed linting ❌"
fi

exit $EXIT_CODE