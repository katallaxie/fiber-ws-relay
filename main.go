package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/katallaxie/template-go/internal/adapters"
	"github.com/katallaxie/template-go/internal/controllers"
	"github.com/katallaxie/template-go/internal/services/noop"

	"github.com/katallaxie/pkg/logger"
	"github.com/katallaxie/pkg/server"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {
	rootCmd.SilenceUsage = true
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	zap.RedirectStdLog(logger.LogSink)

	srv, _ := server.WithContext(ctx)

	ctrl := controllers.New(adapters.NewNoop())

	service := noop.New(ctrl)
	srv.Listen(service, true)

	if err := srv.Wait(); errors.Is(err, &server.Error{}) {
		return err
	}

	return nil
}
