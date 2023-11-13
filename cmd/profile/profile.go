package profile

import (
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var ProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Profile command",
	Long:  `Profile long command`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubCommands() {
	ProfileCmd.AddCommand(ProfileListCmd)
	ProfileCmd.AddCommand(ProfileAddCmd)
}

func init() {

	// Here you will define your flags and configuration settings.
	// ProfileCmd.Flags().StringVar()

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommands()
}
