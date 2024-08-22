package rms

import (
	"context"
	"github.com/oh0123/rakuten-spi-sdk/config"
	"github.com/oh0123/rakuten-spi-sdk/types"
)

// reqbody
type InventoriesReqBody struct {
	Inventories []InventoryBody `json:"inventories"`
}

type InventoryBody struct {
	ManageNumber string `json:"manageNumber"`
	VariantId    string `json:"variantId"`
}

// success respbody
type InventoriesRespBody struct {
	Inventories []Inventories `json:"inventories"`
}

type Inventories struct {
	ManageNumber      string             `json:"manageNumber"`
	VariantId         string             `json:"variantId"`
	Quantity          int                `json:"quantity"`
	OperationLeadTime *OperationLeadTime `json:"operationLeadTime,omitempty"`
	ShipFromIds       *types.Slice[int]  `json:"shipFromIds,omitempty"`
	Created           string             `json:"created"`
	Updated           string             `json:"updated"`
}

type OperationLeadTime struct {
	NormalDeliveryTimeId    int `json:"normalDeliveryTimeId"`
	BackOrderDeliveryTimeId int `json:"backOrderDeliveryTimeId"`
}

// error respbody
type InventoriesErrorRespBody struct {
	Errors []InventoriesErrorDetail `json:"errors"`
}

type InventoriesErrorDetail struct {
	Code     string               `json:"code"`
	Message  string               `json:"message"`
	Metadata *InventoriesMetadata `json:"metadata"`
}

type InventoriesMetadata struct {
	PropertyPath *string `json:"propertyPath"`
}

type InventoryConfig struct {
	Auth      *types.AuthParameter
	ReqBody   *InventoriesReqBody
	RespBody  *InventoriesRespBody
	RespError *InventoriesErrorRespBody

	ClientConfig
}

func ReqInventory(ctx context.Context, cfg InventoryConfig) error {
	uri := config.EndPoint + config.BulkGetInventoryPath
	client := InitClient(&cfg.ClientConfig)
	_, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("Authorization", cfg.Auth.GenToken()).
		SetBody(cfg.ReqBody).
		SetSuccessResult(cfg.RespBody).
		SetErrorResult(cfg.RespError).
		Post(uri)
	if err != nil {
		return err
	}
	return nil
}
