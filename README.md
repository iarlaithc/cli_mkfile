# MkFile - Simple File Creation CLI Tool

A lightweight command-line tool for quickly creating files from your terminal.
I disliked having to remember how to make a file in my head or looking it up

## Features

- Create files with an extremely simple command
- Automatic duplicate file detection
- Clean error handling
- Cross-platform support
- Zero dependencies

## Installation

### Prerequisites

- Go 1.16 or higher
- Make sure your Go environment is properly set up (`GOPATH`, `GOBIN`)

### Building from Source

1. Clone the repository:
```bash
git clone <your-repo-url>
cd mkfile
```

2. Build the binary:
```bash
go build -o mkfile.exe
```

3. Add to system PATH:

#### Windows
- Open "Environment Variables" (Search in Start Menu)
- Under "System Variables" or "User Variables", find "Path"
- Click "Edit" → "New"
- Add the directory containing mkfile.exe
- Open a new terminal for changes to take effect

#### Unix/Linux/MacOS
```bash
sudo mv mkfile /usr/local/bin/
```

## Usage

Basic syntax:
```bash
mkfile filename.txt
```

Examples:
```bash
# Create a new text file
mkfile document.txt

# Create a source file
mkfile main.go

# Create a file in a specific directory
mkfile path/to/file.txt
```

### Error Messages

The tool provides clear feedback for common scenarios:

- If no filename is provided:
  ```
  Usage: mkfile <filename>
  ```

- If file already exists:
  ```
  File Already Exists
  ```

- If file creation fails:
  ```
  Failed To Create File: <error details>
  ```

## Project Structure

```
mkfile/
├── main.go        # Entry point and CLI logic
├── mkfile/        # Core package
│   └── mkfile.go  # File creation functions
└── go.mod        # Module definition
```

## Development

### Running Tests

```bash
go test ./mkfile
```

### Building

To create a custom-named binary:
```bash
go build -o customname.exe
```

## License
MIT License
