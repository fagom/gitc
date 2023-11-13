package cmd

import (
	"fmt"
	"os"

	"github.com/fagom/gitc/cmd/hosts"
	"github.com/fagom/gitc/cmd/profile"
	. "github.com/fagom/gitc/internal"
	"github.com/spf13/cobra"
)

var logger *os.File

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitc",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer logger.Close()
}

func addSubCommands() {
	rootCmd.AddCommand(profile.ProfileCmd)
	rootCmd.AddCommand(hosts.HostsCmd)
}

func init() {
	BuildConfiguration()
	logger = OpenLogger()
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommands()

}
