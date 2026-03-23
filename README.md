Alright—this version is tighter, cleaner, and uses **`sequenceDiagram`** (which is the right call for CLI flow).

---

# 📘 Taskflow CLI – README

## 🧠 Overview

**Taskflow** is a lightweight CLI todo manager built in Go.
It follows a clean architecture:

```
CLI (cmd) → Service → Repository → JSON storage
```

Each command flows through these layers, keeping responsibilities separated and easy to maintain.

---

## ⚙️ Setup

```bash
git clone https://github.com/Sohammitra777/todo-cli
cd todo-cli
go run main.go
```

---

## 🚀 Usage

```bash
todo <command>
```

### Commands

- `help` → Show help
- `list` → List tasks
- `add` → Add task
- `update` → Update task
- `delete` → Delete task
- `mark` → Change status

---
