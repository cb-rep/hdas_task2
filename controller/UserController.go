package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hdas_task/dao"
	"hdas_task/midware"
	"hdas_task/model"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	u := dao.Mgr.Login(username)

	user := model.User{
		Username: username,
		Password: password,
	}

	verifyPassword := midware.VerifyPassword(8, 16, password)

	if u.Username == "" {
		if verifyPassword == true {
			dao.Mgr.Register(&user)
			c.Redirect(301, "/")
		} else {
			c.HTML(200, "register.html", "密码必须包含数字、大写字母、小写字母、特殊字符(如.@$!%*#_~?&^)至少3种的组合且长度在8-16之间")
		}
	} else {
		c.HTML(200, "register.html", "用户名已存在")
	}
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	token, _ := midware.GenToken(username, password)
	//c.SetCookie("token",token,60*60*24*7, "/books","",true,false)
	//c.Set("token",token)

	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在！")
		fmt.Println("用户名不存在！")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Println("登陆成功")
			c.SetCookie("token", token, 60*60*24*7, "/user", "", true, false)
			//c.Set("token",token)
			c.Redirect(301, "/books")
		}
	}
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", 60*60*24*7, "/user", "", true, false)
	c.HTML(200, "index.html", nil)
}

func GetUser(c *gin.Context) {
	//token := c.GetHeader("token")
	token, _ := c.Cookie("token")
	mc, _ := midware.ParseToken(token)
	c.HTML(200, "user.html", mc.Username)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

//func User(c *gin.Context) {
//	c.HTML(200,"books.html",nil)
//}
