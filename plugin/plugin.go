package plugin

/*
  // user interface
  func Load(path string) error {
		cmd, err := NewCmd(path)
		pl := &Plugin{cmd: cmd}

		type, err := pl.Type()
		plugins[type] = pl
  }

  func Release() {
		for _, pl := range plugins {
			pl.close()
		}
  }

  func Get(tag string) (*Plugin, error) {
		return plugins[tag]
  }


	// Plugin
	type Plugin struct {
		cmd ???
	}

	func (o *Plugin) Type() (string, error) {
		err := cmd.SendID(FUNC_TYPE)
		err := cmd.SendArgs([]byte{})
		rets, err := cmd.RecvReturn()
		return string(rets), nil
	}

	func (o *Plugin) Hello() (string, error) {
		err := cmd.SendID(FUNC_HELLO)
		err := cmd.SendArgs([]byte{})
		rets, err := cmd.RecvReturn()
		return string(rets), nil
	}

	func (o *Plugin) Hello2(str string) (string, error) {
		err := cmd.SendID(FUNC_HELLO2)
		err := cmd.SendArgs([]byte{str})
		rets, err := cmd.RecvReturn()
		return string(rets), nil
	}

	func (o *Plugin) close() error {
		err := cmd.SendID(FUNC_CLOSE)
		err := cmd.SendArgs([]byte{})
		return nil
	}


	// vendor interface
	func DispatchLoop(adapt *HelloAdapter) error {
		cmd, err := NewCmd2(os.Stdin, os.Stdout)

		for {
			funcID, err := cmd.RecvID()
			args, err   := cmd.RecvArgs()

			switch (funcID) {
			case FUNC_TYPE:
				rets := adapt.Type()
				err := cmd.SendReturn([]byte(rets))

			case FUNC_HELLO:
				rets := adapt.Hello()
				err := cmd.SendReturn([]byte(rets))

			case FUNC_HELLO2:
				str := string(args)
				rets := adapt.Hello2(str)
				err := cmd.SendReturn([]byte(rets))

			case FUNC_CLOSE:
				return nil
			}
		}
	}


	// Adapter
	type HelloAdapter interface {
		Type() string
		Hello() string
		Hello2(string) string
	}
*/
