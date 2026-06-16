# Go Password Utils

A menu-driven Go command-line application for creating and analysing passwords. Supports manual, clipboard, and random password input, with analysis, transformation, and export tools.

This is a portfolio project and a study on password security. On a production scenario, the password would not be saved as plain text, and the pepper would be stored outside of the database. Exporting to an external file is part of this exercise only and should not be handled this way in a real-world scenario.

Finally, the weights considered for the password strength evaluation were a best guess, and in a real-world scenario they should be set with inputs from a experienced professional in password security.


## Table of Contents
- [Features](#features)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)
- [Installation](#installation)
- [Running](#running)
- [Usage](#usage)
- [Security Notes](#security-notes)
- [Improvements](#improvements)
- [Contributing](#contributing)
- [License](#license)

---

## Features

### Input
- Manual password entry
- Load from clipboard
- Generate a random password

### Analysis
- Character count (no whitespace)
- Word count
- Letter count
- Uppercase letter count
- Lowercase letter count
- Numeric character count
- Special character count
- Repeated character detection
- Password strength analysis (scoring-based)

### Transformations
- Reverse password
- Scramble password
- Add pepper (manual entry)
- Generate salt (random)
- Hash password with Argon2id

### Output
- Print to screen (password, pepper, salt, hashed key)
- Copy to clipboard (choose which value)
- Export to file (`output.txt`)

---

## Project Structure

```text
gopwdutil/
├── main.go
├── ui/
│   ├── main_menu.go
│   ├── input_menu.go
│   ├── analysis_menu.go
│   ├── transform_menu.go
│   └── output_menu.go
├── utils/
│   ├── input.go
│   ├── analysis.go
│   ├── transform.go
│   └── output.go
├── analysis/
│   ├── count.go
│   ├── repeated.go
│   └── strength.go
├── transform/
│   ├── reverse.go
│   ├── scramble.go
│   ├── pepper.go
│   ├── salt.go
│   └── hash.go
├── input/
│   ├── manual.go
│   ├── from_clipboard.go
│   └── random.go
├── output/
│   ├── to_screen.go
│   ├── to_clipboard.go
│   └── to_file.go
└── tools/
    ├── tools.go
    ├── get_choice.go
    ├── contains_byte.go
    └── reset.go
```

---

## Dependencies

| Module | Purpose |
|---|---|
| `golang.org/x/crypto/argon2` | Argon2id password hashing |
| `golang.design/x/clipboard` | Clipboard read/write |

---

## Installation

### Prerequisites

- Go installed on your machine
- A terminal such as PowerShell, Command Prompt, Bash, or zsh

### Clone or initialize the project

```bash
mkdir gopwdutil
cd gopwdutil
go mod init gopwdutil
```

Then add the files and folders from the project structure above.

### Install dependencies

```bash
go mod tidy
```

### Build

#### macOS / Linux

```bash
go build -o gopwdutil
```

#### Windows

```powershell
go build -o gopwdutil.exe
```

---

## Running

### Run directly without building

```bash
go run .
```

### Run the built executable

#### macOS / Linux

```bash
./gopwdutil
```

#### Windows

```powershell
.\gopwdutil.exe
```

### Run tests

```bash
go test ./...
```

---

## Usage

The application uses a hierarchical menu system. All menus loop until you enter `0` to go back.

### Workflow

1. Start the app — the main menu appears showing current password length.
2. Choose **Input tools** to load or generate a password.
3. Choose **Analysis tools** to inspect the password.
4. Choose **Transformation tools** to modify it, add pepper/salt, or hash it.
5. Choose **Output options** to print, copy, or export the results.
6. Choose **Reset** to zero all values from memory.
7. Enter `0` from the main menu to exit.

### Menu Overview

```text
=========== Main Menu ============
Password length: 0
1) Input tools
0) Exit
```

Once a password is loaded:

```text
=========== Main Menu ============
Password length: 14
1) Input tools
2) Analysis tools
3) Transformation tools
4) Output options
5) Reset current password
0) Exit
```

**Input tools**
```text
1) Manual input
2) Load from clipboard
3) Generate random string
0) Return to Main Menu
```

**Analysis tools**
```text
1) Character count
2) Word count
3) Letter count
4) Upper count
5) Lower count
6) Numeric count
7) Special count
8) Repeated counts
9) Strength analysis
0) Return to Main Menu
```

**Transformation tools**
```text
1) Reverse password
2) Scramble password
3) Crack fresh pepper
4) Sprinkle some salt
5) Hash password
0) Return to Main Menu
```

**Output options**
```text
1) To screen
2) To clipboard
3) To file
0) Return to Main Menu
```

---

## Security Notes

- Password, pepper, salt, and hashed key are stored as `[]byte` so they can be explicitly zeroed in memory.
- All sensitive slices are zeroed (byte-by-byte) before the program exits and when reset is chosen.
- All random generation (password, salt, scramble) uses `crypto/rand`.
- Hashing uses **Argon2id** with `t=3`, `m=62 MB`, `p=1`, `keyLen=32 bytes`.
- Pepper is entered manually and never generated by the program, keeping it separate from the database.
- Clipboard is cleared before reading and after reading to avoid leaving data behind.

---

## Improvements

- Input → Random → Allow the user to choose which character sets to include.
- Tools → Refine `SpecialChar` into subcategories (punctuation, math, brackets, etc.).
- Implement batch processing of multiple passwords.

---

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

1. Fork the repository.
2. Create a new branch for your changes.
3. Make your changes and commit them.
4. Push your changes to your forked repository.
5. Open a pull request.

---

## License

This project is licensed under the [MIT License](./LICENSE.md)
