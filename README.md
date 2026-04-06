# B(eta) Compiler

A compiler for the B programming language built with Go and ANTLR4.

## Language Extensions

- Single-line (`//`) comments
- Top-level `extrn` declarations
- `__asm__(...)` inline assembly statements
- `name __asm__(...)` naked functions
- `__variadic__(name, num)` variadic function declarations

See [b-extensions.md](b-extensions.md) for details.

## Requirements

- Go 1.22+
- ANTLR4

## Quick Start

### Install ANTLR4 
https://www.antlr.org/download.html

```bash
# Generate parser
antlr4 -Xexact-output-dir -o internal/parser -Dlanguage=Go b.g4

# Build
go build -o bin/beta ./cmd/beta

# Run
./bin/beta your_program.b
```


## Project Structure

```
.
├── b.g4                    # ANTLR4 grammar
├── b-extensions.md         # Language extension documentation from tsoding
├── cmd/
│   └── beta/
│       └── main.go        # Entry point
├── internal/
│   ├── analysis/
│   │   └── analysis.go    # Parser & error handling
│   └── parser/            # Generated ANTLR files
└── test-syntax.b          # Syntax test file
```

