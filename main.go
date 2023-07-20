package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Parse command-line flags
	altRoot := flag.String("root", "", "Set an alternative Drupal root")
	flag.Parse()

	var drupalRoot string

	// Use the alternative Drupal root if provided
	if *altRoot != "" {
		drupalRoot = *altRoot
	} else {
		// If no alternative root provided, find the Drupal root from the current directory
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			os.Exit(1)
		}
		drupalRoot, err = drushlauncher.FindDrupalRoot(cwd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Construct the full command to run drush
	drushCmd := exec.Command(filepath.Join(drupalRoot, "vendor", "bin", "drush"), flag.Args()...)

	// Pass the current environment variables to the drush command
	drushCmd.Env = os.Environ()

	// Set the correct working directory for the drush command
	drushCmd.Dir = drupalRoot

	// Redirect standard input/output/error for the drush command
	drushCmd.Stdin = os.Stdin
	drushCmd.Stdout = os.Stdout
	drushCmd.Stderr = os.Stderr

	// Run the drush command
	if err := drushCmd.Run(); err != nil {
		fmt.Println("Error executing drush:", err)
		os.Exit(1)
	}
}
