package nikscli

import (
	"fmt"
	"os/exec"
)

func Update(path string, config string) error {
	cmdArgs := []string{"switch", "--flake", path + "#" + config}

	cmd := exec.Command("nixos-rebuild", cmdArgs...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to update system: %w", err)
	}
	fmt.Printf("NixOS Rebuild Output: \n%s\n", string(out))

	return nil
}
