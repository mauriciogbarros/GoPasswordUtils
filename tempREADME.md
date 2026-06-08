# CLI Text Utilities

A small interactive Go command-line application that lets you enter text and run simple analyses and transformations on it.

## Planned Features

- Character count
- Word count
- Vowel count
- Case detection
- Reverse text
- Password strength check

## Suggested Project Structure

```text
textutils/
├── main.go
├── ui/
│   └── menu.go
├── analysis/
│   ├── count.go
│   ├── casecheck.go
│   └── password.go
├── transform/
│   └── reverse.go
└── input/
    └── prompt.go
```

## Prerequisites

- Go installed on your machine
- A terminal such as PowerShell, Command Prompt, Bash, or zsh

## Initialize the Project

If you are starting from scratch, create the folder and initialize a Go module:

```bash
mkdir textutils
cd textutils
go mod init textutils
```

Then add the files and folders from the project spec.

## Build Instructions

From the project root, run:

```bash
go build -o textutils
```

This builds an executable named `textutils` in the current folder.

### Windows

```powershell
go build -o textutils.exe
```

## Run Instructions

### Run directly without building a named executable

```bash
go run .
```

### Run the built executable

#### macOS / Linux

```bash
./textutils
```

#### Windows

```powershell
.\textutils.exe
```

## Example Workflow

1. Start the app.
2. Enter a piece of text when prompted.
3. Pick a menu option.
4. Review the result.
5. Choose another option or exit.

## Example Commands During Development

### Format the code

```bash
go fmt ./...
```

### Run a quick build check

```bash
go build ./...
```

## Notes

- This README assumes the application entry point is `main.go` in the project root.
- If you choose a different module name, replace `textutils` in the `go mod init` command.
- If you later add tests, you can run them with:

```bash
go test ./...
```

## Next Step

Once the README is in place, the fastest next move is to implement:

1. `main.go`
2. menu loop
3. text input
4. one utility such as character count

That gives you a working vertical slice before you split more logic into packages.
