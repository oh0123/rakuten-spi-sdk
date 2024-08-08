package rms

import (
	"rakuten-spi-sdk/config"
	"time"

	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
)

type ClientConfig struct {
	CondFunc   req.RetryConditionFunc
	RetryFunc  req.GetRetryIntervalFunc
	Timeout    time.Duration
	MaxRetries int
}

func InitClient(cfg *ClientConfig) *req.Client {
	return req.
		C().
		SetUserAgent(config.UserAgent).
		SetTimeout(cfg.Timeout).
		SetCommonRetryCount(cfg.MaxRetries).
		SetCommonRetryInterval(cfg.RetryFunc).
		SetCommonRetryCondition(cfg.CondFunc).
		SetJsonMarshal(sonic.Marshal).
		SetJsonUnmarshal(sonic.Unmarshal)
}
