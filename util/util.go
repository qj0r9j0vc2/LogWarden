package util

import (
	"os"
	"os/exec"
)

func Clear() {
	c := exec.Command("Clear")
	c.Stdout = os.Stdout
	err := c.Run()
	if err != nil {
		return
	}
}
