package config

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	TokyoLocation           = "Asia/Tokyo"
	CusTimeFormat           = "2006-01-02T15:04:05"
	EndPoint                = "https://api.rms.rakuten.co.jp"
	SearchOrderPath         = "/es/2.0/order/searchOrder/"
	GetOrderPath            = "/es/2.0/order/getOrder/"
	SearchItemPath          = "/es/2.0/items/search"
	BulkGetInventoryPath    = "/es/2.1/inventories/bulk-get"
	UpdateOrderShippingPath = "/order/updateOrderShipping/"
)

var (
	UserAgent = fmt.Sprintf("rakuten-rms-api-sdk/v1.0.0 (Language=%s; Platform=%s-%s)", strings.Replace(runtime.Version(), "go", "go/", -1), runtime.GOOS, runtime.GOARCH)
)
