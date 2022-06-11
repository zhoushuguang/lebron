package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	OSSEndpoint string
	AccessKeyID string
	AccessKeySecret string
}
