package midware
//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//type Api struct {
//
//}
//
//func (a *Api)GetToken(c *gin.Context)  {
//	username:="cb"
//	password:="123"
//	token, err := GenToken(username, password)
//	if err!= nil{
//		c.JSON(http.StatusOK,gin.H{"code":0,"msg":err})
//		return
//	}
//	c.JSON(http.StatusOK,gin.H{"code":1,"token":token})
//}
