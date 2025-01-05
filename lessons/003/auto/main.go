package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	restore := flag.Bool("r", false, "Restore settings")
	flag.Parse()

	if *restore {
		hideDock(false)
		hideMenu(false)
		return
	}

	err := switchToDesktop2()
	if err != nil {
		fmt.Println("Error switching to Desktop 2:", err)
		return
	}
	time.Sleep(time.Second)
	hideDock(true)
	time.Sleep(time.Second)
	hideMenu(true)
	// Now open VS Code
	// cmd := exec.Command("open", "-n", "/Applications/Visual Studio Code.app")
	cmd := exec.Command("code", "-n")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error starting VS Code:", err)
		return
	}

	almostMaxWindow()
	fmt.Println("Successfully switched to Desktop 2 and started VS Code.")
}

func switchToDesktop2() error {
	applescript := `tell application "System Events"
	key code 19 using {control down}
end tell
    `
	cmd := exec.Command("osascript", "-e", applescript)
	err := cmd.Run()
	return err
}

func hideMenu(hide bool) error {
	v := "true"
	if !hide {
		v = "false"
	}
	applescript := `tell application "System Events" to tell dock preferences to set autohide menu bar to ` + v
	cmd := exec.Command("osascript", "-e", applescript)
	err := cmd.Run()
	return err
}

func hideDock(hide bool) error {
	v := "true"
	if !hide {
		v = "false"
	}
	applescript := `tell application "System Events" to tell dock preferences to set autohide to ` + v
	cmd := exec.Command("osascript", "-e", applescript)
	err := cmd.Run()
	return err
}

func almostMaxWindow() error {
	applescript := `tell application "System Events"
	key code 3 using {shift down, option down, command down}
end tell
    `
	cmd := exec.Command("osascript", "-e", applescript)
	err := cmd.Run()
	return err
}
