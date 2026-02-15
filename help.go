package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

// This variable is populated by build.sh at compile time
var buildDate = "Unknown"

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
	{"general", "dnsreset", "When websites are not loading in browser"},
	{"docker", "ps-all", "List all containers including stopped ones"},
	{"go", "build-macos", "go build -o myapp help.go"},
	{"go", "build-linux", "GOOS=linux GOARCH=amd64 go build -o myapp-linux help.go"},
	{"go", "build-pi", "GOOS=linux GOARCH=arm64 go build -o myapp-pi64 help.go"},
}

func main() {

	// 1. Version Check
    if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version" || os.Args[1] == "version") {
        fmt.Printf("Help Tool v1.0\n")
        fmt.Printf("Build Date: %s\n", header(buildDate))
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

	// Add a footer if showing all scripts
    if filter == "" {
        fmt.Fprintln(w, dim("\nBinary build date: "+buildDate))
    }
    w.Flush()
}