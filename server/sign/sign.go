package sign

import (
	"crypto/md5"
	"fmt"
)

type SignParam struct {
	Body string
	Key  string
}

type VerifyParam struct {
	Body string
	Key  string
	Sign string
}

func Sign(param SignParam) string {
	signStr := fmt.Sprintf("%s.%s.%s", param.Key, param.Body, param.Key)
	bs := md5.Sum([]byte(signStr))
	rtn := fmt.Sprintf("%x", bs)
	return rtn
}

func Verify(param VerifyParam) bool {
	return Sign(SignParam{Body: param.Body, Key: param.Key}) == param.Sign
}
