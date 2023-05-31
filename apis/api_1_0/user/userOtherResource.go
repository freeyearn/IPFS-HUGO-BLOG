package user

import (
	"IPFS-Blog-Hugo/internal/services"
	"IPFS-Blog-Hugo/utils"
	"IPFS-Blog-Hugo/utils/jwt"
	"IPFS-Blog-Hugo/utils/parser"
	"IPFS-Blog-Hugo/utils/security"
	"IPFS-Blog-Hugo/utils/structs"
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Register helps user register an account. Param account and password is needed while wallet isn't needed.
func Register(c *gin.Context) {
	var err error

	var Parser struct {
		Wallet   string `json:"wallet" form:"wallet"`
		Account  string `json:"account" form:"account" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var UserService services.UserService
	UserService.Assign(Parser)

	userId := utils.CreateRandomId(viper.GetInt("security.IDSuffixNum"))
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	UserService.SetUserId(userId)

	if UserService.Wallet != "" {
		wallet, err := security.RSAHelper.Encrypt([]byte(UserService.Wallet))
		UserService.SetWallet(string(wallet))
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
	}

	password, _ := security.RSAHelper.Encrypt([]byte(Parser.Password))
	// 必须编码为base64，解决乱码问题
	UserService.SetPassword(base64.StdEncoding.EncodeToString(password))

	delete(args, "password")
	if args["wallet"] == "" {
		delete(args, "wallet")
	}

	err = UserService.Add(args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", UserService)
}

// Login lets user login into the system. Notice that the password containing a handled timestamp must be signed with RSA public key.
func Login(c *gin.Context) {
	var err error

	var Parser struct {
		Account  string `json:"account" form:"account" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var UserService services.UserService
	UserService.Assign(Parser)

	delete(args, "password")
	args["status"] = 0
	err = UserService.Get(args)
	if err != nil {
		parser.JsonAccessDenied(c, "用户名或密码错误")
		return
	}

	// check password
	pwd, err := base64.StdEncoding.DecodeString(UserService.Password)
	if err != nil {
		parser.JsonAccessDenied(c, "用户名或密码错误")
		return
	}
	password, err := security.RSAHelper.Decrypt(pwd)

	if err != nil {
		parser.JsonAccessDenied(c, "用户名或密码错误")
		return
	}
	pwdCheck, err := base64.StdEncoding.DecodeString(Parser.Password)
	if err != nil {
		parser.JsonAccessDenied(c, "用户名或密码错误")
		return
	}
	passwordCheck, err := security.RSAHelper.TimeStampDecrypt(pwdCheck, viper.GetInt("system.PasswordExpireTime"))
	if err != nil {
		parser.JsonAccessDenied(c, "用户名或密码错误")
		return
	}

	if !bytes.Equal(password, passwordCheck) {
		err = errors.New("wrong account or password, try again")
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var jwtClaim = jwt.JWTClaims{
		Username: UserService.Name,
		UserId:   UserService.UserId,
	}
	token, err := jwtClaim.MakeToken(viper.GetInt("system.TokenExpireTime"), []byte(viper.GetString("system.Secret")))
	if err != nil {
		return
	}

	type Result struct {
		UserId string
		Name   string
		Token  string
	}
	result := Result{
		UserId: UserService.GetUserId(),
		Name:   UserService.GetName(),
		Token:  token,
	}
	parser.JsonOK(c, "", result)
}

// GetUserInfoByToken gets user info with token.
func GetUserInfoByToken(c *gin.Context) {
	var err error

	token := c.GetHeader("token")

	jwtClaim, err := jwt.VerifyToken(token, []byte(viper.GetString("system.Secret")))
	if err != nil {
		parser.JsonAccessDenied(c, "please login")
		return
	}

	var UserService services.UserService
	err = UserService.Get(map[string]any{
		"user_id": jwtClaim.UserId,
	})
	if err != nil {
		parser.JsonParameterIllegal(c, "user not exist", err)
		return
	}

	parser.JsonOK(c, "", UserService)
}
