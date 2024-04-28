package totp

import (
	"himitsu/utils"
	"strings"
	"time"

	totpLib "github.com/pquerna/otp/totp"
)

type TOTP struct {
	Label   string
	Secret  string
	Account string
}

func GetCode(secret string) string {
	code, err := totpLib.GenerateCode(strings.ToUpper(strings.ReplaceAll(secret, " ", "")), time.Now())
	utils.CheckError(err)
	return code

}
