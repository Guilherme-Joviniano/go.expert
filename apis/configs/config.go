package configs

import (
	"fmt"
	"log"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDRiver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	configurator := viper.New()

	configurator.SetConfigName("app_config")
	configurator.AddConfigPath(path)
	configurator.SetConfigType("env")
	configurator.SetConfigFile(".env")
	configurator.AutomaticEnv()

	if err := configurator.ReadInConfig(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	config := conf{
		DBDRiver:      "",
		DBHost:        "",
		DBPort:        "",
		DBUser:        "",
		DBPassword:    "",
		DBName:        "",
		WebServerPort: "",
		JWTSecret:     "",
		JWTExpiresIn:  0,
		TokenAuth:     &jwtauth.JWTAuth{},
	}

	if err := configurator.Unmarshal(&config); err != nil {
		err = fmt.Errorf("error loading viper values: %w", err)

		panic(err)
	}

	config.TokenAuth = jwtauth.New("H256", []byte(config.JWTSecret), nil)

	return &config, nil
}
