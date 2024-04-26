package nikscli

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
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

func ListGenerations() error {

	cmd := exec.Command("nix-env", "--list-generations")

	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("We have some issues getting the generations: %v", err)
	}

	fmt.Println(string(output))

	return nil

}

func CleanGenerations(generations []int) error {
	var cmdArgs []string

	if generations != nil {
		cmdArgs = append(cmdArgs, "--generations")

		for _, gen := range generations {
			cmdArgs = append(cmdArgs, strconv.Itoa(gen))
		}
	} else {

		cmdArgs = []string{"--delete-old"}
		fmt.Println(cmdArgs)
	}

	cmd := exec.Command("nix-collect-garbage", cmdArgs...)
	fmt.Println(cmd)
	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("Failed to run garbage collection: %w", err)
	}

	return nil
}
