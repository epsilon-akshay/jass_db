package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Print()
	fp, err := os.OpenFile("testfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	buf := []byte("line2")
	buf = append(buf, '\n')
	_, err = fp.Write(buf)
	if err != nil {
		panic(err)
	}
	err = fp.Sync()
	if err != nil {
		panic(err)
	}
	fp.Close()
}
