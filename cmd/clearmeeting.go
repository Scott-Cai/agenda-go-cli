package cmd

import (
	"fmt"
	"agenda-go-cli/service"
	"github.com/spf13/cobra"
)
// clearmeetingCmd represents the clearmeeting command
var clearmeetingCmd = &cobra.Command{
	Use:   "clearmeeting",
	Short: "Clear all meeting as a sponsor",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Clear Meeting called")
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("Please Log in firstly!")
		} else {
			cm,flag := service.ClearMeeting(user.Name)
			if flag == true {
				fmt.Println("Successfully clear ", cm," meeting(s)")
			} else {
				fmt.Println("ClearMeeting failed. Check error.log")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(clearmeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
