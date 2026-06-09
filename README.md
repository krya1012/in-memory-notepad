# In-Memory Notepad

A simple in-memory notepad CLI built in Go — a [Hyperskill](https://hyperskill.org/projects/238) project.

## Requirements

- Go (any recent version)

## Run

```bash
go run "In-Memory Notepad/task/main.go"
```

## Usage (Stage 4)

On launch, the program asks for the maximum number of notes to store. It then runs in a loop, prompting `Enter a command and data:` on each iteration.

| Command | Description |
|---|---|
| `create <text>` | Stores a new note; errors if full or text is blank |
| `list` | Prints all notes with 1-based position, or a message if empty |
| `clear` | Deletes all notes |
| `update <pos> <text>` | Replaces the note at the given 1-based position |
| `delete <pos>` | Removes the note at the given position and re-indexes the rest |
| `exit` | Exits the program |
| *(unknown)* | `[Error] Unknown command` |
