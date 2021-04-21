package util

import (
	"os/exec"
	"testing"
)

func TestFoo(t *testing.T) {
	mp := make(map[string]string)

	mp["1"] = "1"

	//t.Log(mp["0"])
	//t.Log(strconv.Atoi(""))

	cmd := exec.Command("vim", "README.md", "+5")
	cmd.Run()
}
