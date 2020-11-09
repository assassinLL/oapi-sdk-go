package main

import (
	"context"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/card"
	coremodel "github.com/larksuite/oapi-sdk-go/core/model"
	"github.com/larksuite/oapi-sdk-go/core/test"
	"github.com/larksuite/oapi-sdk-go/core/tools"
)

func main() {
	var conf = test.GetISVConf("staging")
	header := make(map[string][]string)
	// from http request header
	// and "X-Lark-Request-Timestamp"/"X-Lark-Request-Nonce"/"X-Lark-Signature" validate request, Required!
	header["X-Request-Id"] = []string{"63278309j-yuewuyeu-7828389"}
	header["X-Lark-Request-Timestamp"] = []string{"Monday, 09-Nov-20 23:33:53 CST"}
	header["X-Lark-Request-Nonce"] = []string{"0404f57f-102e-4c91-bb32-a501ad0b7600"}
	header["X-Lark-Signature"] = []string{"26cb59f4f5a91c4147d0xxxxxxxxxc4a36fb2c"}
	header["X-Refresh-Token"] = []string{"acc4d5f2-4bc6-4394-a9d4-45e168fcde97"}
	req := &coremodel.OapiRequest{
		Ctx:    context.Background(),
		Header: coremodel.NewOapiHeader(header),
		Body:   "{json}",
	}
	resp := card.Handle(conf, req)
	fmt.Println(tools.Prettify(resp))
}