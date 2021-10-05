package cmd

import (
	"nina/backend"
	"nina/conf"

	"github.com/spf13/cobra"
)

func BackendRunCmd(f func(backend.Backend)) func(*cobra.Command, []string) {
	return func(*cobra.Command, []string) {
		back := conf.GetBackend()
		back.Init()

		f(conf.GetBackend())
	}
}
