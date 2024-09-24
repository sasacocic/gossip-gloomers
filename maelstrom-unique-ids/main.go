package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type JsonData struct {
	Src   string         `json:"src"`
	Dest  string         `json:"dest"`
	Body  map[string]any `json:"body"`
	Count int            `json:"count"`
}

func helper() int {

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

	file_info, _ := file.Stat()
	b := make([]byte, file_info.Size())
	_, err_read := file.Read(b)
	if err_read != nil {
		log.Fatal(err_read)
	}
	content := strings.TrimSpace(string(b))

	prev_int, _ := strconv.Atoi(content)
	prev_int = prev_int + 1
	file.WriteAt([]byte(strconv.Itoa(prev_int)), 0)
	return prev_int

}

func main() {

	// data, err := os.ReadFile("/tmp/var/output.json")

	// if err != nil {
	// 	fmt.Println("this is a problem")
	// 	log.Fatal(err)
	// }

	// json_data := JsonData{}
	// json.Unmarshal(data, &json_data)
	// fmt.Println(json_data)

	// json_data.Count += 1

	// encoded, err := json.Marshal(json_data)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// writing the encoded bytes seems to just write it as utf-8? well that doesn't make any sense.
	// it's encoded. I encoded it into some bytes. So it's just writing those bytes, and my
	// editor is reading it as utf-8
	// errr := os.WriteFile("output.json", encoded, 0666)

	// if errr != nil {
	// 	log.Fatal(errr)
	// }

	// data, err := os.ReadFile(file)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("data read")
	// count, err := strconv.Atoi(string(data))
	// fmt.Println(count)
	// fmt.Println("data read")

	n := maelstrom.NewNode()

	// counter := 0

	n.Handle("generate", func(msg maelstrom.Message) error {

		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "generate_ok"
		nn := helper()
		body["id"] = nn

		return n.Reply(msg, body)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
