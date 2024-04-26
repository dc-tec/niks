package nikscli

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Update(path string, config string) error {
	cmdArgs := []string{"nixos-rebuild", "switch", "--flake", path + "#" + config}

	cmd := exec.Command("sudo", "-n", "true")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		cmd = exec.Command("doas", cmdArgs...)
		fmt.Printf("command %v\n", cmd)
	} else {
		cmd = exec.Command("sudo", cmdArgs...)
	}

	cmd.Stdout = io.MultiWriter(os.Stdout, &stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Failed to start the update command: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		if strings.Contains(stderr.String(), "permission denied") {
			return fmt.Errorf("Permission Denied, you do not have the correct privileges: %w", err)
		}
		return fmt.Errorf("Command execution failed: %w", err)
	}

	return nil
}

func listGenerations() error {

	cmd := exec.Command("nix-env", "--list-generations")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Stdout = io.MultiWriter(os.Stdout, &stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

	if err := cmd.Start(); err != nil {
		fmt.Printf("We have some issues getting the generations: %v", err)
	}

	return nil

}
