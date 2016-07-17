package boltbrowserweb


import ("github.com/gin-gonic/gin"
"github.com/boltdb/bolt"
"fmt")



var Db *bolt.DB

func Index(c *gin.Context) {

		c.String(200, "Hello ");
	
}


func CreateBucket(c *gin.Context) {



	Db.Update(func(tx *bolt.Tx) error {
    	b, err := tx.CreateBucket([]byte(c.PostForm("bucket")))
    	b=b
    	if err != nil {
        	return fmt.Errorf("create bucket: %s", err)
   		 }
    return nil
	})
		c.String(200, "ok");
	
}

func Buckets(c *gin.Context) {

	

	res:= []string{}

		Db.View(func(tx *bolt.Tx) error {
		
			
        return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
            
            b:=[]string{string(name)}
            res=append(res, b...) 
            return nil
       	 })



        })

		c.JSON(200, res)
	
}