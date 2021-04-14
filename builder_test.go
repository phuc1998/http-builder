package builder

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type CommonResult struct {
	StatusCode int32  `json:"statusCode"`
	Message    string `json:"message"`
	Data       []Data `json:"data"`
}

type PostResponse struct {
	StatusCode int32  `json:"statusCode"`
	Message    string `json:"message"`
}

type Data struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type RequestBody struct {
	CompanyId string `json:"companyId"`
	Name      string `json:"name"`
}

type RequestCallback struct {
	Keyword     string      `http:"keyword,form"`
	UUID        string      `http:"uuid,path"`
	AccessToken string      `http:"access-token,authorizationType"`
	BookingCode string      `http:"bookingCode,query"`
	Body        RequestBody `http:"body,body"`
}

func TestPost(t *testing.T) {
	var (
		response = PostResponse{}
		request  = &RequestCallback{
			Keyword:     "ho chi minh",
			UUID:        "c68c5133-6463-49ed-9fef-4f945f5152d7",
			AccessToken: "ABC.xyz.123",
			BookingCode: "12736",
			Body: RequestBody{
				CompanyId: "cae9753a-ea8a-4ac4-acab-674ca7e673ca",
				Name:      "Phuc dep trai",
			},
		}
	)

	cfg := NewConfiguration().AddBasePath("https://http-builder.free.beeceptor.com")

	ctx := context.WithValue(context.Background(), ContextAccessToken, "key")

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail/:uuid").
		Post().
		BuildRequest(request).
		SetAPIKeyHeader(APIKey{
			Key:   "X-API-Key",
			Value: "abcdefgh123456789",
		}).
		SetBody(request.Body).
		Call(ctx, &response)

	if err != nil {
		t.Error(err)
	}
	fmt.Printf("response %v\n", response)
	t.Logf("response %v\n", response)
}

func TestGet(t *testing.T) {
	var (
		response = CommonResult{}
		request  = &RequestCallback{
			Keyword:     "ho chi minh",
			UUID:        "c68c5133-6463-49ed-9fef-4f945f5152d7",
			AccessToken: "ABC.xyz.123",
			BookingCode: "12736",
		}
	)

	cfg := NewConfiguration().AddBasePath("https://http-builder.free.beeceptor.com")

	customParser := func(resp interface{}, body []byte) error {
		result := resp.(*CommonResult)
		if err := json.Unmarshal(body, result); err != nil {
			return err
		}

		return nil
	}

	ctx := context.Background()

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail/:uuid").
		Get().
		BuildRequest(request).
		SetBearerHeader("abc.xyz.123").
		Call(ctx, &response, customParser)

	if err != nil {
		t.Error(err)
	}
	fmt.Printf("response %v\n", response)
	t.Logf("response %v\n", response)

	ctx = context.WithValue(context.Background(), ContextBasicAuth, BasicAuth{
		UserName: "phucdeptrai",
		Password: "bigphuc",
	})

	_, err = apiClient.Builder("/booking/detail/:uuid").
		Post().
		BuildRequest(request).
		SetAPIKeyHeader(APIKey{
			Key:   "X-API-Key",
			Value: "abcdefgh123456789",
		}).
		SetBody(request.Body).
		Call(ctx, &response)

	if err != nil {
		t.Error(err)
	}
	fmt.Printf("response %v\n", response)
	t.Logf("response %v\n", response)
}

type Response struct {
	StatusCode int32  `json:"statusCode"`
	Message    string `json:"message"`
}

type Request struct {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"
	ID string `http:"id,header"`
}

func main() {
	var (
		response interface{}
		basePath = "http://localhost/cars/v1"
		req      = &Request{
			ID: "123456",
		}
	)

	cfg := NewConfiguration().AddBasePath(basePath)
	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Get().
		SetBody(req).
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail

	log.Println("Response", response)
	if err != nil {
		log.Fatal(err)
	}
}
