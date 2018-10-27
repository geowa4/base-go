package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//ReadConfig reads the configuration from the environment.
func ReadConfig(serviceName string) {
	viper.SetDefault("log_level", "info")
	viper.SetDefault("app_port", "8000")
	viper.SetDefault("metrics_port", "8001")
	viper.SetDefault("db_host", "127.0.0.1")
	viper.SetDefault("db_port", "5432")
	viper.SetDefault("db_user", "postgres")
	viper.SetDefault("db_database", "postgres")
	viper.SetDefault("db_sslmode", "disable")
	viper.SetEnvPrefix(serviceName)
	viper.AutomaticEnv()
}

//DatabaseConnectionString reads the database config variables
//to provide a connection string.
func DatabaseConnectionString() string {
	host := viper.GetString("db_host")
	port := viper.GetInt32("db_port")
	user := viper.GetString("db_user")
	database := viper.GetString("db_database")
	sslmode := viper.GetString("db_sslmode")
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s", host, port, user, database, sslmode)
}
