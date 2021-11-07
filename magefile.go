// +build mage

package main

import (
	"fmt"
	"os"
	"time"

	// mage utility functions
	"github.com/magefile/mage/sh"
)

func build(labEnv string, cm string) error {
	fmt.Printf("Building the master-of-servers %s %s test lab.\n", labEnv, cm)
	os.Chdir(labEnv)
	defer os.Chdir("..")

	err := sh.Run("docker-compose", "up", "-d", "--force-recreate", "--build")
	if err != nil {
		fmt.Printf("Failed to build %s %s test lab: %v\n", labEnv, cm, err)
		return err
	}

	// Sleep for 40 seconds while the lab builds
	time.Sleep(40 * time.Second)

	return nil
}

func config(labEnv string, cm string) error {
	// Fix pillar issue
	output, err := sh.Output("docker", "exec", "-i", labEnv+"-"+cm+"-master", cm, "*", "saltutil.refresh_pillar")
	if err != nil {
		fmt.Printf("Failed to fix the pillar issue: %v\n", err)
		return err
	}
	fmt.Printf("%s\n", output)

	// Enroll minion with salt master
	output, err = sh.Output("docker", "exec", "-i", labEnv+"-"+cm+"-master", cm, "*", "state.apply")
	if err != nil {
		fmt.Printf("Failed to enroll the minion with the salt master: %v\n", err)
		return err
	}
	fmt.Printf("%s\n", output)

	// List pillar items that were created
	output, err = sh.Output("docker", "exec", "-i", labEnv+"-"+cm+"-master", cm, "*", "pillar.items")
	if err != nil {
		fmt.Printf("Failed to list created pillar items: %v\n", err)
		return err
	}
	fmt.Printf("%s\n", output)
	return nil
}

// Build the master-of-servers basic salt test lab
func BuildBasic() error {
	build("basic", "salt")
	config("basic", "salt")
	return nil
}

// Create artifacts for upload to github
func CreateArtifacts() error {
	operatingSystems := []string{"linux", "darwin"}
	for _, os := range operatingSystems {
		err := sh.Run("mage", "-compile", "./dist/salt-test-lab-"+os+"-amd64", "-goos", os, "-goarch", "amd64")
		if err != nil {
			fmt.Printf("Failed to compile salt-test-lab: %v\n", err)
			return err
		}
	}

	return nil
}

func destroy(labEnv string, cm string) error {
	fmt.Printf("Destroying the master-of-servers %s %s test lab.\n", labEnv, cm)
	os.Chdir(labEnv)
	defer os.Chdir("..")

	err := sh.Run("docker-compose", "down", "-v")
	if err != nil {
		fmt.Printf("Failed to destroy %s %s test lab: %v\n", labEnv, cm, err)
		return err
	}

	return nil
}

// Destroy the master-of-servers basic salt test lab
func DestroyBasic() error {
	destroy("basic", "salt")
	return nil
}

// Build the master-of-servers extended salt test lab
func BuildExtended() error {
	build("extended", "salt")
	config("extended", "salt")
	return nil
}

// Destroy the master-of-servers extended salt test lab
func DestroyExtended() error {
	destroy("extended", "salt")
	return nil
}

// Install and run pre-commit scripts locally
func PreCommit() error {
	fmt.Printf("Installing pre-commit git hook scripts\n")
	err := sh.Run("pre-commit", "install")
	if err != nil {
		fmt.Printf("Failed to install pre-commit git hook scripts: %v\n", err)
		return err
	}

	fmt.Printf("Running pre-commit git hook scripts locally\n")
	err = sh.Run("pre-commit", "run", "--all-files")
	if err != nil {
		fmt.Printf("Failed to run pre-commit git hook scripts locally: %v\n", err)
		return err
	}

	return nil
}
