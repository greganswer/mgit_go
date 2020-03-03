package cmd

import (
	"fmt"
	"os"

	"github.com/greganswer/mgit_go/git"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// Used for flags.
var (
	baseBranch    string
	commitMessage string
	issueID       string
	yes           bool
)

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
	baseBranch = git.DefaultBaseBranch()
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mgit.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&yes, "yes", "y", false, "assume \"yes\" as answer to all prompts and run non-interactively")
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

// FailOrOK displays a failure message if there's an error otherwise displays "OK".
func FailOrOK(err error) {
	FailIfError(err)
	finished()
}

// Confirm returns true if the user confirms or if the yes flag is present.
func Confirm(message string) bool {
	if yes {
		return true
	}
	prompt := promptui.Prompt{
		Label:     message,
		IsConfirm: true,
	}
	_, err := prompt.Run()
	return err == nil
}

func skip(messages ...interface{}) {
	fmt.Println(color.YellowString("SKIP:"), fmt.Sprintf(messages[0].(string), messages[1:]...))
	fmt.Println("")
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
