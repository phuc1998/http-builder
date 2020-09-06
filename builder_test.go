package builder

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

type CommonResult struct {
	StatusCode int32  `json:"statusCode"`
	Message    string `json:"message"`
	Data       []Data `json:"data"`
}

type Data struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type RequestBody struct {
	Muid         string `json:"muid"`
	MAccessToken string `json:"maccesstoken"`
	CompanyId    string `json:"companyId"`
}

type RequestCallback struct {
	Keyword      string      `http:"keyword,query"`
	Muid         string      `http:"muid,header"`
	MAccessToken string      `http:"maccesstoken,header"`
	BookingCode  string      `http:"bookingCode,path"`
	Body         RequestBody `http:"body,body"`
}

func TestGet(t *testing.T) {
	var (
		response = CommonResult{}
		request  = &RequestCallback{
			Keyword:      "ho chi minh",
			Muid:         "---jjjjj-sdfd",
			MAccessToken: "sdfdsf---jjj---",
			BookingCode:  "12736",
			Body: RequestBody{
				CompanyId: "111111111111111",
			},
		}
	)
	cfg := NewConfiguration()
	cfg.BasePath = "https://http-builder.free.beeceptor.com"
	cfg.HTTPClient = http.DefaultClient

	customParser := func(resp interface{}, body []byte) error {
		result := resp.(*CommonResult)
		r := string(body)
		result.Message = r
		return nil
	}

	ctx := context.WithValue(context.Background(), ContextAccessToken, "key")

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail/:bookingCode").
		Post().
		BuildRequest(request).
		SetBody(request.Body).
		Call(ctx, &response, customParser)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	fmt.Printf("response %v", response)
}
