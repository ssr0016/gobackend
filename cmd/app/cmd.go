package app

import (
	"context"

	"backend/pkg/infra/log"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	serviceName = "backend"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: serviceName,
		Run: func(cmd *cobra.Command, args []string) {
			runServer(false)
		},
	}

	return cmd
}

func runServer(isApiServer bool) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log, err := log.New(serviceName)
	if err != nil {
		return err
	}

	undo := zap.ReplaceGlobals(log)
	defer func() {
		if err != nil {
			log.Error(err.Error())
		}

		undo()
		log.Sync()
	}()

	s, err := NewServer(isApiServer)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	s.Run(ctx)

	return nil
}
