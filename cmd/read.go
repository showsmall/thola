package cmd

import (
	"fmt"
	"github.com/inexio/thola/core/request"
	"github.com/spf13/cobra"
)

func init() {
	addDeviceFlags(readCMD)
	rootCMD.AddCommand(readCMD)
}

var readCMD = &cobra.Command{
	Use:   "read",
	Short: "Read out information of a device",
	Long: "Read out information of a device.\n\n" +
		"You need to specify the information which you want to read out with a subcommand.",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.UsageString())
	},
}

func getReadRequest() request.ReadRequest {
	return request.ReadRequest{
		BaseRequest: getBaseRequest(),
	}
}
