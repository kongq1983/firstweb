package myweb

//import "github.com/gin-gonic/gin"
import (
	"com.kq/dao"
	"com.kq/dao/employee"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http" // 1.18 这个会有问题  time
	"strconv"
	//"dto/Employee"
	"com.kq/dto"
)

// https://github.com/gin-gonic/gin/
func main111() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func Init() {

	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	// add
	r.POST("/employee", func(c *gin.Context) {
		//var username = c.PostForm("username")
		//var name = c.PostForm("name")

		var username = c.Query("username")
		var name = c.Query("name")
		var age, err = strconv.Atoi(c.Query("age"))
		if err != nil {
			age = 0
		}

		var emp = dto.Employee{Username: username, Name: name, Age: age}

		//fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		fmt.Printf("username: %s; age: %d; name: %s; \n", emp.Username, emp.Age, emp.Name)

		employee.Add(emp)

		//c.JSON(http.StatusOK, emp)
		c.JSON(http.StatusOK, gin.H{"username": username, "name": name, "age": age, "method": "post"})
	})

	// modify
	r.PUT("/employee", func(c *gin.Context) {
		//var username = c.PostForm("username")
		//var name = c.PostForm("name")

		var id, err0 = strconv.Atoi(c.Query("id"))
		if err0 != nil {
			id = -1
		}
		var username = c.Query("username")
		var name = c.Query("name")
		var age, err = strconv.Atoi(c.Query("age"))
		if err != nil {
			age = 0
		}

		var emp = dto.Employee{Id: id, Username: username, Name: name, Age: age}
		employee.Update(emp)

		//fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		fmt.Printf("username: %s; age: %d; name: %s; \n", emp.Username, emp.Age, emp.Name)

		//c.JSON(http.StatusOK, emp)
		c.JSON(http.StatusOK, gin.H{"username": username, "name": name, "age": age, "method": "put"})
	})

	r.GET("/employee/:id", func(c *gin.Context) {

		var id, err0 = strconv.Atoi(c.Param("id"))
		if err0 != nil {
			id = -1
		}

		log.Printf("GET id = %s \n", id)

		var emp = employee.Get(id)

		c.JSON(http.StatusOK, gin.H{"employee": emp, "method": "get"})
	})

	r.DELETE("/employee/:id", func(c *gin.Context) {
		var id, _ = strconv.Atoi(c.Param("id"))
		log.Printf("DELETE id = %s \n", id)

		employee.Delete(id)

		c.JSON(http.StatusOK, gin.H{"id": id, "username": "guest", "name": "匿名用户", "age": 18, "method": "delete"})
	})

	// modify
	r.PUT("/user", func(c *gin.Context) {
		var emp dto.Employee

		if err := c.ShouldBindJSON(&emp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//var emp = Employee{id:id,username:username,name:name,age:age}

		//fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		//fmt.Printf("id:%d ; username: %s; age: %d; name: %s; \n", emp.id,emp.username, emp.age, emp.name)

		//c.JSON(http.StatusOK, emp)
		c.JSON(http.StatusOK, gin.H{"user": emp})
	})

	r.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/json1", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	if err := dao.InitDb(); err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	fmt.Println("init DB success")

	// 监听并在 0.0.0.0:8080 上启动服务
	//r.Run()
	Start(r)

}

func Start(r *gin.Engine) {
	log.Printf("the server is started ! bind port:8080 ! \n")

	// // 监听并在 0.0.0.0:8080 上启动服务
	r.Run()
}
