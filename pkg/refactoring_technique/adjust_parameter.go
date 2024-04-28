package refactoringtechnique

import "net/http"

type HttpRequestSource interface {
	ModifyUrl(path string) string
}

type MockHttpRequestSource struct {
	Url string
}

func (mockHttpRequestSource *MockHttpRequestSource) ModifyUrl(path string) string {
	return mockHttpRequestSource.Url + "/" + path
}

type HttpRequestSourceImpl struct {
	Req *http.Request
}

func (httpRequestSource *HttpRequestSourceImpl) ModifyUrl(path string) string {
	return httpRequestSource.Req.URL.Host + "/" + path
}

type HttpHandler struct {
	HttpRequestSource
}
