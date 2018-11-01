package cmd

import (
	"fmt"
	"agenda-go-cli/service"
	"github.com/spf13/cobra"
)

// deleteaccountCmd represents the deleteaccount command
var deleteaccountCmd = &cobra.Command{
	Use:   "deleteaccount",
	Short: "to delete current user",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Delete account called")
		if user,flag := service.GetCurUser(); flag != true {
			fmt.Println("You need login firstly!")
		} else {
			if dflag := service.DeleteUser(user.Name); dflag != true {
				fmt.Println("Error occurred when delete account.Please check error.log")
			} else {
				fmt.Println("Successfully Delete")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteaccountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteaccountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteaccountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
