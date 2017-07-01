package plugin

import (
	"fmt"
	"os"
)

// function ID
type FuncID int

const (
	FUNC_TYPE FuncID = iota
	FUNC_CLOSE
	FUNC_HELLO
	FUNC_HELLO2
)

// loaded plugin list
var plugins = map[string]*Plugin{}

//
// Plugin
//
type Plugin struct {
	cmd *Cmd
}

func (o *Plugin) Type() (string, error) {
	rets, err := o.call(FUNC_TYPE, []byte{})
	if err != nil {
		return "", err
	}
	return string(rets), nil
}

func (o *Plugin) Hello() (string, error) {
	rets, err := o.call(FUNC_HELLO, []byte{})
	if err != nil {
		return "", err
	}
	return string(rets), nil
}

func (o *Plugin) Hello2(str string) (string, error) {
	rets, err := o.call(FUNC_HELLO2, []byte(str))
	if err != nil {
		return "", err
	}
	return string(rets), nil
}

func (o *Plugin) close() error {
	if err := o.cmd.SendID(FUNC_CLOSE); err != nil {
		return err
	}
	if err := o.cmd.SendArgs([]byte{}); err != nil {
		return err
	}
	return nil
}

func (o *Plugin) call(id FuncID, args []byte) ([]byte, error) {
	if err := o.cmd.SendID(id); err != nil {
		return nil, err
	}
	if err := o.cmd.SendArgs(args); err != nil {
		return nil, err
	}
	return o.cmd.RecvReturn()
}

//
// User Interface
//
func Load(path string) error {
	cmd, err := NewCmd(path)
	if err != nil {
		return err
	}

	pl := &Plugin{cmd: cmd}
	t, err := pl.Type()
	if err != nil {
		pl.close()
		return err
	}

	plugins[t] = pl
	return nil
}

func Release() {
	for _, pl := range plugins {
		pl.close()
	}
	plugins = map[string]*Plugin{}
}

func Get(tag string) (*Plugin, error) {
	if pl, ok := plugins[tag]; ok {
		return pl, nil
	} else {
		return nil, fmt.Errorf("Not Loaded %s", tag)
	}
}

//
// vendor interface
//
func DispatchLoop(adapt HelloAdapter) error {
	cmd, err := NewCmd2(os.Stdin, os.Stdout)
	if err != nil {
		return err
	}

	for {
		funcID, err := cmd.RecvID()
		if err != nil {
			return err
		}
		args, err := cmd.RecvArgs()
		if err != nil {
			return err
		}

		switch funcID {
		case FUNC_TYPE:
			rets := adapt.Type()
			err := cmd.SendReturn([]byte(rets))
			if err != nil {
				return err
			}

		case FUNC_HELLO:
			rets := adapt.Hello()
			err := cmd.SendReturn([]byte(rets))
			if err != nil {
				return err
			}

		case FUNC_HELLO2:
			str := string(args)
			rets := adapt.Hello2(str)
			err := cmd.SendReturn([]byte(rets))
			if err != nil {
				return err
			}

		case FUNC_CLOSE:
			return nil
		}
	}
}

//
// Adapter
//
type HelloAdapter interface {
	Type() string
	Hello() string
	Hello2(string) string
}
