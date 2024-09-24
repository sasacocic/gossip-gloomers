package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"strconv"
	"syscall"

	"example.com/greetings"
	"example.com/hello/nest"
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

	fmt.Println(file.Name())
	file_info, _ := file.Stat()
	b := make([]byte, file_info.Size())
	amount, err_read := file.Read(b)
	if err_read != nil {
		log.Fatal(err_read)
	}
	content := strings.TrimSpace(string(b))
	fmt.Println("amount: ", amount, "content: '", content, "'")

	prev_int, _ := strconv.Atoi(content)
	fmt.Println("prev_val:", prev_int)
	prev_int = prev_int + 1
	fmt.Println("new_val:", prev_int)
	fmt.Println("what is written: ", string([]byte(strconv.Itoa(prev_int))))
	file.WriteAt([]byte(strconv.Itoa(prev_int)), 0)
	// file.Sync()

	fmt.Println("programming completing")

	return

	numss := []int{1, 2, 3, 4}

	for ind, val := range numss {
		fmt.Println(ind, val)
	}

	two := [][]int{}
	two = append(two, []int{9, 10, 11})

	nums := [][]int{[]int{99}, []int{2}}
	slices.SortFunc(nums, func(a, b []int) int {
		return a[0] - b[0]
	})

	message := greetings.Hello("Gladys")
	fmt.Println(message)
	fmt.Println(wowzer())
	fmt.Println(nest.Nester())

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	var s int = 1
	for s < 10 {
		fmt.Println(s)
		s += 1
	}

	var strs [10]string
	for i := 0; i < cap(strs); i++ {
		strs[i] = "wowzers"
	}
	strs_two := [3]int{1, 2, 3}
	fmt.Println(len(strs_two))
	fmt.Println(strs)

	var i int = 1

	for i < 10 {
		fmt.Println(i)
		i += 1
	}

	var (
		age  int    = 99
		name string = "bob"
	)
	fmt.Println(age, name)

}
