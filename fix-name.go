package main


import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)


func RunCmd(cmd string, args ...string) bool {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		fmt.Printf("OUTPUT: %s\n", out)
		return false
	}
	return true
}


func main() {
	nshelper := "C:\\Program Files (x86)\\NPVR\\NScriptHelper.exe"

	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "FILENAME")
		os.Exit(1)
	}
	origFilename := os.Args[1]

	re := regexp.MustCompile(`_(\d{8})_(\d{8})\.ts$`)
	newFilename := re.ReplaceAllString(origFilename, " - S${1}E${2}.ts")
	fmt.Println("New filename will be", newFilename)
	success := true
	if origFilename != newFilename {
		success =
			RunCmd("cmd", "/C", "move", origFilename, newFilename) &&
			RunCmd(nshelper, "-Rename", origFilename, newFilename)
	}
	if success {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

