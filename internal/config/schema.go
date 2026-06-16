package config

var Settings SettingsSchema

type SettingsSchema struct {
	// gateway
	GRPCAddr string `mapstructure:"GRPC_ADDR"`
	BankAddr string `mapstructure:"BANK_ADDR"`

	// internal TCP
	ServerAddr   string `mapstructure:"SERVER_ADDR"`
	PingInterval int64  `mapstructure:"PING_INTERVAL"`

	// logging
	SYSLogAddr string `mapstructure:"SYSLOG_TARGET"`
	HMACKey    string `mapstructure:"HMAC_KEY"`
}
