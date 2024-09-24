package main

import (
	"encoding/json"
	"log"
	"strings"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func tt(name string) (one string, two string) {
	things := strings.Split(name, " ")
	one = things[0]
	two = things[1]
	return
}

func main() {

	n := maelstrom.NewNode()
	n.Handle("echo", func(msg maelstrom.Message) error {

		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "echo_ok"

		return n.Reply(msg, body)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
