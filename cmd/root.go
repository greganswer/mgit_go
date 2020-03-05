package cmd

import (
	"fmt"
	"os"

	"github.com/greganswer/mgit_go/git"
	"github.com/greganswer/mgit_go/issues"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
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

// TODO: Extract the following helper functions

// Fail exits the program with a standardized error message.
func Fail(err error) {
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	fmt.Println(red("FAIL:"), err)
	os.Exit(1)
}

// FailIfError exits the program with a standardized error message if an error occurred.
func FailIfError(err error) {
	if err != nil {
		Fail(err)
	}
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
	message := fmt.Sprintf(messages[0].(string), messages[1:]...)
	fmt.Println(color.YellowString("SKIP:"), message)
	fmt.Println("")
}

func warn(messages ...interface{}) {
	message := fmt.Sprintf(messages[0].(string), messages[1:]...)
	fmt.Println(color.YellowString("WARN:"), message)
}

func emphasis(message string) string {
	return color.CyanString(message)
}

func finished() {
	c := color.New(color.FgGreen, color.Bold)
	c.Println("OK")
	fmt.Println()
}

// getCommitMessage get a commit message from the branch or from user input.
func getCommitMessage(branchName string) string {
	issue, err := issues.FromBranch(branchName)
	switch err.(type) {
	case *issues.BranchHasNoIDError:
		warn(err.Error())
	default:
		Fail(err)
	}
	return issue.String()
}

func addAllFiles() {
	fmt.Println("Adding all files...")
	FailOrOK(git.AddAll())
}

func commitChanges(message string) {
	fmt.Println("Committing files...")
	FailOrOK(git.Commit(message))
}

func pushChanges(branch string) {
	fmt.Printf("Pushing changes on %s branch to remote...\n", emphasis(branch))
	FailOrOK(git.Push(branch))
}

func getIssueFromCurrentBranch() issues.Issue {
	currentBranch, err := git.CurrentBranch()
	FailIfError(err)
	i, err := issues.FromBranch(currentBranch)
	FailIfError(err)
	return i
}

func getIssueFromTracker(issueID string) issues.Issue {
	fmt.Printf("Retrieving issue data from %s...\n", emphasis(issues.APIURL(issueID)))
	i, err := issues.FromTracker(issueID)
	FailOrOK(err)
	return i
}
