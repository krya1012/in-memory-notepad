# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a [Hyperskill](https://hyperskill.org/projects/238) Go project — an in-memory notepad CLI program built in four incremental stages. The sole source file is `In-Memory Notepad/task/main.go`.

## Commands

```bash
# Run the program
go run "In-Memory Notepad/task/main.go"

# Run the Hyperskill test suite (requires Python + hs-test-python)
pip install -r requirements.txt
python "In-Memory Notepad/task/tests.py"
```

## Architecture

There is a single file: `In-Memory Notepad/task/main.go`. No packages, no external Go dependencies.

The program is a REPL loop:
1. On startup, prompt `Enter the maximum number of notes:` and read an integer (the storage cap).
2. Loop: print `Enter a command and data:`, read a full line, split on the first space into `command` and `data`.
3. Dispatch on `command`:

| Command  | Behavior |
|----------|----------|
| `create` | Append `data` (trimmed) to the slice; errors if full or data is blank |
| `list`   | Print `[Info] N: text` for each note; print `[Info] Notepad is empty` if none |
| `clear`  | Delete all notes |
| `update` | Replace note at 1-based position; errors on missing/invalid position or blank note |
| `delete` | Remove note at position and re-index remaining notes |
| `exit`   | Print `[Info] Bye!` and terminate |
| *other*  | Print `[Error] Unknown command` |

Use `bufio.Scanner` + `strings.SplitN` (or slice expressions) to handle whitespace-containing input correctly.

## Stage Progression

The lesson has four stages, each adding to the previous:
1. **Basic CLI program** — REPL loop, echo unknown commands, handle `exit`
2. **Core functionality** — `create` (fixed cap 5), `list`, `clear`
3. **Dynamic storage** — cap read from stdin, `[Info] Notepad is empty`, `[Error] Unknown command`, validate blank `create`
4. **Working with the notes** — `update` and `delete` with full position validation

## Test Framework

Tests are in `In-Memory Notepad/task/tests.py` using the `hs-test-python` framework (`TestedProgram` drives the Go binary via stdin/stdout). The dependency is installed from the URL in `requirements.txt`.
