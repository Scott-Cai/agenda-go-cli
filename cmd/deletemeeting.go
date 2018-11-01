package cmd

import (
	"fmt"
	"agenda-go-cli/service"
	"github.com/spf13/cobra"
)

// deletemeetingCmd represents the deletemeeting command
var deletemeetingCmd = &cobra.Command{
	Use:   "deletemeeting",
	Short: "delete meeting",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Delete Meeting called")
		tmp_t, _ := cmd.Flags().GetString("title")
		if tmp_t == "" {
			fmt.Println("Error: Please input meeting title")
			return
		}
		if user,flag := service.GetCurUser(); flag != true {
			fmt.Println("Error: Please login firstly!")
		} else {
			if c := service.DeleteMeeting(user.Name, tmp_t); c == 0 {
				fmt.Println("Error: Meeting not exist or you're not a Sponsor of it")
			} else {
				fmt.Println("Delete Successfully")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(deletemeetingCmd)

	// Here you will define your flags and configuration settings.
	deletemeetingCmd.Flags().StringP("title", "t","", "the title of meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deletemeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deletemeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
