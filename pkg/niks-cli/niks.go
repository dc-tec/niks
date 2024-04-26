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

	cmd := exec.Command("sudo", cmdArgs...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	fmt.Printf("out: %s\n", err)

	if err != nil {
		cmd = exec.Command("doas", cmdArgs...)
	} else {
		cmd = exec.Command("sudo", cmdArgs...)
	}

	cmd.Stdout = io.MultiWriter(os.Stdout, &stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

	if err := cmd.Start(); err != nil {
		if strings.Contains(stderr.String(), "permission denied") {
			return fmt.Errorf("Failed to start the update command: %w", err)
		}
		return fmt.Errorf("Command execution failed: %w", err)
	}

	return nil
}
