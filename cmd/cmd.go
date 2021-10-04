package cmd

import (
	"bufio"
	"fmt"
	"log"
	"nina/backend"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	back backend.Backend

	rootCmd = &cobra.Command{
		Use:   "nina",
		Short: "CLI to interact with Noko time tracking",
		Long:  "A commandline client written in golang to help interact with Noko time tracker",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/nina.yaml)")

	// Use the real noko backend by default
	back = &backend.RealBackend{}

	// Add the individual commands
	rootCmd.AddCommand(NewTimerCmd(back))
	rootCmd.AddCommand(NewProjectCmd(back))
	rootCmd.AddCommand(NewEntryCmd(back))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("nina.yaml")
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()

	back.Init()
}

func promptForConfirmation(text string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", text)

		resp, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		resp = strings.ToLower(strings.TrimSpace(resp))

		if resp == "yes" || resp == "y" {
			return true
		} else if resp == "no" || resp == "n" {
			return false
		}
	}
}
