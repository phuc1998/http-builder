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
}

type RequestBody struct {
	Muid         string `json:"muid"`
	MAccessToken string `json:"maccesstoken"`
	CompanyId    string `json:"companyId"`
}

type RequestCallback struct {
	Muid         string      `http:"muid,header"`
	MAccessToken string      `http:"maccesstoken,header"`
	Body         RequestBody `http:"body,body"`
}

func TestGet(t *testing.T) {
	var (
		response = &CommonResult{}
		request  = &RequestCallback{

			Body: RequestBody{
				CompanyId:    "111111111111111",
				Muid:         "---jjjjj-sdfd",
				MAccessToken: "sdfdsf---jjj---",
			},
		}
	)
	cfg := NewConfiguration()
	cfg.BasePath = "http://sb-vexere.zpapps.vn/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/create_v2").Post().
		BuildRequest(request).
		Call(context.Background(), response)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	fmt.Printf("response %v", response)
}
