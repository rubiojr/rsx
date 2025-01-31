# Quick Start Guide

### 1. Install RSX

```bash
go install github.com/rubiojr/rsx@latest
```

If you want RSX with SQLite and FTS5 (full-text search) support:
```bash
CGO_ENABLED=1 go install --tags fts5,semver github.com/rubiojr/rsx@latest
```

### 2. Create Your First Project

```bash
# Create a new project called 'hello-world'
rsx new hello-world
cd hello-world
```

### 3. Project Structure
You'll see the following files created:
```
hello-world/
├── main.risor    # Main entry point
└── lib/          # Directory for additional modules
```

### 4. Write Your First Script

Edit `main.risor`:
```Go
import rsx

# Print a message
rsx.log("Hello, World!")

# Use some basic functionality
for range 3 {
	rsx.log("count")
}

# Use built-in SQLite
# Make sure to build RSX with SQLite and FTS5 support
import sql
db := sql.open("sqlite::memory:")
db.exec("CREATE TABLE test (name TEXT)")
db.exec("INSERT INTO test VALUES ('RSX')")
result := db.query("SELECT * FROM test")
rsx.log(result)
```

### 5. Development

During development, you can run your script directly:
```bash
rsx run
```

### 6. Build and Run

When ready to distribute:
```bash
# Build the binary
rsx build

# Run your application
./hello-world
```

### 7. Using External Modules

Add external Risor modules:
```Go
import rsx

# Load a module from GitHub
rsx.load("gh:rubiojr/risor-libs/lib/pool", { branch: "main" })

# Use the module
pool.new(2)  # Create a worker pool with 2 workers
pool.queue(func() { rsx.log("Task executed!") })
pool.wait()
```
