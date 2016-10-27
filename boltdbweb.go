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
	dbName     string
	port       string 
	staticPath string
)

func init() {
	// Read the static path from the environment if set.
	staticPath = os.Getenv("BOLTDBWEB_STATIC_PATH")
	if staticPath == "" {
		staticPath = "."
	}
	dbName = os.Getenv("BOLTDBWEB_DB_NAME")
	port = os.Getenv("BOLTDBWEB_PORT")
	if port == "" {
		port = "8080"
	}
	flag.StringVar(&dbName, "db-name", dbName, "Name of the database")
	flag.StringVar(&port, "port", port, "Port for the web-ui")
	flag.StringVar(&staticPath, "static-path", staticPath, "Path for the static content")
}

func main() {
	flag.Parse()
	args := flag.Args()

	fmt.Print(" ")
	log.Info("starting boltdb-browser..")
	if dbName == nil && len(args) > 0 {
		dbName = args[0]
	}

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
