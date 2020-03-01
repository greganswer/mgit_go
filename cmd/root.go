package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// TODO: Find a way to remove these variables.
var commitMessage, baseBranch, issueID string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "mgit",
	Short:   "run Git work flows for GitHub with issue tracking ticket numbers",
	Version: "0.0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	FailIfError(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mgit.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		FailIfError(err)

		// Search config in home directory with name ".mgit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mgit")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// TODO: Extract the following helper functions

// FailIfError exits and returns a standardized error message if an error occurred.
func FailIfError(err error) {
	if err == nil {
		return
	}
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	fmt.Println("")
	fmt.Println(red("FAIL:"), err)
	os.Exit(1)
}

// Confirm returns true if the user confirms.
func Confirm(message string) bool {
	prompt := promptui.Prompt{
		Label:     message,
		IsConfirm: true,
	}
	_, err := prompt.Run()
	return err == nil
}

// ConfirmOrAbort displays a prompt and exits if the user selects no.
func ConfirmOrAbort(message, abortMessage string) {
	if !Confirm(message) {
		fmt.Println(abortMessage)
		os.Exit(0)
	}
}

func info(message string) string {
	return color.CyanString(message)
}

func success(message string) string {
	return color.GreenString(message)
}

func finished() {
	c := color.New(color.FgGreen, color.Bold)
	c.Println("OK")
	fmt.Println()
}
