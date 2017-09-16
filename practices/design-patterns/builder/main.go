package main

import (
	"encoding/json"
	"fmt"
)

type Message interface {
	Encode() ([]byte, error)
	Text() (string, error)
}

type jsonMsg map[string]interface{}

func (j jsonMsg) Encode() ([]byte, error) {
	return json.Marshal(j)
}

func (j jsonMsg) Text() (string, error) {
	b, err := json.Marshal(j)

	if err != nil {
		return "", err
	}
	return string(b), nil
}

type stringMsg string

func (str stringMsg) Encode() ([]byte, error) {
	return []byte(str), nil
}

func (str stringMsg) Text() (string, error) {
	return string(str), nil
}

// The Builder, which hide the complex implementation of
// How to encode the message
func encode(m Message) ([]byte, error) {
	b, err := m.Encode()
	return b, err
}

func send(b []byte) {
	fmt.Println("Sending the bytes", b)
}

func main() {
	var msg1 jsonMsg = map[string]interface{}{
		"ping": "pong",
		"foo":  "bar",
	}
	var msg2 stringMsg = "This is test string"

	// Use the encode() as Builder
	if b, err := encode(msg1); err == nil {
		send(b)
	}
	if b, err := encode(msg2); err == nil {
		send(b)
	}
}
