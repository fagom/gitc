package hosts

import (
	"github.com/spf13/cobra"
)

// hostsCmd represents the hosts command
var HostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostsCmd.PersistentFlags().String("foo", "", "A help for foo")
	HostsCmd.AddCommand(ListCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
