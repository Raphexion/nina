package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

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

	// Add the individual commands
	rootCmd.AddCommand(NewTimerCmd())
	rootCmd.AddCommand(NewProjectCmd())
	rootCmd.AddCommand(NewEntryCmd())
	rootCmd.AddCommand(NewWorkdayCmd())
	rootCmd.AddCommand(NewWorkweekCmd())
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
