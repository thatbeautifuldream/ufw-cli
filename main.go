package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var dryRun bool

// Install UFW if not present
func installUFW() {
	fmt.Println("Checking if UFW is installed...")
	if _, err := exec.Command("ufw", "--version").Output(); err != nil {
		fmt.Println("UFW not found. Installing...")
		if dryRun {
			fmt.Println("[Dry Run] Would execute: sudo apt-get install ufw -y")
			return
		}
		cmd := exec.Command("sudo", "apt-get", "install", "ufw", "-y")
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
		{"ufw", "--force", "enable"},
		{"ufw", "allow", "ssh"},
		{"ufw", "allow", "http"},
		{"ufw", "allow", "https"},
	}
	for _, cmdArgs := range commands {
		if dryRun {
			fmt.Printf("[Dry Run] Would execute: sudo %v\n", strings.Join(cmdArgs, " "))
			continue
		}
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
			if dryRun {
				fmt.Printf("[Dry Run] Would execute: sudo ufw allow %s\n", port)
				continue
			}
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

// Enable/Disable UFW
func toggleUFW(enable bool) {
	action := "disable"
	if enable {
		action = "enable"
	}

	cmdArgs := []string{"ufw"}
	if enable {
		cmdArgs = append(cmdArgs, "--force", "enable")
	} else {
		cmdArgs = append(cmdArgs, "disable")
	}

	if dryRun {
		fmt.Printf("[Dry Run] Would execute: sudo %v\n", strings.Join(cmdArgs, " "))
		return
	}

	cmd := exec.Command("sudo", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to %s UFW: %v\n", action, err)
		return
	}
	fmt.Printf("UFW %sd successfully.\n", action)
}

// Reset UFW
func resetUFW() {
	if dryRun {
		fmt.Println("[Dry Run] Would execute: sudo ufw reset")
		return
	}
	cmd := exec.Command("sudo", "ufw", "reset")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to reset UFW:", err)
		return
	}
	fmt.Println("UFW reset successfully.")
}

// Set default policies
func setDefaultPolicies() {
	if dryRun {
		fmt.Println("[Dry Run] Would execute: sudo ufw default deny incoming")
		fmt.Println("[Dry Run] Would execute: sudo ufw default allow outgoing")
		return
	}
	cmds := [][]string{
		{"ufw", "default", "deny", "incoming"},
		{"ufw", "default", "allow", "outgoing"},
	}
	for _, cmdArgs := range cmds {
		cmd := exec.Command("sudo", cmdArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to set default policy %v: %v\n", cmdArgs, err)
		}
	}
	fmt.Println("Default policies set successfully.")
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ufw-cli",
		Short: "A CLI tool to manage UFW (Uncomplicated Firewall)",
		Long: `ufw-cli helps users install, configure, and manage UFW for their servers.
		
Basic commands:
  install    - Install UFW if not already present
  setup      - Configure basic UFW rules (SSH, HTTP, HTTPS)
  configure  - Add custom port rules
  status     - Show current UFW status
  enable     - Enable the UFW firewall
  disable    - Disable the UFW firewall
  reset      - Reset UFW to default settings
  default    - Set default policies for incoming and outgoing traffic

Example usage:
  ufw-cli install     # Install UFW
  ufw-cli setup       # Set up basic rules
  ufw-cli configure   # Add custom ports
  ufw-cli enable      # Turn on the firewall
  ufw-cli status      # Check firewall status`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to ufw-cli! Use --help to see available commands.")
		},
	}

	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Simulate the command without executing")

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

	// Enable Command
	var enableCmd = &cobra.Command{
		Use:   "enable",
		Short: "Enable UFW",
		Run: func(cmd *cobra.Command, args []string) {
			toggleUFW(true)
		},
	}

	// Disable Command
	var disableCmd = &cobra.Command{
		Use:   "disable",
		Short: "Disable UFW",
		Run: func(cmd *cobra.Command, args []string) {
			toggleUFW(false)
		},
	}

	// Reset Command
	var resetCmd = &cobra.Command{
		Use:   "reset",
		Short: "Reset UFW to default settings",
		Run: func(cmd *cobra.Command, args []string) {
			resetUFW()
		},
	}

	// Default Policies Command
	var defaultCmd = &cobra.Command{
		Use:   "default",
		Short: "Set default policies for incoming and outgoing traffic",
		Run: func(cmd *cobra.Command, args []string) {
			setDefaultPolicies()
		},
	}

	// Add commands to root
	rootCmd.AddCommand(installCmd, setupCmd, configureCmd, statusCmd, enableCmd, disableCmd, resetCmd, defaultCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}