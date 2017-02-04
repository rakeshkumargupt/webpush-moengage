package main

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"github.com/rakeshkumargupt/router"
)

//Initialize gin Engine
var ginRouter *gin.Engine

func init() {

	// config file name
	viper.SetConfigName("res")

	//full path of config file from $GOPATH
	viper.AddConfigPath("$GOPATH/src/github.com/rakeshkumargupt/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to Read file")
		return
	}
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		viper.WatchConfig()
	})
	startServer()
}



func main() {

	ginRouter.Run(viper.GetString("server") + ":" + viper.GetString("port"))

}

func startServer(){

	ginRouter = router.SetRoutes()

}

