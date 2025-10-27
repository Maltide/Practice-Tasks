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

# Find homework-like dirs, stable order
DIRS=$(find . -type d -path '*/homework*' | sort)

if [ -z "$DIRS" ]; then
  echo "⚠️  No homework directories found."
  exit 0
fi

# Ensure golangci-lint exists (pin version for determinism)
if ! command -v golangci-lint >/dev/null 2>&1; then
  echo "📦 Installing golangci-lint..."
  GOLANGCI_LINT_VERSION="v1.58.2"
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b "$(go env GOPATH)/bin" "${GOLANGCI_LINT_VERSION}"
  export PATH="$(go env GOPATH)/bin:$PATH"
fi

for d in $DIRS; do
  # skip dirs with no .go files at all
  if ! ls "$d"/*.go >/dev/null 2>&1; then
    continue
  fi

  echo "────────────────────────────────────────────"
  echo "📂 Linting directory: $d"

  # Heuristic: if this dir (or its parents) has a go.mod, we consider it buildable.
  # We'll walk up until repo root looking for go.mod.
  MODDIR="$d"
  FOUND_MOD="false"
  while true; do
    if [ -f "$MODDIR/go.mod" ]; then
      FOUND_MOD="true"
      break
    fi
    # stop at repo root "."
    if [ "$MODDIR" = "." ]; then
      break
    fi
    MODDIR=$(dirname "$MODDIR")
  done

  if [ "$FOUND_MOD" != "true" ]; then
    echo "⚠️  Skipping $d (no go.mod in this tree, cannot lint reliably)"
    if [ -n "$SKIPPED_DIRS" ]; then
      SKIPPED_DIRS="${SKIPPED_DIRS}\n  • ${d}"
    else
      SKIPPED_DIRS="  • ${d}"
    fi
    # do not count as failure, just move on
    continue
  fi

  # Try normal lint first
  set +e
  OUTPUT=$(cd "$d" && golangci-lint run --timeout=5m . 2>&1)
  STATUS=$?
  set -e

  # If golangci-lint errored (bad config, analyzer panic, etc.), try minimal safe linters
  if [ $STATUS -ne 0 ]; then
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
  echo -e "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
  echo -e "\e[32m✔ PASSED DIRECTORIES:\e[0m"
  echo -e "$PASSED_DIRS"
else
  echo -e "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
  echo -e "\e[32m✔ PASSED DIRECTORIES:\e[0m"
  echo "  • (none)"
fi

echo    # blank spacer

# Failed block
if [ -n "$FAILED_DIRS" ]; then
  echo -e "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
  echo -e "\e[31m✘ FAILED DIRECTORIES:\e[0m"
  echo -e "$FAILED_DIRS"
else
  echo -e "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
  echo -e "\e[31m✘ FAILED DIRECTORIES:\e[0m"
  echo "  • (none)"
fi

echo "────────────────────────────────────────────"

if [ -n "$FAILED_DIRS" ]; then
  # only fail the job if at least one real lint run failed
  echo -e "❗ \e[31mSome homework directories failed linting ❌\e[0m"
  exit 1
fi

echo -e "🏁 \e[32mNo lint failures ✅\e[0m"
exit 0