package cli

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/odpf/salt/cmdx"
	"github.com/spf13/cobra"
)

type CDK struct {
	Config *cmdx.Config
}

// New root command.
func New() *cobra.Command {
	var cmd = &cobra.Command{
		Use:           "dex <command> <subcommand> [flags]",
		Short:         "Data experience console",
		Long:          "Data experience console.",
		SilenceUsage:  true,
		SilenceErrors: true,
		Annotations: map[string]string{
			"group": "core",
			"help:learn": heredoc.Doc(`
				Use 'dex <command> --help' for info about a command.
				Read the manual at https://odpf.github.io/dex/
			`),
			"help:feedback": heredoc.Doc(`
				Open an issue here https://github.com/odpf/dex/issues
			`),
		},
	}

	cdk := &CDK{Config: cmdx.SetConfig("dex")}

	cmd.AddCommand(
		serverCommand(),
		configCmd(cdk),
		versionCmd(),
	)

	// Help topics.
	cmdx.SetHelp(cmd)
	cmd.AddCommand(
		cmdx.SetCompletionCmd("dex"),
		cmdx.SetHelpTopicCmd("environment", envHelp),
		cmdx.SetRefCmd(cmd),
	)

	cmdx.SetClientHook(cmd, func(cmd *cobra.Command) {
		// client config.
		cmd.PersistentFlags().StringP("host", "h", "", "Server host address")
	})

	return cmd
}
