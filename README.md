# In-Memory Notepad

A simple in-memory notepad CLI built in Go — a [Hyperskill](https://hyperskill.org/projects/238) project.

## Requirements

- Go (any recent version)

## Run

```bash
go run "In-Memory Notepad/task/main.go"
```

## Usage (Stage 3)

On launch, the program asks for the maximum number of notes to store. It then runs in a loop, prompting `Enter a command and data:` on each iteration.

| Command | Description |
|---|---|
| `create <text>` | Stores a new note; errors if full or text is blank |
| `list` | Prints all notes with 1-based position, or a message if empty |
| `clear` | Deletes all notes |
| `exit` | Exits the program |
| *(unknown)* | `[Error] Unknown command` |
