package middleware

import (
	"ginserver/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 用户签名，不能泄漏
var Jwtkey = []byte("daiow3wnqdw09qn")

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 生成 JWT Token
func GenJwtToken(username string) (string, int) {
	expiredTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		//Password:       "",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(), // 过期时间
			Issuer:    "ginblog",          // 签名发行者
		},
	}

	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)

	token, err := reqClaims.SignedString(Jwtkey)
	if err != nil {
		return "", utils.ERROR
	}
	return token, utils.SUCCESS
}

// 验证 token
func CheckJwtToken(token string) (*MyClaims, int) {
	settoken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Jwtkey, nil
	})

	if key, ok := settoken.Claims.(*MyClaims); ok && settoken.Valid {
		return key, utils.SUCCESS
	} else {
		return nil, utils.ERROR
	}
}

// 验证 token
func VerifyJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("authorization")
		code := utils.SUCCESS
		// request token is null
		if token == "" {
			code = utils.USER_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": utils.GetErrMsg(code),
				"data":    "",
			})
		}

		key, CheckCode := CheckJwtToken(token)
		// token 验证失败
		if CheckCode == utils.ERROR {
			code = utils.USER_TOKEN_VERIFY_ERROR
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": utils.GetErrMsg(code),
				"data":    "",
			})
			c.Abort()
		}

		// token 过期
		if time.Now().Unix() > key.ExpiresAt {
			code = utils.USER_TOKEN_EXPIRED
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": utils.GetErrMsg(code),
				"data":    "",
			})
			c.Abort()
		}

		// token 验证通过，继续后面的中间件处理
		c.Set("username", key.Username)
		c.Next()
	}
}
