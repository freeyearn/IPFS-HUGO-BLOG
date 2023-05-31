package currentUser

import (
	"errors"
	"github.com/gin-gonic/gin"
	"sync"
)

type CurrentUser struct {
	UserID   string `json:"user_id"`
	UserType string `json:"user_type"`
	//OtherInfo map[string]interface{} `json:"other_info"`
}

var userPool = &sync.Pool{
	New: func() any {
		return new(CurrentUser)
	},
}

const (
	UnKnown = "0" // 身份未知
	Admin   = "1" // 超级管理
	Common  = "2" // 普通用户
)

type userTempService interface {
	GetModelMap() (map[string]interface{}, error)
	Get() error
	SetUserID(string)
}

func GetUser(c *gin.Context) (*CurrentUser, error) {
	temp, ok := c.Get("user")
	if !ok {
		return nil, errors.New("无用户")
	}
	user := temp.(*CurrentUser)
	return user, nil
}

func NewUser(userID, userType string) (*CurrentUser, error) {
	var err error
	user := userPool.Get().(*CurrentUser)
	user.UserID = userID
	user.UserType = userType
	//user.OtherInfo, err = rpcReq.GetUserInfo(userID, userType)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Release(user *CurrentUser) {
	user.UserID = ""
	user.UserType = ""
	//user.OtherInfo = map[string]interface{}{}
	userPool.Put(user)
}

// AuthType 验证允许的身份
func (u CurrentUser) AuthType(allowRole ...string) bool {
	if u.IsAdmin() {
		return true
	}
	for _, s := range allowRole {
		if s == u.UserType {
			return true
		}
	}
	return false
}

// AuthSelf 只允许自己
func (u CurrentUser) AuthSelf(UserID string) bool {
	if u.IsAdmin() {
		return true
	}
	if u.UserID == UserID {
		return true
	}
	return false
}

func (u CurrentUser) IsAdmin() bool {
	if u.UserType == Admin {
		return true
	}
	return false
}
