package runnable

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/calmonr/cicd/cmd/playground/app"
	"github.com/calmonr/cicd/internal/cli"
	"github.com/calmonr/cicd/internal/command"
	"github.com/calmonr/cicd/internal/logger"
	"github.com/calmonr/cicd/internal/service/greeting"
	greetingv1 "github.com/calmonr/cicd/pkg/proto/greeting/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewRoot() command.Runnable {
	return func(cmd *cobra.Command, args []string) error {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()

		logger, err := logger.New()
		if err != nil {
			return fmt.Errorf("could not create logger: %w", err)
		}

		defer func() {
			// it should be safe to ignore according to https://github.com/uber-go/zap/issues/328
			_ = logger.Sync()
		}()

		// read config from file if the config file is specified
		{
			if file := viper.GetString(cli.ConfigFileFlag); file != "" {
				viper.SetConfigFile(file)

				if err := viper.ReadInConfig(); err != nil {
					log.Fatalf("could not read config file: %v", err)
				}
			}
		}

		cfg := new(app.Config).Fill()

		listener, err := net.Listen(cfg.GRPCServer.Network, cfg.GRPCServer.Address)
		if err != nil {
			return fmt.Errorf("could not listen on %s: %w", cfg.GRPCServer.Address, err)
		}

		server := grpc.NewServer()

		greetingv1.RegisterGreetingServiceServer(server, &greeting.Service{})

		reflection.Register(server)

		g, ctx := errgroup.WithContext(ctx)

		g.Go(func() error {
			logger.Info("serving grpc server", zap.String("address", cfg.GRPCServer.Address))

			if err := server.Serve(listener); err != nil {
				return fmt.Errorf("could not serve grpc: %w", err)
			}

			return nil
		})

		g.Go(func() error {
			<-ctx.Done()

			logger.Info("gracefully stopping the application")

			server.GracefulStop()

			return nil
		})

		if err := g.Wait(); err != nil {
			return fmt.Errorf("could not stop application: %w", err)
		}

		return nil
	}
}
