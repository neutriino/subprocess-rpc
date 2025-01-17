package subrpc

import (
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/neutriino/libatc"
)

// ProcessOptions allows for passing process options to NewProcess
type ProcessOptions struct {
	Name         string
	Type         string
	Config       map[string]interface{}
	Handler      interface{}
	ExePath      string
	Socket       string
	Env          map[string]string
	Token        string
	StartupDelay time.Duration
}

// ProcessInfo holds information about running processes
type ProcessInfo struct {
	Name         string
	Type         string
	Config       map[string]interface{}
	Token        string
	Handler      interface{}
	CMD          *libatc.Subprocess
	Options      ProcessOptions
	Running      bool
	Terminate    chan bool
	PID          int
	Socket       string
	RPC          *rpc.Client
	StartupDelay time.Duration
}

// ProcessInput type
type ProcessInput struct {
	Socket       string
	ServerSocket string
	Token        string
	Config       []byte
}
