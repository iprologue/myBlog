package util

import "github.com/iprologue/myBlog/pkg/setting"

func SetUp() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
