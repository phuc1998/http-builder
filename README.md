# http-builder

## Introduction:

> **_http-builder_** is an enhanced http client which provides you a http client package to help
> your create requests easier and reusable.

**Features:**

- You can customize requests with http client build in Go.

- It supports for adding a base path for many requests in a group.

- You also add default headers into your requests.

- It supports tags for building request struct into request parameter types: query param, path
  param, form data, body and **hybrid parameters**.
- Unmarshalling responses have declared struct in json or your response custom parser.

## Installation

```shell
go get -u github.com/phuc1998/http-builder
```

## Usage:

### Create a simple request

> All configuration is default.

```go
func main() {
  var response interface{}

  cfg := NewConfiguration()
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail").
    Get().                               //support Get, Post, Put, Delete (Default: Get)
    Call(context.Background(), response) //request url: http://localhost/booking/detail

  if err != nil {
    log.Fatal(err)
  }
  log.Println("Response", response)
}
```

### Create a request with customize client config

```go
func main() {
  var (
    response interface{}
    basePath = "http://localhost/cars/v1"
  )

  proxyURL, err := url.Parse("http://proxy.bigphuc.com")
  if err != nil {
    panic(err)
  }

  // Customize http client for a proxy transport and max connection client.
  httpClient := &http.Client{
    Transport: &http.Transport{
      Proxy:           http.ProxyURL(proxyURL),
      MaxConnsPerHost: 10,
    },
    Timeout: 60,
  }

  cfg := NewConfiguration().
    // Now you call many request start with base path: http://localhost/cars/v1
    AddBasePath(basePath).
    // APIClient now can send request through a proxy with 10 max connections per host
    AddHTTPClient(httpClient).
    // This header will be added to every requests.
    AddDefaultHeader("Age", "23")

  apiClient := NewAPIClient(cfg)
  _, err = apiClient.Builder("/booking/detail").
    Get().
    Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail
                                          //             --header 'age="23"'

  if err != nil {
    log.Fatal(err)
  }
  log.Println("Response", response)
}
```

### Custom response parser:

> You can use create your own response struct and use the default parser or customize your own one.

```go
type Response struct {
	StatusCode int32  `json:"statusCode"`
	Message    string `json:"message"`
}

func main() {
  var response = Response{}

  // With request response header text/plain this parser add body into Response.Message field
  customParser := func(resp interface{}, body []byte) error {
    result := resp.(*Response)
    r := string(body)
    result.Message = r
    return nil
  }

  cfg := NewConfiguration()
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail").
    Get().
    Call(context.Background(), &response, customParser) // You can add customParser or not

  log.Println("Response", response)
  if err != nil {
    log.Fatal(err)
  }
}
```

### Create complex request

> The **_Content-Type_** header default is _application/json_.

#### Create request with header

> Use **_SetHeader_**(key, value).

```go
func() {
  var (
    response interface{}
    basePath = "http://localhost/cars/v1"
  )

  cfg := NewConfiguration().AddBasePath(basePath)
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail").
    Get().
    SetHeader("Content-Type", "text/plaint").
    Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail
                                          //             --header 'content-type="text/plaint"'

  log.Println("Response", response)
  if err != nil {
    log.Fatal(err)
  }
}

```

> Use **_BuildHeader_**(object interface{}) to build header from struct.

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

func () {
  var (
    response interface{}
    basePath = "http://localhost/cars/v1"
    req = &Request{
      ID: "123456",
    }
  )

  cfg := NewConfiguration().AddBasePath(basePath)
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail").
    Get().
    BuildHeader(req).
    Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail

  log.Println("Response", response)
  if err != nil {
    log.Fatal(err)
  }
}
```

#### Create request with query param

> Use **_SetQuery_**(key, value).

```go
func() {
  var (
    response interface{}
    basePath = "http://localhost/cars/v1"
  )

  cfg := NewConfiguration().AddBasePath(basePath)
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail").
    Get().
    SetQuery("id", "123456").
    Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail?id=123456

  log.Println("Response", response)
  if err != nil {
    log.Fatal(err)
  }
}
```

> Use **_BuildQuery_**(object) to build query param from struct.

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

func () {
  var (
    response interface{}
    basePath = "http://localhost/cars/v1"
    req = &Request{
      ID: "123456",
    }
  )

  cfg := NewConfiguration().AddBasePath(basePath)
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail").
    Get().
    BuildQuery(req).
    Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail?id=123456

  log.Println("Response", response)
  if err != nil {
    log.Fatal(err)
  }
}
```

#### Create request with path param

> Use **_SetPath_**(key, value).

```go
func() {
  var (
    response interface{}
    basePath = "http://localhost/cars/v1"
  )

  cfg := NewConfiguration().AddBasePath(basePath)
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail/:id").
    Get().
    SetPath("id", "123456").
    Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail/123456

  log.Println("Response", response)
  if err != nil {
    log.Fatal(err)
  }
}
```

> Use **_BuildPath_**(object) to build path param from struct.

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

func () {
  var (
    response interface{}
    basePath = "http://localhost/cars/v1"
    req = &Request{
      ID: "123456",
    }
  )

  cfg := NewConfiguration().AddBasePath(basePath)
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail/:id").
    Get().
    BuildPath(req).
    Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail/123456

  log.Println("Response", response)
  if err != nil {
    log.Fatal(err)
  }
}
```

#### Create request with form data

> Use **_SetFormParam_**(key, value).

If you want **_ContentType_**: "_multipart/form-data_" use **_UseMultipartFormData()_**. Moreover,
if you want "application/x-www-form-urlencoded" use **_UseXFormURLEncoded()_** or
**_SetContentType(contentType)_** before **_Call_** function.

```go
func() {
  var (
    response interface{}
    basePath = "http://localhost/cars/v1"
  )

  cfg := NewConfiguration().AddBasePath(basePath)
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail").
    Post().
    SetFormParam("id", "123456").
    UseMultipartFormData().
    Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail
                                          //             --form 'id="123456"'
  log.Println("Response", response)
  if err != nil {
    log.Fatal(err)
  }
}
```

> Use **_BuildForm_**(object) to build form param from struct.

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

func () {
  var (
    response interface{}
    basePath = "http://localhost/cars/v1"
    req = &Request{
      ID: "123456",
    }
  )

  cfg := NewConfiguration().AddBasePath(basePath)
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail").
    Post().
    BuildForm(req).
    UseMultipartFormData().
    Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail
                                          //             --form 'id="123456"'
  log.Println("Response", response)
  if err != nil {
    log.Fatal(err)
  }
}
```

#### Create request with body data

> Use **_SetBody_**(body). It does not have **_BuildBody_** function.

```go
type RequestBody struct {
  ID string `json:"id"`
}

func () {
  var (
    response interface{}
    basePath = "http://localhost/cars/v1"
    req := &RequestBody{
      ID: "123456",
    }
  )

  cfg := NewConfiguration().AddBasePath(basePath)
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail").
    Post().
    Setbody(body).
    Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail

  log.Println("Response", response)
  if err != nil {
    log.Fatal(err)
  }
}
```

### Create hybrid request

> Use **_BuildRequest_**(object) to build hybrid param from struct, but it not support for build
> body. You can use **_SetBody_** instead.

```go
type Request struct {
  Muid string    `http:"muid,header"`
  KeyWord string `http:"keyword,query"`
}

func () {
  var (
    response interface{}
    basePath = "http://localhost/cars/v1"
    req = &Request{
      Muid: "abc",
      Keyword: "cba",
    }
  )

  cfg := NewConfiguration().AddBasePath(basePath)
  apiClient := NewAPIClient(cfg)
  _, err := apiClient.Builder("/booking/detail").
    Post().
    BuildRequest(req).
    Call(context.Background(), response) //request url: http://localhost/cars/v1/booking/detail?keyword=cba

  log.Println("Response", response)
  if err != nil {
    log.Fatal(err)
  }
}
```
