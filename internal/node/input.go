package node

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Input struct {
	buf *bufio.Scanner
}

func NewInput(rd io.Reader) *Input {
	return &Input{buf: bufio.NewScanner(rd)}
}

func (i *Input) ReadMessage(v any) error {
	if i.buf.Scan() {
		bt := i.buf.Bytes()
		if len(bt) == 0 {
			return errors.New("Empty message")
		}
		if err := json.Unmarshal(bt, v); err != nil {
			return fmt.Errorf("Failed to read message: %s", err)
		}
		return nil
	}
	return io.EOF
}
