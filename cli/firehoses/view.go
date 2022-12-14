package firehoses

import (
	"github.com/odpf/salt/printer"
	"github.com/spf13/cobra"

	"github.com/odpf/dex/cli/cdk"
	"github.com/odpf/dex/generated/client/operations"
)

func viewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view <project> <name>",
		Short: "View a firehose",
		Long:  "Display information about a firehose",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			spinner := printer.Spin("")
			defer spinner.Stop()

			client := initClient(cmd)

			params := operations.GetFirehoseParams{
				ProjectSlug: args[0],
				FirehoseUrn: args[1],
			}

			res, err := client.Operations.GetFirehose(&params)
			if err != nil {
				return err
			}
			firehose := res.GetPayload()

			return cdk.Display(cmd, firehose, cdk.YAMLFormat)
		},
	}

	return cmd
}
