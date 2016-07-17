package main

import "github.com/gin-gonic/gin"
import "github.com/evnix/boltdb-browser/web"

import ("fmt"
		log "github.com/Sirupsen/logrus"
		 "github.com/boltdb/bolt"
		)

var db *bolt.DB

func main() {

	fmt.Print(" ")
	log.Info("starting Ashtra...")
	var err error
	db,err=bolt.Open("ashtra.db", 0600, nil)
	boltbrowserweb.Db=db
	err=err


	r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.GET("/", boltbrowserweb.Index)

    r.GET("/buckets", boltbrowserweb.Buckets)
    r.POST("/createBucket", boltbrowserweb.CreateBucket)


    r.Static("/web", "./web")


    r.Run()


}