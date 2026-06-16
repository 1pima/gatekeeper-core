package setup

import (
	"github.com/spf13/viper"
)

// Configure Для сериализации .env конфигурации приложения в необходимую структуру
func Configure(cfg any) error {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	return nil
}
