package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

// This variable is populated by build.sh at compile time
var (
    buildDate string = "Unknown" // Default value, will be overridden by build.sh
    gitHash   string = "unknown" // Default value, will be overridden by build.sh
)

// Global UI Styles
var (
	header  = color.New(color.FgCyan, color.Bold).SprintFunc()
	cmdCol  = color.New(color.FgYellow).SprintFunc()
	dim     = color.New(color.Faint).SprintFunc()
)

// 1. Define the structure
type HelpItem struct {
	Section string
	Command string
	Details string
}

// 2. Create the data store (the "Source of Truth")
var helpRegistry = []HelpItem{
	{"general", "dnsreset", "When websites are not loading in MacOS browser"},
	{"docker", "doccker ps", "List running containers, -a for all"},
	{"go", "go build -o help help.go", "build for macos"},
	{"go", "GOOS=linux GOARCH=amd64 go build -o help-linux help.go", "build for linux"},
	{"go", "GOOS=linux GOARCH=arm64 go build -o help-pi64 help.go", "build for pi"},
}

func main() {

	// 1. Version Check
    if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version" || os.Args[1] == "version") {
        fmt.Printf("Help Tool v1.0\n")
        fmt.Printf("Version: %s (%s)\n", gitHash, buildDate)
        return
    }

	// 2. Setup the tabwriter for aligned output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer w.Flush()

	// Determine what to filter for
	filter := ""
	if len(os.Args) > 1 {
		filter = strings.ToLower(os.Args[1])
		// Support "help [section]" syntax by checking the second arg
		if filter == "help" && len(os.Args) > 2 {
			filter = strings.ToLower(os.Args[2])
		}
	}

	// Logic for shorthand
	if filter == "g" { filter = "general" }
	if filter == "d" { filter = "docker" }

	if filter == "" {
		fmt.Println(color.New(color.Bold, color.Underline).Sprint("AVAILABLE SCRIPTS (ALL)"))
	}

	currentSection := ""

	// 3. The Loop: One loop to rule them all
	for _, item := range helpRegistry {
		// Filter logic: if filter is empty, show all. Otherwise, match section.
		if filter != "" && item.Section != filter {
			continue
		}

		// Print Section Header only when it changes
		if item.Section != currentSection {
			currentSection = item.Section
			fmt.Fprintln(w, header("\n"+strings.ToUpper(currentSection)+" COMMANDS"))
		}

		// Print the Command Row
		fmt.Fprintf(w, "  %s\t%s\n", cmdCol(item.Command), dim(item.Details))
	}

	w.Flush() // Ensure all output is printed before the footer

	// Add a footer if showing all scripts
    if filter == "" {
        fmt.Printf("\nVersion: %s (%s)\n", gitHash, buildDate)
    }
    w.Flush()
}