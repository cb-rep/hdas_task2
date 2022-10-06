package midware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var Key =[]byte("secret")

func GenToken(username string, password string) (string, error) {
	c:=MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour *2).Unix(),
			Issuer: "chenbo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Key)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token,err:=jwt.ParseWithClaims(tokenString,&MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return Key,nil
		})
	if err!= nil {
        return nil,err
    }
    if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
    	return claims,nil
	}
	return nil,errors.New("Invalid token")
}


//func SetToken() {
//
//}

//type MyClaims struct {
//	Username string `json:"username"`
//	jwt.StandardClaims
//}
//
//
//func Verify() {
//
//	mySigningKey:=[]byte("woshiqimiaofangguowoba")
//
//	c:=MyClaims{
//		Username: "qimiao",
//		StandardClaims: jwt.StandardClaims{
//			NotBefore: time.Now().Unix() - 60,
//			ExpiresAt: time.Now().Unix() + 60*60*2,
//			Issuer: "qimiao",
//		},
//	}
//	t:=jwt.NewWithClaims(jwt.SigningMethodHS256,c)
//	s, e := t.SignedString(mySigningKey)
//	if e!= nil {
//		fmt.Printf("%s",e)
//  }
//  fmt.Println(s)
//
//	token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return mySigningKey, nil
//	})
//
//	if err != nil {
//		fmt.Printf("%s",err)
//		return
//	}
//	fmt.Println(token.Claims.(*MyClaims))
//	fmt.Println(token.Claims.(*MyClaims).Username)
//}

