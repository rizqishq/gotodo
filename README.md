# 🧠 Task Tracker CLI

A simple command-line interface (CLI) application written in Go to track your daily tasks.  
You can add, update, delete, and mark tasks as in-progress or done — all saved locally in a JSON file.

---

## 🚀 Features

- ✅ Add new tasks  
- 📋 List all tasks or filter by status  
- ✏️ Update existing task descriptions  
- 🗑️ Delete tasks  
- 🚧 Mark tasks as in-progress  
- ✅ Mark tasks as done  
- 📁 Persistent storage using a JSON file (`tasks.json`)

---

## 🛠 Requirements

- Go 1.18 or later  
- No external libraries required

---

## 📦 Installation

Clone the repository:

   ```bash
   git clone https://github.com/rizqishq/ToDo-CLI
   cd ToDo-CLI
   ```
---

## 📚 Usage

> The CLI uses **positional arguments**, so no flags are needed.

### ➕ Add a task

You can add a task using **add** command:
```bash
go run . add "Write blog post"
-- or
go run *.go add "Write blog post"
```

### 📋 List tasks

```bash
go run . list                    # List all tasks
go run . list todo               # List only todo tasks
go run . list in-progress        # List in-progress tasks
go run . list done               # List completed tasks
```

### ✏️ Update task description

```bash
go run . update <id> "New description"
```

### 🗑️ Delete a task

```bash
go run . delete <id>
```

### 🚧 Mark task as in-progress

```bash
go run . mark-in-progress <id>
```

### ✅ Mark task as done

```bash
go run . mark-done <id>
```

---

## 📝 Task Format

Each task is stored in `tasks.json` in the following format:

```json
{
  "id": 1,
  "description": "Write blog post",
  "status": "todo",
  "createdAt": "2025-07-13T10:20:00Z",
  "updatedAt": "2025-07-13T10:20:00Z"
}
```

---

## 📁 File Storage

All tasks are stored in:
```text
./tasks.json
```

If the file does not exist, it will be created automatically when the first task is added.

---

## 🧼 Output Example
- Add task:

```bash
$ go run . add "Finish Go CLI"

✅ Task added successfully!
🆔 ID         : 1
📄 Description: Finish Go CLI
📌 Status     : todo 📝
🕒 Created At : 2025-07-13 18:00
```
- List task:
```bash
$ go run . list

📋 Listing Tasks (total: 1)

[🆔 1]
📄 Description: Finish Go CLI
📌 Status     : todo 📝
🕒 Created At : 2025-07-13 18:00
🕒 Updated At : 2025-07-13 18:00
------------------------------------
```

---

## 🧪 Testing

Try each of the following:
- `add` a few tasks
- `list` to view them
- `update` to change descriptions
- `delete` to remove
- `mark-in-progress` and `mark-done` to change status

---