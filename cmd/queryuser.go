package cmd

import (
	"fmt"
	"agenda-go-cli/service"
	"github.com/spf13/cobra"
)

// queryuserCmd represents the queryuser command
var queryuserCmd = &cobra.Command{
	Use:   "queryuser",
	Short: "To query user",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Query User called")
		if _, flag := service.GetCurUser(); flag != true {
			fmt.Println("Please Log in firstly")
		}
		ru := service.ListAllUser()
		for _, u := range ru {
			fmt.Println("----------------")
			fmt.Println("Username: ", u.Name)
			fmt.Println("Phone: ",u.Phone)
			fmt.Println("Email: ",u.Email)
			fmt.Println("----------------")
		}
	},
}

func init() {
	RootCmd.AddCommand(queryuserCmd)

	// Here you will define your flags and configuration settings.
	// queryuserCmd.Flags().StringP("username", "u", "", "specify user")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
