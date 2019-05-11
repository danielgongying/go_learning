package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
var db=make(map[string]string)

func hello()  {
	r:=gin.New()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello gin")

	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK,"pong000")

	})
	//http.Handle("/",r)
	r.Run()


}
func basic() {
	r:=gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK,"basic pong")

	})
	r.GET("/user/:name", func(c *gin.Context) {
		user:=c.Params.ByName("name")
		fmt.Println(user)
		db["daniel"] = "owner"
		db [ "France" ] = "Paris"
		db [ "Italy" ] = "罗马"
		db [ "Japan" ] = "东京"
		db [ "India " ] = "新德里"
		fmt.Println(db )
		value,ok:=db[user]
		if ok {
			c.JSON(http.StatusOK,gin.H{
				"user":user,
				"value":value,
			})

		}else {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
				"status": "no value",
			})
		}


	})
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"daniel": "123456",
	}))
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey)
		fmt.Println(user)


	})

	r.Run()

}
func main() {
	//hello()
	basic()
}

