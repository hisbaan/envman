package utils

import (
	"os"
)

func IsSymlinkOrDoesNotExist(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return true
		}
		return false
	}
	return fi.Mode()&os.ModeSymlink != 0
}
