package main

import "github.com/gin-gonic/gin"
import "github.com/evnix/boltdbweb/web"

import ("fmt"
		log "github.com/Sirupsen/logrus"
		 "github.com/boltdb/bolt"
		 "os")		

var db *bolt.DB

func main() {

	fmt.Print(" ")
	log.Info("starting boltdb-browser..")

	if(len(os.Args)<2){

		fmt.Println("Usage: "+os.Args[0]+" <DBfilename>[required] <port>[optional]")
		os.Exit(0)
	}

	var err error
	db,err=bolt.Open(os.Args[1], 0600, nil)
	boltbrowserweb.Db=db
	
	if(err!=nil){

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


    r.Static("/web", "./web")

    port:="8080";

    if(len(os.Args)>2){

    	port=os.Args[2]
    }

    r.Run(":"+port)


}
