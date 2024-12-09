package config

import "github.com/joho/godotenv"

type Config interface {
	AddresGRPC() string
	LogLevel() string
	DSN() string
}

// Struct of configuration application
type config struct {
	grpchost string
	grpcport string
	loglevel string
	PGdsn    string
}

// Load config file to enviroment variables
func Load(path string) (Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	grpc, err := parseGRPCEnv()
	if err != nil {
		return nil, err
	}

	loglevel, err := parseLogLevelenv()
	if err != nil {
		return nil, err
	}

	pg, err := parsePGEnv()
	if err != nil {
		return nil, err
	}
	return &config{grpchost: grpc.host,
			grpcport: grpc.port,
			loglevel: loglevel.logLevel,
			PGdsn:    pg.dsn,
		},
		nil
}
