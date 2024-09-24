package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {

	const file_path = "/var/tmp/count.txt"
	file, err := os.OpenFile(file_path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lock_err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	if lock_err != nil {
		log.Fatal(lock_err)
	}
	defer syscall.Flock(int(file.Fd()), syscall.LOCK_UN)

	fmt.Print("acquired lock")
	fmt.Println(file.Name())

	fmt.Print("runner done")
}
