//nolint:dupl
package firehose

import (
	"fmt"

	"github.com/odpf/salt/printer"
	"github.com/spf13/cobra"

	"github.com/odpf/dex/generated/client/operations"
)

func startCommand(cfgLoader ConfigLoader) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start <project> <firehoseURN>",
		Short: "Start the firehose if it's currently stopped.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error { //nolint:dupl
			spinner := printer.Spin("")
			defer spinner.Stop()

			params := &operations.StartFirehoseParams{
				FirehoseUrn: args[1],
				ProjectSlug: args[0],
				Body:        struct{}{},
			}

			client := initClient(cmd, cfgLoader)
			_, err := client.Operations.StartFirehose(params)
			if err != nil {
				return err
			}

			fmt.Println("Start request accepted. Use view command to check status.")
			return nil
		},
	}
	return cmd
}
