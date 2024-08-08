package rms

import (
	"context"
	"rakuten-spi-sdk/config"
	"rakuten-spi-sdk/pkg"
	"rakuten-spi-sdk/types"
)

// paramters
type QueryParameters struct {
	Title                     *string       `json:"title,omitempty"`
	Tagline                   *string       `json:"tagline,omitempty"`
	ManageNumber              *string       `json:"manageNumber,omitempty"`
	ItemNumber                *string       `json:"itemNumber,omitempty"`
	ArticleNumber             *string       `json:"articleNumber,omitempty"`
	VariantId                 *string       `json:"variantId,omitempty"`
	MerchantDefinedSkuId      *string       `json:"merchantDefinedSkuId,omitempty"`
	GenreId                   *string       `json:"genreId,omitempty"`
	ItemType                  *ItemType     `json:"itemType,omitempty"`
	IsAsuraku                 *bool         `json:"isAsuraku,omitempty"`
	StandardPriceFrom         *int          `json:"standardPriceFrom,omitempty"`
	StandardPriceTo           *int          `json:"standardPriceTo,omitempty"`
	IsVariantStockout         *bool         `json:"isVariantStockout,omitempty"`
	IsItemStockout            *bool         `json:"isItemStockout,omitempty"`
	PurchasablePeriodFrom     *types.JpDate `json:"purchasablePeriodFrom,omitempty"`
	PurchasablePeriodTo       *types.JpDate `json:"purchasablePeriodTo,omitempty"`
	IsHiddenItem              *bool         `json:"isHiddenItem,omitempty"`
	IsHiddenVariant           *bool         `json:"isHiddenVariant,omitempty"`
	IsSearchable              *bool         `json:"isSearchable,omitempty"`
	IsYamiichi                *bool         `json:"isYamiichi,omitempty"`
	PointApplicablePeriodFrom *types.JpDate `json:"pointApplicablePeriodFrom,omitempty"`
	PointApplicablePeriodtTo  *types.JpDate `json:"pointApplicablePeriodTo,omitempty"`
	IsOptimizedPoint          *bool         `json:"isOptimizedPoint,omitempty"`
	PointRate                 *int          `json:"pointRate,omitempty"`
	MaxPointRate              *int          `json:"maxPointRate,omitempty"`
	CategoryId                *string       `json:"categoryId,omitempty"`
	IsBackOrder               *bool         `json:"isBackOrder,omitempty"`
	IsPostageIncluded         *bool         `json:"isPostageIncluded,omitempty"`
	CreatedFrom               *types.JpDate `json:"createdFrom,omitempty"`
	CreatedTo                 *types.JpDate `json:"createdTo,omitempty"`
	UpdatedFrom               *types.JpDate `json:"updatedFrom,omitempty"`
	UpdatedTo                 *types.JpDate `json:"updatedTo,omitempty"`
	SortKey                   *SortKey      `json:"sortKey,omitempty"`
	SortOrder                 *SortOrder    `json:"sortOrder,omitempty"`
	Offset                    *int          `json:"offset,omitempty"`
	Hits                      *int          `json:"hits,omitempty"`
	CursorMark                *string       `json:"cursorMark,omitempty"`
	IsCategoryIncluded        *bool         `json:"isCategoryIncluded,omitempty"`
	IsReviewIncluded          *bool         `json:"isReviewIncluded,omitempty"`
	IsInventoryIncluded       *bool         `json:"isInventoryIncluded,omitempty"`
}

type ItemType string

const (
	NORMAL       ItemType = "NORMAL"
	PRE_ORDER    ItemType = "PRE_ORDER"
	SUBSCRIPTION ItemType = "SUBSCRIPTION"
)

type SortOrder string

const (
	DESC SortOrder = "desc"
	ASC  SortOrder = "asc"
)

type SortKey string

const (
	UPDATED                  SortKey = "updated"
	CREATED                  SortKey = "created"
	ITEM_DISPLAY_SEQUENCE    SortKey = "itemDisplaySequence"
	MANAGE_NUMBER            SortKey = "manageNumber"
	PURCHASABLE_PERIOD_START SortKey = "purchasablePeriodStart"
	PURCHASABLE_PERIOD_END   SortKey = "purchasablePeriodEnd"
	POINT_CAMPAIGN_START     SortKey = "pointCampaignStart"
	POINT_CAMPAIGN_END       SortKey = "pointCampaignEnd"
	POINT_RATE               SortKey = "pointRate"
	REVIEW_COUNT             SortKey = "reviewCount"
	REVIEW_AVERAGE_RATING    SortKey = "reviewAverageRating"
)

// success respbody
type ItemRespBody struct {
	NumFound       int           `json:"numFound"`
	Offset         int           `json:"offset"`
	NextCursorMark string        `json:"nextCursorMark"`
	Results        []ItemResults `json:"results"`
}

type ItemResults struct {
	Item      ItemBody                `json:"item"`
	Category  *ItemCategory           `json:"category"`
	Review    *ItemReview             `json:"review"`
	Inventory *types.Map[string, any] `json:"inventory"`
}

type ItemBody struct {
	ManageNumber           string                         `json:"manageNumber"`
	ItemNumber             *string                        `json:"itemNumber,omitempty"`
	Title                  string                         `json:"title"`
	Tagline                *string                        `json:"tagline,omitempty"`
	ProductDescription     *types.Map[string, any]        `json:"productDescription,omitempty"`
	SalesDescription       *string                        `json:"salesDescription,omitempty"`
	Precautions            *types.Map[string, any]        `json:"precautions,omitempty"`
	ItemType               string                         `json:"itemType"`
	Images                 *types.Maps[string, any]       `json:"images,omitempty"`
	WhiteBgImage           *types.Map[string, any]        `json:"whiteBgImage,omitempty"`
	Video                  *types.Map[string, any]        `json:"video,omitempty"`
	GenreId                string                         `json:"genreId"`
	Tags                   *types.Slice[int]              `json:"tags,omitempty"`
	HideItem               bool                           `json:"hideItem"`
	UnlimitedInventoryFlag bool                           `json:"unlimitedInventoryFlag"`
	CustomizationOptions   *types.Maps[string, any]       `json:"customizationOptions,omitempty"`
	ReleaseDate            *string                        `json:"releaseDate,omitempty"`
	PurchasablePeriod      *types.Map[string, any]        `json:"purchasablePeriod,omitempty"`
	Subscription           *types.Map[string, any]        `json:"subscription,omitempty"`
	Features               types.Map[string, any]         `json:"features"`
	AccessControl          *types.Map[string, any]        `json:"accessControl,omitempty"`
	Payment                types.Map[string, any]         `json:"payment"`
	PointCampaign          *types.Map[string, any]        `json:"pointCampaign,omitempty"`
	ItemDisplaySequence    int                            `json:"itemDisplaySequence"`
	Layout                 types.Map[string, any]         `json:"layout"`
	VariantSelectors       *types.Maps[string, any]       `json:"variantSelectors,omitempty"`
	Variants               types.Map[string, ItemVariant] `json:"variants"`
	Created                string                         `json:"created"`
	Updated                string                         `json:"updated"`
}

type ItemCategory struct {
	CategoryIds *types.Slice[string] `json:"categoryIds"`
}

type ItemReview struct {
	Count         *int     `json:"count"`
	AverageRating *float64 `json:"averageRating"`
}

type ItemVariant struct {
	MerchantDefinedSkuId    *string                  `json:"merchantDefinedSkuId,omitempty"`
	SelectorValues          *types.Map[string, any]  `json:"selectorValues,omitempty"`
	Images                  *types.Maps[string, any] `json:"images,omitempty"`
	RestockOnCancel         *bool                    `json:"restockOnCancel,omitempty"`
	BackOrderFlag           *bool                    `json:"backOrderFlag,omitempty"`
	NormalDeliveryDateId    *int                     `json:"normalDeliveryDateId,omitempty"`
	BackOrderDeliveryDateId *int                     `json:"backOrderDeliveryDateId,omitempty"`
	OrderQuantityLimit      *int                     `json:"orderQuantityLimit,omitempty"`
	ReferencePrice          *types.Map[string, any]  `json:"referencePrice,omitempty"`
	Features                *types.Map[string, any]  `json:"features,omitempty"`
	Hidden                  *bool                    `json:"hidden,omitempty"`
	StandardPrice           *string                  `json:"standardPrice,omitempty"`
	SubscriptionPrice       *types.Map[string, any]  `json:"subscriptionPrice,omitempty"`
	ArticleNumberForSet     *types.Slice[string]     `json:"articleNumberForSet,omitempty"`
	ArticleNumber           *types.Map[string, any]  `json:"articleNumber,omitempty"`
	Shipping                *types.Map[string, any]  `json:"shipping,omitempty"`
	AsurakuDeliveryId       *int                     `json:"asurakuDeliveryId,omitempty"`
	Specs                   *types.Maps[string, any] `json:"specs,omitempty"`
	Attributes              *types.Maps[string, any] `json:"attributes,omitempty"`
}

// error respbody
type ItemErrorRespBody struct {
	Errors []ItemErrorDetail `json:"errors"`
}

type ItemErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SearchItemConfig struct {
	Auth      *types.AuthParameter
	ReqParams *QueryParameters
	RespBody  *ItemRespBody
	RespError *ItemErrorRespBody

	ClientConfig
}

func SearchItemReq(ctx context.Context, cfg SearchItemConfig) error {
	uri := config.EndPoint + config.GetOrderPath
	client := InitClient(&cfg.ClientConfig)
	params, err := pkg.StructToMap(cfg.ReqParams, "json", "")
	if err != nil {
		return err
	}
	_, err = client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("Authorization", cfg.Auth.GenToken()).
		SetQueryParamsAnyType(params).
		SetSuccessResult(cfg.RespBody).
		SetErrorResult(cfg.RespBody).
		Get(uri)
	if err != nil {
		return err
	}
	return nil
}
