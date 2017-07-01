package main

import (
	"fmt"

	"github.com/hidez8891/go-pipe-plugin-sample/plugin"
)

type Hello struct {
}

func (o *Hello) Type() string {
	return "tag_hello"
}

func (o *Hello) Hello() string {
	return fmt.Sprintf("Hello!")
}

func (o *Hello) Hello2(str string) string {
	return fmt.Sprintf("Hello %s!", str)
}

func main() {
	err := plugin.DispatchLoop(&Hello{})
	if err != nil {
		panic(err)
	}
}
