package helper

import "os"

// FileExists Checks whether a file or directory exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil || !os.IsNotExist(err)
}
