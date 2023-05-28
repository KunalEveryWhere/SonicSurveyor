package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func main() {
	// Command to run with sudo
	cmd := exec.Command("sudo", "echo", "Hello, World!")

	// Set the appropriate options to run with elevated privileges
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: 0, // UID of the superuser (root)
			Gid: 0, // GID of the superuser (root)
		},
	}

	// Run the command and get the output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output
	fmt.Println("Command output:", string(output))
}
