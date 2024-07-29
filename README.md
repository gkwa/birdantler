# birdantler

Purpose: birdantler is a command-line tool for watching directories and files for changes. It allows you to specify patterns and event types to filter the changes you're interested in.

## Installation

On macOS/Linux:
```bash
brew install gkwa/homebrew-tools/birdantler
```

On Windows:
```powershell
TBD
```

## Usage

Basic usage:
```bash
birdantler hello [flags]
```

Flags:
- `-d, --dir string`: Directory to watch for changes (default ".")
- `-p, --pattern stringSlice`: File patterns to watch (e.g. '*.txt', '*.md')
- `-t, --type string`: Filter type (create, write, remove, rename, chmod) (default "write")
- `-v, --verbose`: Enable verbose mode
- `--log-format string`: Log format (json or text, default is text)

## Examples

1. Watch current directory for all changes:
```bash
birdantler hello
```

2. Watch a specific directory for changes:
```bash
birdantler hello -d /path/to/watch
```

3. Watch for changes to markdown files:
```bash
birdantler hello -p "*.md"
```

4. Watch for changes to text and markdown files:
```bash
birdantler hello -p "*.txt" -p "*.md"
```

5. Watch for file creations:
```bash
birdantler hello -t create
```

6. Watch for file removals in a specific directory:
```bash
birdantler hello -d /path/to/watch -t remove
```

7. Watch for changes to Python files with verbose logging:
```bash
birdantler hello -p "*.py" -v
```

8. Watch for file renames with JSON logging:
```bash
birdantler hello -t rename --log-format json
```

## Cheat Sheet

| Command | Description |
|---------|-------------|
| `birdantler hello` | Watch current directory for all changes |
| `birdantler hello -d /path` | Watch specific directory |
| `birdantler hello -p "*.ext"` | Watch files with specific extension |
| `birdantler hello -t create` | Watch for file creations |
| `birdantler hello -t write` | Watch for file modifications (default) |
| `birdantler hello -t remove` | Watch for file deletions |
| `birdantler hello -t rename` | Watch for file renames |
| `birdantler hello -t chmod` | Watch for permission changes |
| `birdantler hello -v` | Enable verbose logging |
| `birdantler hello --log-format json` | Enable JSON logging format |

Remember to use quotes around file patterns when using wildcards to prevent shell expansion.
```
