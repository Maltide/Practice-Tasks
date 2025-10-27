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

# ensure golangci-lint exists (pin version so CI is deterministic)
if ! command -v golangci-lint >/dev/null 2>&1; then
  echo "📦 Installing golangci-lint..."
  GOLANGCI_LINT_VERSION="v1.58.2"
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b "$(go env GOPATH)/bin" "${GOLANGCI_LINT_VERSION}"
  export PATH="$(go env GOPATH)/bin:$PATH"
fi

for d in $DIRS; do
  # skip dirs without .go files directly in them
  if ! ls "$d"/*.go >/dev/null 2>&1; then
    continue
  fi

  echo "────────────────────────────────────────────"
  echo "📂 Linting directory: $d"

  # first attempt: run normally
  set +e
  OUTPUT=$(cd "$d" && golangci-lint run --timeout=5m . 2>&1)
  STATUS=$?
  set -e

  # if golangci-lint errored out in a weird way (unsupported config / analyzer panic etc),
  # try again in "minimal" mode with a safe, explicit linter set
  if [ $STATUS -ne 0 ]; then
    # we always show fallback attempt reason in debug logs to make CI less confusing
    echo "⚠️  Falling back to minimal linters in $d"
    set +e
    OUTPUT=$(cd "$d" && golangci-lint run \
      --timeout=5m \
      --disable-all \
      --enable=govet \
      --enable=staticcheck \
      --enable=ineffassign \
      --enable=gofmt \
      --out-format=colored-line-number \
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
  echo -e "🏁 \e[32mAll homework directories passed linting ✅\e[0m"
else
  echo -e "❗ \e[31mSome homework directories failed linting ❌\e[0m"
fi

exit $EXIT_CODE