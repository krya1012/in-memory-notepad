# In-Memory Notepad

A simple in-memory notepad CLI built in Go — a [Hyperskill](https://hyperskill.org/projects/238) project.

## Requirements

- Go (any recent version)

## Run

```bash
go run "In-Memory Notepad/task/main.go"
```

## Usage (Stage 2)

The program runs in a loop, prompting `Enter a command and data:` on each iteration.

| Command | Description |
|---|---|
| `create <text>` | Stores a new note (max 5) |
| `list` | Prints all stored notes with 1-based position |
| `clear` | Deletes all notes |
| `exit` | Exits the program |
