package main

import (
	"fmt"

	"github.com/hidez8891/go-pipe-plugin-sample/plugin"
)

func main() {
	err := plugin.Load("./sample/hello/hello.exe")
	if err != nil {
		panic(err)
	}
	defer plugin.Release()

	pl, err := plugin.Get("tag_hello")
	if err != nil {
		panic(err)
	}

	expect := "Hello!"
	str, err := pl.Hello()
	if err != nil {
		panic(err)
	}
	if str != expect {
		panic(fmt.Errorf("ERR: get %s, want %s", str, expect))
	}

	expect = "Hello world!"
	str, err = pl.Hello2("world")
	if err != nil {
		panic(err)
	}
	if str != expect {
		panic(fmt.Errorf("ERR: get %s, want %s", str, expect))
	}

	fmt.Println("sample is successed")
}
