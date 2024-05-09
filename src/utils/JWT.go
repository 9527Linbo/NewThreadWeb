package utils

import (
	"NewThread/src/pojo"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type MyClaims struct {
	Account            string //信息项
	jwt.StandardClaims        // claims设置项
}

var TokenExpireDuration int //JWT有效时间
var MySecret []byte         // 签名的密钥

// 生成JWT令牌
func GenerateToken(usermsg pojo.RecvUserMsg) (string, error) {
	TokenExpireDuration = viper.GetInt("JWT.TTL")                                 //默认6小时
	MySecret = []byte(viper.GetString("JWT.Secret"))                              //密钥
	expiresat := time.Now().Add(time.Duration(TokenExpireDuration) * time.Second) // 六小时后
	claims := &MyClaims{
		Account: usermsg.Account, //用户账号信息（不建议展示密码信息）
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresat.Unix(), //设置到期时间
		},
	}

	// 生成Token，指定签名算法和claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名
	tokenString, err := token.SignedString(MySecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenstring string) (*MyClaims, error) {
	claims := &MyClaims{}
	_, err := jwt.ParseWithClaims(tokenstring, claims, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	// 若token只是过期claims是有数据的，若token无法解析claims无数据
	return claims, err
}
