package cmd

import (
	"nina/backend"

	"github.com/spf13/cobra"
)

func BackendRunCmd(m backend.Backend, f func(backend.Backend)) func(*cobra.Command, []string) {
	return func(*cobra.Command, []string) {
		f(m)
	}
}
