package logreader

import (
	"os"
	"os/exec"
)

var (
	DEFAULT_PATCH = "./"
)

func getLogData(filename, updatedPath string) (string,error) {
	var path = DEFAULT_PATCH
	if updatedPath != "" {
		path = updatedPath
	}

	cmd := exec.Command("tail", path+filename)
	cmd.Stdin = os.Stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
	return cmd.String(),nil
}

func GetLogDataOnNewPatch(filename, updatedPath string) (string,error) {
	return getLogData(filename, updatedPath)
}

func GetLogData(filename string) (string,error) {
	return getLogData(filename, "")
}
