package testgorepo2

import (
	"fmt"
	"os/exec"
)

func _11() {
	cmd := exec.Command("ls", "-l", "/var/log/")
	out, _ := cmd.CombinedOutput()
	fmt.Printf("combined out:\n%s\n", string(out))
}
