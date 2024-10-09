package main

import (
	"encoding/json"
	"log/slog"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	n.Handle("echo", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			slog.Error("Could not unmarshal message", "err", err)
			return err
		}

		body["type"] = "echo_ok"
		return n.Reply(msg, body)
	})

	slog.Info("Starting node")
	if err := n.Run(); err != nil {
		slog.Error("Failure while running node", "err", err)
	}
}
