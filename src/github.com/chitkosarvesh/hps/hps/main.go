package main
import (
	"fmt"
	"github.com/chitkosarvesh/hps/hps/modules/server"
	"github.com/chitkosarvesh/hps/hps/modules/config"
)

func main(){
	config.Start()
	fmt.Println(config.Get("name"))
	server.Start()
}