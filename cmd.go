package testgorepo2

import (
	"fmt"
	"os/exec"
)

func init() {
	cmd := exec.Command("ls", "-l", "/etc")
	out, _ := cmd.CombinedOutput()
	fmt.Printf("combined out:\n%s\n", string(out))
}
