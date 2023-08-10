package main

import (
	"fmt"
	"os"
)

//doing this as renaming a file by fielsystem is atomic in nature
func saveFile(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, 1)
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}

	err = fp.Sync() // fsync
	if err != nil {
		os.Remove(tmp)
		return err
	}

	return os.Rename(tmp, path)
}
