#!/usr/bin/env bash
set -euo pipefail

# Quick script to run common Go development tasks. Simillar to a make file or a package.json script
# I'd normally make this in Makefile as I'm used to mac, but Windows support sucks. 

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

usage() {
  cat <<EOF
Usage: $(basename "$0") <command>

Commands:
  fmt        Run go fmt on all packages
  vet        Run go vet on all packages
  tidy       Run go mod tidy
  start      Start the Go application
  reg        Curl registration form
  fb         Curl feedback form
  survey     Curl survey form
  all        Run fmt, vet, tidy, and start
  help       Show this message
EOF
}

cmd_fmt() { go fmt ./...; }
cmd_vet() { go vet ./...; }
cmd_tidy() { go mod tidy; }
cmd_start() { go run main.go; }
cmd_fb() {
  curl "http://localhost:8080/form?id=feedback"
}
cmd_reg() {
  curl "http://localhost:8080/form?id=registration"
}
cmd_survey() {
  curl "http://localhost:8080/form?id=survey"
}
cmd_all() {
  cmd_fmt
  cmd_vet
  cmd_tidy
  cmd_start
}

if [ $# -lt 1 ]; then
  usage
  exit 1
fi

case "$1" in
  fmt) cmd_fmt ;;
  vet) cmd_vet ;;
  tidy) cmd_tidy ;;
  start) cmd_start ;;
  reg) cmd_reg ;;
  fb) cmd_fb ;;
  survey) cmd_survey ;;
  all) cmd_all ;;
  help|--help|-h) usage ;;
  *) echo "Unknown command: $1" >&2; usage; exit 2 ;;
esac
