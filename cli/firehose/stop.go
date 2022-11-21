package firehose

import (
	"fmt"

	"github.com/odpf/salt/printer"
	"github.com/spf13/cobra"

	"github.com/odpf/dex/generated/client/operations"
)

func stopCommand(cfgLoader ConfigLoader) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Stop <project> <firehoseURN>",
		Short: "Stop the firehose if it's currently running.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			spinner := printer.Spin("")
			defer spinner.Stop()

			client := initClient(cfgLoader)
			params := &operations.StopFirehoseParams{
				FirehoseUrn: args[1],
				ProjectSlug: args[0],
				Body:        struct{}{},
			}

			_, err := client.Operations.StopFirehose(params)
			if err != nil {
				return err
			}

			fmt.Println("Stop request accepted. Use view command to check status.")
			return nil
		},
	}
	return cmd
}
