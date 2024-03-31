package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")

	if userType != role {
		err = fmt.Errorf("UnAuthenticated to access this resource")
		return err
	}
	return nil
}

// user can access this resource only via his token or he is an ADMIN...
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	uid := c.GetString("uid")
	userType := c.GetString("user_type")

	if err := CheckUserType(c, "ADMIN"); err == nil {
		return nil
	}

	if uid == userId && userType == "USER" {
		return nil
	}

	err = fmt.Errorf("UnAuthenticated to access this resource")
	return err
}
