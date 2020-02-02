package config
import (
	"github.com/spf13/viper"
	"fmt"
)
func Start() {
	viper.SetConfigName("hps")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../../../../conf")
	err :=viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
func Get(key string) string {
	return viper.GetString(key)
}
func GetBool(key string) bool {
	return viper.GetBool(key)
}