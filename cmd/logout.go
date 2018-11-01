package cmd

import (
	"fmt"
	"agenda-go-cli/service"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "User logout",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Logout called")
		if err := service.UserLogout(); err != true {
			fmt.Println("Some error happened when log out, please read error.log for details")
		} else {
			fmt.Println("Logout Successfully")
		}
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
