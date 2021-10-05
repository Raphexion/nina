package main

import (
	"nina/backend"
	"nina/cmd"
	"nina/conf"
)

func main() {
	back := &backend.RealBackend{}
	conf.SetBackend(back)

	cmd.Execute()
}
