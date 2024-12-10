package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// Install UFW if not present
func installUFW() {
	fmt.Println("Checking if UFW is installed...")
	if _, err := exec.Command("ufw", "--version").Output(); err != nil {
		fmt.Println("UFW not found. Installing...")
		cmd := exec.Command("sudo", "apt-get", "install", "ufw", "-y") // Assuming Debian-based systems
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Failed to install UFW:", err)
			return
		}
		fmt.Println("UFW installed successfully.")
	} else {
		fmt.Println("UFW is already installed.")
	}
}

// Set up basic UFW rules
func setupUFW() {
	fmt.Println("Setting up basic UFW rules...")
	commands := [][]string{
		{"ufw", "enable"},
		{"ufw", "allow", "ssh"},
		{"ufw", "allow", "http"},
		{"ufw", "allow", "https"},
	}
	for _, cmdArgs := range commands {
		cmd := exec.Command("sudo", cmdArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to execute %v: %v\n", cmdArgs, err)
		}
	}
	fmt.Println("Basic UFW setup complete.")
}

// Prompt user for additional ports
func configureAdditionalPorts() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter additional ports to open (comma-separated, e.g., 8080,3306):")
	input, _ := reader.ReadString('\n')
	ports := strings.Split(strings.TrimSpace(input), ",")
	for _, port := range ports {
		port = strings.TrimSpace(port)
		if port != "" {
			cmd := exec.Command("sudo", "ufw", "allow", port)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf("Failed to allow port %s: %v\n", port, err)
			}
		}
	}
	fmt.Println("Custom ports configured.")
}

// Show UFW status
func showUFWStatus() {
	cmd := exec.Command("sudo", "ufw", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to retrieve UFW status:", err)
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ufw-cli",
		Short: "A CLI tool to manage UFW (Uncomplicated Firewall)",
		Long:  "ufw-cli helps users install, configure, and manage UFW for their servers.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to ufw-cli! Use --help to see available commands.")
		},
	}

	// Install Command
	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "Check and install UFW if not present",
		Run: func(cmd *cobra.Command, args []string) {
			installUFW()
		},
	}

	// Setup Command
	var setupCmd = &cobra.Command{
		Use:   "setup",
		Short: "Set up basic UFW rules",
		Run: func(cmd *cobra.Command, args []string) {
			setupUFW()
		},
	}

	// Configure Command
	var configureCmd = &cobra.Command{
		Use:   "configure",
		Short: "Configure additional UFW rules",
		Run: func(cmd *cobra.Command, args []string) {
			configureAdditionalPorts()
		},
	}

	// Status Command
	var statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Show UFW status",
		Run: func(cmd *cobra.Command, args []string) {
			showUFWStatus()
		},
	}

	// Add commands to root
	rootCmd.AddCommand(installCmd, setupCmd, configureCmd, statusCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
