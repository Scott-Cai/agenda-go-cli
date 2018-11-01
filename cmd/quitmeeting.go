package cmd

import (
	"fmt"
	"agenda-go-cli/service"
	"github.com/spf13/cobra"
)

// quitmeetingCmd represents the quitmeeting command
var quitmeetingCmd = &cobra.Command{
	Use:   "quitmeeting",
	Short: "quit a meeting",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Quit Meeting called")
		tmp_t, _ := cmd.Flags().GetString("title")
		if tmp_t == "" {
			fmt.Println("Please input meeting title")
			return
		}
		if user,flag := service.GetCurUser(); flag != true {
			fmt.Println("Error: Please login firstly!")
		} else {
			if flag := service.QuitMeeting(user.Name, tmp_t); flag == false {
				fmt.Println("Error: Meeting not exist or you're not a participator of it")
			} else {
				fmt.Println("Quit Successfully")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(quitmeetingCmd)

	// Here you will define your flags and configuration settings.
	quitmeetingCmd.Flags().StringP("title", "t", "", "the title of meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
