# Web Push in go
## Web Push using Go and Gin-Gonic framework with Moengage webpush API

go get -u github.com/rakeshkumargupt/webpush-moengage

# Example
package main

import (

	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/rakeshkumargupt/router"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
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

func startServer() {

	ginRouter = router.SetRoutes()

}

# router.go

func SetRoutes() *gin.Engine{

	router := gin.Default()
	// Endpoint call url : localhost:8080/sendwebpush
	router.POST("/sendwebpush", utils.Validator(models.UiWebpushModel{}), handler.SendWebpush)
	return router
}
# Dependencies
  1. Install Go 
  2. Get Moengage Subscription Id
  
  
# Similar Projects / Inspired By

   [Push Encryption (Go)](https://github.com/GoogleChrome/push-encryption-go)
   
   [go-ecdh](https://github.com/wsddn/go-ecdh)
   

# References

For more information visit these Google Developers links:

https://developers.google.com/web/updates/2016/03/web-push-encryption 

https://developers.google.com/web/updates/2016/07/web-push-interop-wins
  

