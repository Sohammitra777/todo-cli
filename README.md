# 📘 Taskflow CLI (Go • Fully Type-Safe)

## 🧠 Overview

**Taskflow** is a lightweight CLI todo manager built in Go.
It follows a clean architecture:

```
CLI (cmd) → Service → Repository → JSON storage
```

Each command flows through these layers, keeping responsibilities separated and easy to maintain.

A sample solution for the Task Tracker CLI challenge from roadmap.sh.

This project demonstrates a fully type-safe command-line application built in Go, capable of managing tasks (add, update, delete, mark status, and list) with persistent JSON storage.

Challenge: https://roadmap.sh/projects/task-tracker

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
taskflow <command>
```

---

## 📚 Command Documentation

Each command has detailed documentation and sequence diagrams:

* ➕ [Add Task](./docs/add.md)
* 📋 [List Tasks](./docs/list.md)
* ✏️ [Update Task](./docs/update.md)
* ❌ [Delete Task](./docs/delete.md)
* ✅ [Mark Task](./docs/mark.md)

---

## 🧾 Available Commands

| Command | Description        |
| ------- | ------------------ |
| help    | Show help menu     |
| list    | List tasks         |
| add     | Add new task       |
| update  | Update task        |
| delete  | Delete task        |
| mark    | Update task status |

---

## 🆘 Help

```bash
taskflow help
```

Displays all available commands and usage instructions.

---

## 🏗️ Project Structure

```
.
├── cmd/        # CLI handlers
├── service/    # Business logic
├── repo/       # Data persistence (JSON)
├── utils/      # Helper functions
├── docs/       # Command documentation
└── main.go     # Entry point
```

---

## 🧠 Design Philosophy

* Simple and minimal
* Clear separation of concerns
* Easy to extend (add DB, API, etc.)

---

## 👨‍💻 Author

Soham Mitra
GitHub: https://github.com/Sohammitra777

---

## 🙏 Final Note

Thank you—at last—for taking the time to check out this project.
If you found it useful or interesting, feel free to explore, contribute, or share feedback.
