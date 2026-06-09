# In-Memory Notepad

A simple in-memory notepad CLI built in Go — a [Hyperskill](https://hyperskill.org/projects/238) project.

## Requirements

- Go (any recent version)

## Run

```bash
go run "In-Memory Notepad/task/main.go"
```

## Usage (Stage 1)

The program runs in a loop, prompting `Enter a command and data:` on each iteration.

| Input | Output |
|---|---|
| `exit` | `[Info] Bye!` — program exits |
| anything else | echoes the command and data back |
