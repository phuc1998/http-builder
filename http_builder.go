package builder

import (
	_context "context"
	"fmt"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/phuc1998/http-builder/structs"
)

//ParserCustomHandle custom parser
type ParserCustomHandle func(interface{}, []byte) error

// Linger please
var (
	_ _context.Context
)

type builder struct {
	a                        *service
	uri                      string
	localVarAcceptHeader     []string
	localVarHTTPMethod       string
	localVarPostBody         interface{}
	localVarFormFileName     string
	localVarFileName         string
	localVarFileBytes        []byte
	localVarHeaderParams     map[string]string
	localVarQueryParams      _neturl.Values
	localVarFormParams       _neturl.Values
	localVarHTTPContentTypes []string
}

func (a *service) Builder(uri string, acceptHeader ...string) *builder {
	var bd = &builder{}
	bd.localVarQueryParams = _neturl.Values{}
	bd.localVarFormParams = _neturl.Values{}
	bd.localVarAcceptHeader = []string{"application/json"}
	bd.localVarAcceptHeader = append(bd.localVarAcceptHeader, acceptHeader...)
	bd.a, bd.uri = a, uri
	bd.localVarHTTPMethod = _nethttp.MethodGet
	bd.localVarHeaderParams = make(map[string]string)
	return bd
}

func (b *builder) Get() *builder {
	b.localVarHTTPMethod = _nethttp.MethodGet
	return b
}

func (b *builder) Post() *builder {
	b.localVarHTTPMethod = _nethttp.MethodPost
	return b
}

func (b *builder) Put() *builder {
	b.localVarHTTPMethod = _nethttp.MethodPut
	return b
}

func (b *builder) Delete() *builder {
	b.localVarHTTPMethod = _nethttp.MethodDelete
	return b
}

func (b *builder) SetBody(body interface{}) *builder {
	b.localVarPostBody = body
	return b
}

func (b *builder) SetHeader(key string, value interface{}) *builder {
	b.localVarHeaderParams[key] = parameterToString(value, "")
	return b
}

func (b *builder) SetBasicAuthHeader(value BasicAuth) *builder {
	b.localVarHeaderParams[AuthorizationHeader] = fmt.Sprintf(headerFormatString, BasicAuthHeader, basicAuth(value.UserName, value.Password))
	return b
}

func (b *builder) SetBearerHeader(value string) *builder {
	b.localVarHeaderParams[AuthorizationHeader] = fmt.Sprintf(headerFormatString, BearerHeader, parameterToString(value, ""))
	return b
}

func (b *builder) SetAPIKeyHeader(value APIKey) *builder {
	b.localVarHeaderParams[value.Key] = parameterToString(value.Value, "")
	return b
}

func (b *builder) SetContentType(contentType string) *builder {
	b.localVarHTTPContentTypes = append(b.localVarHTTPContentTypes, contentType)
	return b
}

func (b *builder) SetPath(key string, value interface{}) *builder {
	b.uri = strings.ReplaceAll(b.uri, fmt.Sprintf(":%s", key), parameterToString(value, ""))
	return b
}

func (b *builder) SetFormFileName(formFileName string) *builder {
	b.localVarFormFileName = formFileName
	return b
}

func (b *builder) SetFileName(fileName string) *builder {
	b.localVarFileName = fileName
	return b
}

func (b *builder) SetFileBytes(files []byte) *builder {
	b.localVarFileBytes = files
	return b
}

func (b *builder) SetQuery(key string, value interface{}) *builder {
	b.localVarQueryParams.Add(key, parameterToString(value, ""))
	return b
}

func (b *builder) SetFormParam(key string, value interface{}) *builder {
	b.localVarFormParams.Add(key, parameterToString(value, ""))
	return b
}

func (b *builder) BuildQuery(queryObject interface{}) *builder {
	var (
		structField = structs.Map(queryObject)
		queryMap    = structField["_query_"]
	)
	if queryMap != nil {
		for key, value := range queryMap.(map[string]interface{}) {
			b.localVarQueryParams.Add(key, parameterToString(value, ""))
		}
		return b
	}
	for key, value := range structField {
		b.localVarQueryParams.Add(key, parameterToString(value, ""))
	}
	return b
}

//BuildPath Replace :key by value
func (b *builder) BuildPath(pathObject interface{}) *builder {
	var (
		structField = structs.Map(pathObject)
		pathMap     = structField["_path_"]
	)
	if pathMap != nil {
		for key, value := range pathMap.(map[string]interface{}) {
			b.uri = strings.ReplaceAll(b.uri, fmt.Sprintf(":%s", key), parameterToString(value, ""))
		}
		return b
	}
	for key, value := range structField {
		b.uri = strings.ReplaceAll(b.uri, fmt.Sprintf(":%s", key), parameterToString(value, ""))
	}
	return b
}

func (b *builder) BuildHeader(headerObject interface{}) *builder {
	var (
		structField = structs.Map(headerObject)
		headerMap   = structField["_header_"]
	)
	if headerMap != nil {
		for key, value := range headerMap.(map[string]interface{}) {
			b.localVarHeaderParams[key] = parameterToString(value, "")
		}
		return b
	}
	for key, value := range structField {
		b.localVarHeaderParams[key] = parameterToString(value, "")
	}
	return b
}

func (b *builder) BuildForm(formObject interface{}) *builder {
	var (
		structField = structs.Map(formObject)
		formMap     = structField["_form_"]
	)
	if formMap != nil {
		for key, value := range formMap.(map[string]interface{}) {
			b.localVarFormParams.Add(key, parameterToString(value, ""))
		}
		return b
	}
	for key, value := range structField {
		b.localVarFormParams.Add(key, parameterToString(value, ""))
	}
	return b
}

//BuildRequest default base on http tag, not support build body
func (b *builder) BuildRequest(request interface{}) *builder {
	var (
		structField = structs.Map(request)
		headerMap   = structField["_header_"]
		queryMap    = structField["_query_"]
		pathMap     = structField["_path_"]
		formMap     = structField["_form_"]
	)
	if headerMap != nil {
		for key, value := range headerMap.(map[string]interface{}) {
			b.localVarHeaderParams[key] = parameterToString(value, "")
		}
	}
	if queryMap != nil {
		for key, value := range queryMap.(map[string]interface{}) {
			b.localVarQueryParams.Add(key, parameterToString(value, ""))
		}
	}
	if pathMap != nil {
		for key, value := range pathMap.(map[string]interface{}) {
			b.uri = strings.ReplaceAll(b.uri, fmt.Sprintf(":%s", key), parameterToString(value, ""))
		}
	}
	if formMap != nil {
		for key, value := range formMap.(map[string]interface{}) {
			b.localVarFormParams.Add(key, parameterToString(value, ""))
		}
	}
	return b
}

func (b *builder) UseXFormURLEncoded() *builder {
	b.localVarHTTPContentTypes = []string{"application/x-www-form-urlencoded"}
	return b
}

func (b *builder) UseMultipartFormData() *builder {
	b.localVarHTTPContentTypes = []string{"multipart/form-data"}
	return b
}

func (b *builder) UseApplicationJSON() *builder {
	b.localVarHTTPContentTypes = []string{"application/json"}
	return b
}

func (b *builder) UseApplicationXML() *builder {
	b.localVarHTTPContentTypes = []string{"application/xml"}
	return b
}

func (b *builder) Call(ctx _context.Context, response interface{}, parserCustom ...ParserCustomHandle) (*_nethttp.Response, error) {
	localVarPath := b.a.client.cfg.BasePath + b.uri
	localVarHTTPContentType := selectHeaderContentType(b.localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		b.localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// set Accept authorizationType
	localVarHTTPHeaderAccept := selectHeaderAccept(b.localVarAcceptHeader)
	if localVarHTTPHeaderAccept != "" {
		b.localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := b.a.client.prepareRequest(ctx, localVarPath, b.localVarHTTPMethod, b.localVarPostBody, b.localVarHeaderParams, b.localVarQueryParams, b.localVarFormParams, b.localVarFormFileName, b.localVarFileName, b.localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := b.a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	if len(parserCustom) > 0 {
		err = parserCustom[0](response, localVarBody)
		if err != nil {
			newErr := GenericOpenAPIError{
				body:  localVarBody,
				error: err.Error(),
			}
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, nil
	}

	err = b.a.client.decode(&response, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
