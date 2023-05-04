package internal

type Config struct {
	Grpc     grpc     `yaml:"grpc"`
	Postgres postgres `yaml:"postgres"`
}
type grpc struct {
	ServiceName string `yaml:"serviceName"`
	Port        string `yaml:"port"`
}
type postgres struct {
	ConnectionURL string `yaml:"connectionURL"`
}
