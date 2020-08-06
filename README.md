# http-builder

## Giới thiệu
> Thư viện ***http-builder*** giúp dễ dàng tạo một request

## Cách sử dụng

### Tạo một request đơn giản

```go

func() {
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Get(). //support Get, Post, Put, Delete (Default: Get)
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail
}
  	
```

### Tạo request phức tạp

- Mặc định request sẽ sử dụng ContentType là "application/json"

> Tạo request với header (dùng hàm ***SetHeader***(key, value))

```go

func() {
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Get().
		SetHeader("id", "123456").
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail
}
  	
```

> Tạo request với header (dùng hàm ***BuildHeader***(object interface{}))

```go

type Request struct {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"
	ID string `http:"id,header"` 
}

func() {
	req := &Request{
		ID: "123456",
	}
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Get().
		BuildHeader(req).
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail
}
  	
```

> Tạo request với query param (dùng hàm ***SetQuery***(key, value))

```go

func() {
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Get().
		SetQuery("id", "123456").
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail?id=123456
}
  	
```

> Tạo request với query param (dùng hàm ***BuildQuery***(object))

```go

type Request struct {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"
	ID string `http:"id,query"` 
}

func() {
	req := &Request{
		ID: "123456",
	}
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Get().
		BuildQuery(req).
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail?id=123456
}
  	
```

> Tạo request với path param (dùng hàm ***SetPath***(key, value))

```go

func() {
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail/:id").
		Get().
		SetPath("id", "123456").
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail/123456
}
  	
```

> Tạo request với path param (dùng hàm ***BuildPath***(object))

```go

type Request struct {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"
	ID string `http:"id,path"` 
}

func() {
	req := &Request{
		ID: "123456",
	}
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail/:id").
		Get().
		BuildPath(req).
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail/123456
}
  	
```


> Tạo request với form data (dùng hàm ***SetFormParam***(key, value)) 

- Nếu muốn request sử dụng ContentType là "multipart/form-data" thì dùng hàm ***UseMultipartFormData()***, nếu muốn dùng "application/x-www-form-urlencoded" thì sử dụng hàm ***UseXFormURLEncoded()*** hoặc ***SetContentType(contentType)*** trước khi ***Call***

```go

func() {
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Get().
        	SetFormParam("id", "123456").
        	UseMultipartFormData().
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail
}
  	
```

> Tạo request với form data (dùng hàm ***BuildForm***(object))

```go

type Request struct {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"
	ID string `http:"id,form"` 
}

func() {
	req := &Request{
		ID: "123456",
	}
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Get().
        	BuildForm(req).
        	UseMultipartFormData().
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail?id=123456
}
  	
```

> Tạo request với body data (dùng hàm ***SetBody***(body)) 

```go
type RequestBody struct {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"
	ID string `json:"id"` 
}

func() {
	req := &RequestBody{
		ID: "123456",
	}
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Post().
		Setbody(body).
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail
}
  	
```

> Tạo request với body data (dùng hàm ***BuildBody***(object))

```go

type RequestBody struct {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"
	ID string `json:"id"` 
}


type Request struct {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"
	Body RequestBody `http:"body,body"` //`http:",body"`
}

func() {
	req := &Request{
		Body: RequestBody{
			ID: "123456",
		},
	}
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Post().
		BuildBody(req).
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail
}
  	
```

> Tạo request với nhiều yếu tố (dùng hàm ***BuildRequest***(object))

```go

type RequestBody struct {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"
	ID string `json:"id"` 
}


type Request struct {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"
	Muid string 	`http:"muid,header"`
	KeyWord string `http:"keyword,query"`
	Body RequestBody `http:"body,body"` //`http:",body"`
}

func() {
	req := &Request{
		Muid: "abc",
		Keyword: "cba",
		Body: RequestBody{
			ID: "123456",
		},
	}
    	cfg := NewConfiguration()
	cfg.BasePath = "http://localhost/cars/v1"
	cfg.HTTPClient = http.DefaultClient

	apiClient := NewAPIClient(cfg)
	_, err := apiClient.Builder("/booking/detail").
		Post().
		BuildRequest(req).
		Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail?keyword=cba
}
  	
```
