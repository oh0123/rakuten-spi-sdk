package types

import (
	"encoding/base64"
	"fmt"
)

type CommonMessageModelResponse struct {
	// MessgeType はメッセージ種別です。以下のいずれかが入力されます。
	// ・INFO
	// ・ERROR
	// ・WARNING
	MessageType string `json:"messageType"`

	// MessageCode はメッセージコードです。
	// メッセージコードは https://webservice.rms.rakuten.co.jp/merchant-portal/view?contents=/ja/common/1-1_service_index/rakutenpayorderapi/rakutenpaymsgcodereference を参照してください(要ログイン)
	MessageCode string `json:"messageCode"`

	// Message はメッセージです。
	Message string `json:"message"`
}

type AuthParameter struct {
	ServiceSecret, LicenseKey string
}

func (p *AuthParameter) GenToken() string {
	return fmt.Sprintf("ESA %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", p.ServiceSecret, p.LicenseKey))))
}
