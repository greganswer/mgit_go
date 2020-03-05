package cmd

import (
	"fmt"

	"github.com/greganswer/mgit_go/git"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Used for flags.
var (
	baseBranch    string
	cfgFile       string
	commitMessage string
	issueID       string
	yes           bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "mgit",
	Short:   "run Git work flows for GitHub with issue tracking ticket numbers",
	Version: "0.1.0",
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
