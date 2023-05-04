package main

import (
	"context"
	"fmt"
	"github.com/SafetyLink/authenticationService/internal"
	"github.com/SafetyLink/authenticationService/internal/controller"
	"github.com/SafetyLink/commons/config"
	"github.com/SafetyLink/commons/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(logger.InitLogger),
		fx.Provide(config.ReadConfig[internal.Config]),

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
