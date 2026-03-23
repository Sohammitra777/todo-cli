### Usage

```bash
todo update 1 "Buy milk"
```

### Description

Updates task description by ID.

### Sequence Diagram

```mermaid
sequenceDiagram
    participant CLI
    participant CMD
    participant Service
    participant Repo
    participant File

    CLI->>CMD: HandleUpdate(args)
    CMD->>CMD: parseUpdateArgs()

    alt invalid args
        CMD-->>CLI: error + usage
    else valid
        CMD->>Service: UpdateTask(id, desc)
        Service->>Repo: GetTasks()
        Repo->>File: read JSON
        File-->>Repo: tasks

        Service->>Service: find task

        alt not found
            Service-->>CMD: error
            CMD-->>CLI: "Invalid Task Id"
        else found
            Service->>Repo: StoreFile(updated tasks)
            Repo->>File: write JSON
            CMD-->>CLI: success message
        end
    end
```

---
