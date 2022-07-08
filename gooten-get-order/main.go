package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

func main() {
	// gootenID := "5847766"
	gootenID := "584776"
	gootenToken := "9d3c83751c6a4d3aa8f39996d00e9351"
	gootenOrderPriceResponse, err := GetOrderGootenPrice(gootenID, gootenToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(gootenOrderPriceResponse)
	fmt.Println(gootenOrderPriceResponse.PartnerShipping)
}

type GootenOrderPriceReponse struct {
	PartnerSubTotal float64 `json:"PartnerSubTotal"`
	PartnerTotal    float64 `json:"PartnerTotal"`
	PartnerShipping float64 `json:"PartnerShipping"`
	PartnerTax      float64 `json:"PartnerTax"`
}

func GetOrderGootenPrice(gootenID, gootenToken string) (*GootenOrderPriceReponse, error) {
	var gootenOrderPriceResponse GootenOrderPriceReponse
	url := fmt.Sprintf("https://www.gooten.com/admin/api/OrderDetailsApi/getOrderBilling?orderId=%s", gootenID)
	client := resty.New()
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+gootenToken).
		SetResult(&gootenOrderPriceResponse).
		Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed when get gooten order price")
	}
	if res.IsError() {
		return nil, errors.Errorf("status: %s, repsonse: %v", res.Status(), res.String())
	}
	if res.String() == "null" {
		return nil, nil
	}
	return &gootenOrderPriceResponse, nil
}
