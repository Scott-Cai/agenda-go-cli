package cmd

import (
	"fmt"
	//"strings"
	"agenda-go-cli/service"
	"github.com/spf13/cobra"
)


// removeparticipatorCmd represents the removeparticipator command
var removeparticipatorCmd = &cobra.Command{
	Use:   "removeparticipator",
	Short: "remove participator(s)",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Remove Participator called")
		tmp_t,_ := cmd.Flags().GetString("title")
		tmp_p,_ := cmd.Flags().GetStringSlice("participator")
		if tmp_t == "" || len(tmp_p) == 0 {
			fmt.Println("Please input title and participator(s)(input like \"name1, name2\")")
			return
		}
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("Please login firstly")
		} else {
			// participators := strings.Split(tmp_p, ",")
			flag := service.RemoveMeetingParticipator(user.Name, tmp_t, tmp_p)
			if flag != true {
				fmt.Println("Unexpected error. Check error.log for detail")
			} else {
				fmt.Println("Remove successfully")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(removeparticipatorCmd)

	// Here you will define your flags and configuration settings.
	removeparticipatorCmd.Flags().StringP("title", "t", "", "the title of the meeting")
	removeparticipatorCmd.Flags().StringSliceP("participator", "p", nil, "the participator(s) of the meeting, input like \"name1, name2\"")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeparticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeparticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
