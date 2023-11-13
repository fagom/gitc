package hosts

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hosts = [2]string{"github", "gitlab"}

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List supported hosts",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Supported hosts:")
		for i := 0; i < len(hosts); i++ {
			fmt.Printf("%d. %s\n", i+1, hosts[i])
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
