package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"os"
)

func main() {
	fsm := NewSimpleFSM()

	fmt.Println(fsm.Apply("SET foo bar"))
}

// SImple FSM represents a finite state machine that stores simple key value pairs
type SimpleFSM struct {
	state map[string]string
}

type Snapshot struct {
	state map[string]string
}

// new simple FSM creates and initializes a new FSM
func NewSimpleFSM() *SimpleFSM {
	return &SimpleFSM{state: make(map[string]string)}
}

// apply processes a command and updates the state
func (fsm *SimpleFSM) Apply(command string) interface{} {
	parts := strings.Split(command, " ")
	switch parts[0] {
	case "SET":
		key, value := parts[1], parts[2]
		fsm.state[key] = value
		return nil
	case "GET":
		key := parts[1]
		return fsm.state[key]
	case "DELETE":
		key := parts[1]
		delete(fsm.state, key)
		return nil
	default:
		return fmt.Errorf("unknown command: %s", command)
	}
}

// Snapshot() type creates a snapshot of the FSM state
func (fsm *SimpleFSM) Snapshot() (*Snapshot, error) {
	return &Snapshot{state: fsm.state}, nil
}

// Persist writes the snapshot to a file
func (snap *Snapshot) Persist(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(snap.state)
}
