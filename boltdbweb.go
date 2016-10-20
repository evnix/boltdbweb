package main

import (
	"github.com/evnix/boltdbweb/web"
	"github.com/gin-gonic/gin"
)

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
)

var (
	db         *bolt.DB
	dbName     = flag.String("db-name", "", "Name of the database")
	port       = flag.String("port", "8080", "Port for the web-ui")
	staticPath = flag.String("static-path", ".", "Path for the static content")
)

func main() {
	flag.Parse()

	fmt.Print(" ")
	log.Info("starting boltdb-browser..")

	if dbName == nil {

		fmt.Println("Usage: " + os.Args[0] + " --db-name=<DBfilename>[required] --port=<port>[optional] --static-path=<static-path>[optional]")
		os.Exit(0)
	}

	var err error
	db, err = bolt.Open(*dbName, 0600, nil)
	boltbrowserweb.Db = db

	if err != nil {

		fmt.Println(err)
		os.Exit(0)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", boltbrowserweb.Index)

	r.GET("/buckets", boltbrowserweb.Buckets)
	r.POST("/createBucket", boltbrowserweb.CreateBucket)
	r.POST("/put", boltbrowserweb.Put)
	r.POST("/get", boltbrowserweb.Get)
	r.POST("/deleteKey", boltbrowserweb.DeleteKey)
	r.POST("/deleteBucket", boltbrowserweb.DeleteBucket)
	r.POST("/prefixScan", boltbrowserweb.PrefixScan)

	r.Static("/web", *staticPath+"/web")

	r.Run(":" + *port)

}
