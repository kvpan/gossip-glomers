package node_test

import (
	"io"
	"strings"
	"testing"

	"github.com/kvpan/gossip-glomers/internal/node"
	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	type testMessage struct {
		Type  string `json:"type"`
		Value int    `json:"value"`
	}

	t.Run("Invalid message", func(t *testing.T) {
		var msg testMessage
		in := node.NewInput(strings.NewReader("invalid"))
		err := in.ReadMessage(&msg)
		assert.Error(t, err)
	})

	t.Run("Empty message", func(t *testing.T) {
		var msg testMessage
		in := node.NewInput(strings.NewReader("\n"))
		err := in.ReadMessage(&msg)
		assert.Error(t, err)
	})

	t.Run("Single message", func(t *testing.T) {
		var msg testMessage
		in := node.NewInput(strings.NewReader(`{"type":"test","value":1}`))
		err := in.ReadMessage(&msg)
		assert.Equal(t, testMessage{Type: "test", Value: 1}, msg)
		assert.NoError(t, err)
	})

	t.Run("Multiple messages", func(t *testing.T) {
		in := node.NewInput(strings.NewReader(`{"type":"test","value":1}` + "\n" + `{"type":"test","value":2}`))
		var messages []testMessage
		for {
			var msg testMessage
			err := in.ReadMessage(&msg)
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			messages = append(messages, msg)
		}

		expected := []testMessage{
			{Type: "test", Value: 1},
			{Type: "test", Value: 2.},
		}
		assert.Equal(t, expected, messages)
	})
}
