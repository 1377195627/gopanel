package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"gitlab.com/xiayesuifeng/gopanel/core"
	"log"
	"os"
)

var (
	port = flag.Int("p", 8080, "port")
	help = flag.Bool("h", false, "help")
)

func main() {
	router := gin.Default()

	if err := router.Run(); err != nil {
		log.Panicln(err)
	}
}

func init() {
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	err := core.ParseConf("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("please config config.json")
			os.Exit(0)
		}
		log.Panicln(err)
	}
}
