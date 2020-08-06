# http-builder

## Giới thiệu
> Thư viện ***http-builder*** giúp dễ dàng tạo một request

## Cách sử dụng

### Tạo một request Get đơn giản

```go

  	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Get().
		Call(context.Background(), response)

```
