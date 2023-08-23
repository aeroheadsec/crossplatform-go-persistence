package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("By using this script created by https://github.com/dharmade")
	fmt.Println("You swear and agree that you will not use it for any malicious use")
	fmt.Println("Go to the github repository for the full ToS.\n")
	fmt.Println("Copying to startup....")
	copyToStartup()
}

func copyToStartup() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	startupPath := ""
	switch runtime.GOOS {
	case "windows":
		startupPath = filepath.Join(usr.HomeDir, "AppData", "Roaming", "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
	case "darwin":
		startupPath = filepath.Join(usr.HomeDir, "Library", "LaunchAgents")
	case "linux":
		startupPath = filepath.Join(usr.HomeDir, ".config", "autostart")
	default:
		fmt.Println("Unsupported operating system.")
		return
	}

	destPath := filepath.Join(startupPath, filepath.Base(exePath))
	if destPath == exePath {
		fmt.Println("Script is already in the startup folder.")
		return
	}

	err = copyFile(exePath, destPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func copyFile(srcPath, destPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	return nil
}
