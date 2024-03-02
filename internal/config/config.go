package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func setup() {
	viper.SetConfigName("app.conf")
	viper.AddConfigPath("../../configs/app.conf")
}

func DB_URI() string {
	setup()
	fmt.Println(viper.GetString("POSTGRES_DB_URI"))
	return viper.GetString("POSTGRES_DB_URI=")
}
