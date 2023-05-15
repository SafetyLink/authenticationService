package main

import (
	"context"
	"fmt"
	"github.com/SafetyLink/authenticationService/internal"
	"github.com/SafetyLink/authenticationService/internal/controller"
	"github.com/SafetyLink/authenticationService/internal/domain/authenticationService"
	"github.com/SafetyLink/authenticationService/internal/infra/adapter/sql"
	"github.com/SafetyLink/authenticationService/internal/infra/bcryptRepository"
	"github.com/SafetyLink/authenticationService/internal/infra/userRepository"
	"github.com/SafetyLink/commons/config"
	"github.com/SafetyLink/commons/grpc"
	"github.com/SafetyLink/commons/logger"
	"github.com/SafetyLink/commons/otel"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/zap"
	grpc2 "google.golang.org/grpc"
	"net"
)

func init() {
	//load env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	fx.New(
		fx.Provide(logger.InitLogger),
		fx.Provide(config.ReadConfig[internal.Config]),
		fx.Provide(otel.InitTracer),

		fx.Provide(sql.NewPostgresProvider),

		fx.Provide(userRepository.NewUserRepository),
		fx.Provide(bcryptRepository.NewBcryptRepository),

		fx.Provide(authenticationService.NewAuthenticationService),

		fx.Provide(SetupAuthenticationServiceGrpcServer),

		fx.Invoke(controller.RegisterAuthenticationService),

		fx.Invoke(StartGrpcServer),
	).Run()
}

func StartGrpcServer(lc fx.Lifecycle, s *grpc2.Server, lis net.Listener, logger *zap.Logger, config *internal.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info(fmt.Sprintf("Starting %s Grpc Server on port %s!", config.Grpc.ServiceName, config.Grpc.Port))

			go func() {
				if err := s.Serve(lis); err != nil {
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

func SetupAuthenticationServiceGrpcServer(logger *zap.Logger, config *internal.Config) (net.Listener, *grpc2.Server) {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", config.Grpc.Port)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to start %s grpc server", config.Grpc.ServiceName), zap.Error(err))
	}

	return lis, s
}
