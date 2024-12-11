package app

import (
	"context"

	"github.com/BelyaevEI/playlist/internal/config"
	"github.com/BelyaevEI/playlist/internal/interceptor"
	"github.com/BelyaevEI/playlist/internal/logger"
	descAuth "github.com/BelyaevEI/playlist/pkg/auth_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// Calls all dependences application
func (a *App) initDependens(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initLogger,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {

	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()), grpc.UnaryInterceptor(interceptor.LogInterceptor))

	reflection.Register(a.grpcServer)

	//start initializating DI - conteiner
	descAuth.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.AuthImpl(ctx))

	return nil
}

// Inititalizating entity service provider
func (a *App) initServiceProvider(_ context.Context) error {

	cfg, err := config.Load("./config.env")
	if err != nil {
		return err
	}
	a.serviceProvider = newServiceProvider(cfg)

	return nil
}

// initLogger initialization entity logger
func (a *App) initLogger(_ context.Context) error {
	logger.Init(logger.GetCore(logger.GetAtomicLevel(a.serviceProvider.config.LogLevel())))

	return nil
}
