
## README.md

````markdown
# Help-CLI 🚀

A high-performance, cross-platform CLI tool built in **Go** to manage and document my personal home-lab scripts and workflows.

Born out of a love for the "80s-style" binary runtime, this tool provides a centralized, colorized, and searchable index of commands for macOS, Linux, and Raspberry Pi—all from a single source of truth.

## 🛠 Features

- **Single Binary:** Zero dependencies. No `node_modules`, no Python environments.
- **Data-Driven:** All commands are stored in a Go slice for easy maintenance.
- **Cross-Platform:** One-touch builds for macOS (M3), AMD64 Linux, and ARM64 Raspberry Pi.
- **Smart Formatting:** Uses `tabwriter` for perfect terminal alignment and `fatih/color` for visual hierarchy.
- **Versioned:** Build dates are etched into the binary at compile-time using linker flags.

## 🏗 Project Structure

- `help.go`: The core logic, command registry, and UI engine.
- `build.sh`: The master build pipeline for cross-compilation.
- `dist/`: (Ignored) Destination for platform-specific binaries.

## 🚀 Getting Started

### Prerequisites

- Go 1.24+
- `github.com/fatih/color`

### Installation

1. Clone the repository to your Mac.
2. Initialize the module:
   ```bash
   go mod init help-cli
   go mod tidy
   ```
````

### Building

Use the included build script to generate binaries for all your systems:

```bash
# Standard build (outputs to ./dist)
./build.sh

# Production build (outputs to ./dist AND installs to ~/bin)
./build.sh -p

```

## 📖 Usage

Run the tool without arguments to see all available scripts:

```bash
help

```

Filter by a specific section (supports shorthands):

```bash
help go      # Show only Go tools
help d       # Show only Docker commands

```

Check the binary's health and build date:

```bash
help -v

```

## ⚙️ Maintenance: Adding Commands

To add a new script to your archive, simply add a line to the `helpRegistry` in `help.go`:

```go
var helpRegistry = []HelpItem{
    {"network", "ping-test", "Pings the home server to check connectivity"},
    // Add your new command here
}

```

## 🔒 License

Custom / Personal. Built for the Home Lab.


