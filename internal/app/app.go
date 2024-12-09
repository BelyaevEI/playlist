package app

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/BelyaevEI/platform_common/pkg/closer"
	initutils "github.com/BelyaevEI/test-assignment/internal/init_utils"
	"github.com/BelyaevEI/test-assignment/internal/logger"

	"google.golang.org/grpc"
)

// App represents the app
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// NewApp creates and initializate a new app.
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDependens(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run application
func (a *App) Run(ctx context.Context) error {

	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := a.runGRPCServer()
		if err != nil {
			logger.Error(fmt.Sprintf("failed to run grpc server: %v", err))
		}

	}()

	initutils.GracefulShutdown(ctx, cancel, wg)
	return nil
}

// Start listen grpc server
func (a *App) runGRPCServer() error {
	logger.Info(fmt.Sprintf("GRPC server is running on %s", a.serviceProvider.config.AddresGRPC()))
	list, err := net.Listen("tcp", a.serviceProvider.config.AddresGRPC())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
