package config

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	Port       string `mapstructure:"PORT"`
	
	DOCTOR_SVC string `mapstructure:"DOCTOR_SVC"`
	KEY_ID_fOR_PAY     string `mapstructure:"KEY_ID_fOR_PAY"`
	SECRET_KEY_FOR_PAY string `mapstructure:"SECRET_KEY_FOR_PAY"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "PORT","DOCTOR-SVC","KEY_ID_fOR_PAY","SECRET_KEY_FOR_PAY",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil

}
