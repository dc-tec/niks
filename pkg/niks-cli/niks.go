package nikscli

import (
	"fmt"
	"os/exec"
)

func Update(config string, dryRun bool) error {
	cmdArgs := []string{"switch", "--flake .#", config}

	if dryRun {
		cmdArgs = append(cmdArgs, "--dry-run")
	}

	cmd := exec.Command("nixos-rebuild", cmdArgs...)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Failed to update system: %w", err)
	}

	return nil
}
