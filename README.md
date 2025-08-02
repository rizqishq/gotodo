# 🧠 Gotodo CLI

A simple command-line interface (CLI) built with Go and [Cobra](https://github.com/spf13/cobra) to manage daily tasks efficiently from your terminal.  
Tasks are saved locally using a JSON file, and you can add, update, delete, and mark them as in-progress or done.



## 📦 Installation

To install, simply use `go install`:

```bash
go install github.com/rizqishq/gotodo
```

This will install the CLI tool named `gotodo` into your `$GOBIN`, and you'll be able to run it globally from any terminal.


## 🧪 Example Usage

The CLI uses structured subcommands (thanks to Cobra). Here's how you use it:

### ➕ Add a task

```bash
gotodo add "Learn Go"
```

### 📋 List tasks

```bash
gotodo list                  # All tasks
gotodo list todo             # Only tasks with 'todo' status
gotodo list in-progress      # Tasks in progress
gotodo list done             # Completed tasks
```

### ✏️ Update task description

```bash
gotodo update <id> "New Description"
```

### 🗑️ Delete a task

```bash
gotodo delete <id>
```

### 🚧 Mark task as in-progress

```bash
gotodo mark-in-progress <id>
```

### ✅ Mark task as done

```bash
gotodo mark-done <id>
```


## 📁 Task Storage

All tasks are saved to a local JSON file at:

```
./tasks.json
```

If this file doesn't exist, it will be created automatically on the first task addition.

---

## 📝 Task Format

Each task is stored with the following structure:

```json
{
  "id": 1,
  "description": "Belajar Go",
  "status": "todo",
  "createdAt": "2025-08-02T20:00:00Z",
  "updatedAt": "2025-08-02T20:00:00Z"
}
```

---

## 🧼 Output Example

### ✅ Add task

```bash
gotodo add "Belajar Go"

✅ Task added successfully!
🆔 ID         : 1
📄 Description: Belajar Go
📌 Status     : todo 📝
🕒 Created At : 2025-08-02 20:00
```

### 📋 List tasks

```bash
gotodo list

📋 Listing Tasks (total: 1)

[🆔 1]
📄 Description: Belajar Go
📌 Status     : todo 📝
🕒 Created At : 2025-08-02 20:00
🕒 Updated At : 2025-08-02 20:00
------------------------------------
```