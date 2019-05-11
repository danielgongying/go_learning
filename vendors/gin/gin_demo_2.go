package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)
/*1.获取路径中的参数*/
func routerParams()  {
	router:=gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200,"hello world")
	})
	router.GET("/user/:name", func(c *gin.Context) {
		name:= c.Params.ByName("name")
		c.String(http.StatusOK,"hello "+ name)


	})
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name:=c.Params.ByName("name")
		action:=c.Params.ByName("action")
		message:=name+" is " + action
		c.String(http.StatusOK,message)

	})
	router.Run(":8080")

}
/*2.获取Get参数:http://localhost:8080/
welcome?firstname=Jane&lastname=Doe*/
func getParams()  {
	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		firstname:=c.DefaultQuery("firstname","Guest")
		lastname:=c.Query("lastname")
		c.String(200,"hello "+firstname+" "+lastname)

	})
	router.Run()
}
/*3.获取Post参数
 */
func postParams() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}
/*获取post的参数：http://localhost:8080/
post?id=6&page=1&name=daniel&message=hello*/
func post() {
	router:=gin.Default()
	router.POST("/post", func(c *gin.Context) {
		id:=c.Query("id")
		page:=c.DefaultQuery("page","0")
		name:=c.PostForm("name")
		message:=c.PostForm("message")
		fmt.Println("id:%s,page:%s,name:%s,message:%s",id,page,name,message)

	})
	router.Run()
}
func upload() {
	router := gin.Default()
	router.MaxMultipartMemory = 8<<20
	fmt.Println(router.MaxMultipartMemory)
	router.POST("/upload", func(c *gin.Context) {
		file,_:=c.FormFile("file")
		fmt.Println(file.Filename)
		c.SaveUploadedFile(file,"/users/daniel/desktop")

		c.String(http.StatusOK,file.Filename+" upload success!")
	})
	router.Run()
}
func uploadMore() {
	router := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件到指定的路径
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}
/*路由分组*/
func group() {
	//router := gin.Default()

	// Simple group: v1
	//v1 := router.Group("/v1")
	{
		//v1.POST("/login", loginEndpoint)
		//v1.POST("/submit", submitEndpoint)
		//v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	//v2 := router.Group("/v2")
	{
		//v2.POST("/login", loginEndpoint)
		//v2.POST("/submit", submitEndpoint)
		//v2.POST("/read", readEndpoint)
	}

	//router.Run(":8080")
}
/*日志文件*/
func ginlog()  {
	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run()

}
/*自定义日志格式
 */
func customLog() {
	router := gin.New()

	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
// 绑定为json
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
/*模型绑定和验证*/
func modleBind() {
	router := gin.Default()

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</password>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")


}

func main() {
	//routerParams()
	//getParams()
	//postParams()
	//post()
	//upload()
	//ginlog()
	//customLog()
	modleBind()
}


