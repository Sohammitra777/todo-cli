package utils

import "os"

func CheckAndCreateStorageFile(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		f, err := os.Create(file)
		PrintErr(err)
		defer f.Close()
	}
}