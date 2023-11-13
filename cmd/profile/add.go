package profile

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/fagom/gitc/internal"
	. "github.com/fagom/gitc/models"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var ProfileAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add git profile",
	Long:  `The command lets you add a git profile reference in your host machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Add logic here
		data := []User{
			User{
				PassKey: password,
				Host:    host,
			},
		}
		Logger.Println("testing")

		// CheckUserAuthKey(data)

		file, _ := json.MarshalIndent(data, "", "\t")
		_ = os.WriteFile("test.json", file, 0644)
	},
}

var (
	userName string
	password string
	host     string
)

func init() {
	ProfileAddCmd.Flags().StringVarP(&password, "passkey", "p", "", "Password/Auth key of git source")
	ProfileAddCmd.Flags().StringVarP(&host, "host", "a", "", `Hostname of git source. Run "gitc hosts ls" to get a list of supported hosts`)

	if err := ProfileAddCmd.MarkFlagRequired("passkey"); err != nil {
		fmt.Println(err)
	}

	if err := ProfileAddCmd.MarkFlagRequired("host"); err != nil {
		fmt.Println(err)
	}
}
