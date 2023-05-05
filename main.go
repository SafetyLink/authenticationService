package main

import (
	"context"
	"fmt"
	"github.com/SafetyLink/authenticationService/internal"
	"github.com/SafetyLink/authenticationService/internal/controller"
	"github.com/SafetyLink/authenticationService/internal/infra/adapter/sql"
	"github.com/SafetyLink/authenticationService/internal/infra/userRepository"
	"github.com/SafetyLink/commons/config"
	"github.com/SafetyLink/commons/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
)

func main() {
	fx.New(
		fx.Provide(logger.InitLogger),
		fx.Provide(config.ReadConfig[internal.Config]),

		fx.Provide(sql.NewPostgresProvider),

		fx.Provide(userRepository.NewUserRepository),

		fx.Provide(SetupAuthenticationServiceGrpcServer),

		fx.Provide(controller.NewAuthenticationServiceGrpcServer),
		fx.Invoke(StartGrpcServer),
	).Run()
}

func StartGrpcServer(lc fx.Lifecycle, grpcServer *controller.AuthenticationServiceGrpcServer, logger *zap.Logger, config *internal.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info(fmt.Sprintf("Starting %s Grpc Server on port %s!", config.Grpc.ServiceName, config.Grpc.Port))
			go func() {
				if err := grpcServer.Gs.Serve(grpcServer.Lis); err != nil {
					logger.Info(fmt.Sprintf("Failed to start %s Grpc Server on port %s!", config.Grpc.ServiceName, config.Grpc.Port))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

func SetupAuthenticationServiceGrpcServer(logger *zap.Logger, config *internal.Config) net.Listener {
	lis, err := net.Listen("tcp", config.Grpc.Port)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to start %s grpc server", config.Grpc.ServiceName), zap.Error(err))
	}

	return lis
}
