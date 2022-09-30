package cli

import (
	"context"

	"github.com/MakeNowJust/heredoc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/spf13/cobra"
	shieldv1beta1 "go.buf.build/odpf/gwv/odpf/proton/odpf/shield/v1beta1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/odpf/dex/config"
	"github.com/odpf/dex/internal/server"
	"github.com/odpf/dex/pkg/logger"
	"github.com/odpf/dex/pkg/telemetry"
)

func serverCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server <command>",
		Aliases: []string{"s"},
		Short:   "Server management",
		Long:    "Server management commands.",
		Example: heredoc.Doc(`
			$ dex server start
			$ dex server start -c ./config.yaml
		`),
		Annotations: map[string]string{
			"group": "core",
		},
	}

	cmd.AddCommand(startCommand())
	return cmd
}

func startCommand() *cobra.Command {
	var configFile string

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the server",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(configFile)
		if err != nil {
			return err
		}

		zapLog, err := logger.New(&cfg.Log)
		if err != nil {
			return err
		}

		telemetry.Init(cmd.Context(), cfg.Telemetry, zapLog)
		nrApp, err := newrelic.NewApplication(
			newrelic.ConfigAppName(cfg.Telemetry.ServiceName),
			newrelic.ConfigLicense(cfg.Telemetry.NewRelicAPIKey),
		)
		if err != nil {
			return err
		}
		return runServer(cmd.Context(), nrApp, zapLog, cfg)
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "./config.yaml", "Config file path")

	return cmd
}

func runServer(baseCtx context.Context, nrApp *newrelic.Application, zapLog *zap.Logger, cfg config.Config) error {
	ctx, cancel := context.WithCancel(baseCtx)
	defer cancel()

	shieldConn, err := grpc.Dial(cfg.Shield.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	return server.Serve(ctx, cfg.Service.Addr(), nrApp, zapLog,
		shieldv1beta1.NewShieldServiceClient(shieldConn),
	)
}