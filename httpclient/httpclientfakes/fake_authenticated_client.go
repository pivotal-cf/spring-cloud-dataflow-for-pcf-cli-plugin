// Code generated by counterfeiter. DO NOT EDIT.
package httpclientfakes

import (
	"io"
	"net/http"
	"sync"

	"github.com/pivotal-cf/spring-cloud-dataflow-for-pcf-cli-plugin/httpclient"
)

type FakeAuthenticatedClient struct {
	DoAuthenticatedGetStub        func(url string, accessToken string) (io.ReadCloser, int, http.Header, error)
	doAuthenticatedGetMutex       sync.RWMutex
	doAuthenticatedGetArgsForCall []struct {
		url         string
		accessToken string
	}
	doAuthenticatedGetReturns struct {
		result1 io.ReadCloser
		result2 int
		result3 http.Header
		result4 error
	}
	doAuthenticatedGetReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 int
		result3 http.Header
		result4 error
	}
	DoAuthenticatedDeleteStub        func(url string, accessToken string) (int, error)
	doAuthenticatedDeleteMutex       sync.RWMutex
	doAuthenticatedDeleteArgsForCall []struct {
		url         string
		accessToken string
	}
	doAuthenticatedDeleteReturns struct {
		result1 int
		result2 error
	}
	doAuthenticatedDeleteReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	DoAuthenticatedPostStub        func(url string, bodyType string, body string, accessToken string) (io.ReadCloser, int, error)
	doAuthenticatedPostMutex       sync.RWMutex
	doAuthenticatedPostArgsForCall []struct {
		url         string
		bodyType    string
		body        string
		accessToken string
	}
	doAuthenticatedPostReturns struct {
		result1 io.ReadCloser
		result2 int
		result3 error
	}
	doAuthenticatedPostReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 int
		result3 error
	}
	DoAuthenticatedPutStub        func(url string, accessToken string) (int, error)
	doAuthenticatedPutMutex       sync.RWMutex
	doAuthenticatedPutArgsForCall []struct {
		url         string
		accessToken string
	}
	doAuthenticatedPutReturns struct {
		result1 int
		result2 error
	}
	doAuthenticatedPutReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedGet(url string, accessToken string) (io.ReadCloser, int, http.Header, error) {
	fake.doAuthenticatedGetMutex.Lock()
	ret, specificReturn := fake.doAuthenticatedGetReturnsOnCall[len(fake.doAuthenticatedGetArgsForCall)]
	fake.doAuthenticatedGetArgsForCall = append(fake.doAuthenticatedGetArgsForCall, struct {
		url         string
		accessToken string
	}{url, accessToken})
	fake.recordInvocation("DoAuthenticatedGet", []interface{}{url, accessToken})
	fake.doAuthenticatedGetMutex.Unlock()
	if fake.DoAuthenticatedGetStub != nil {
		return fake.DoAuthenticatedGetStub(url, accessToken)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.doAuthenticatedGetReturns.result1, fake.doAuthenticatedGetReturns.result2, fake.doAuthenticatedGetReturns.result3, fake.doAuthenticatedGetReturns.result4
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedGetCallCount() int {
	fake.doAuthenticatedGetMutex.RLock()
	defer fake.doAuthenticatedGetMutex.RUnlock()
	return len(fake.doAuthenticatedGetArgsForCall)
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedGetArgsForCall(i int) (string, string) {
	fake.doAuthenticatedGetMutex.RLock()
	defer fake.doAuthenticatedGetMutex.RUnlock()
	return fake.doAuthenticatedGetArgsForCall[i].url, fake.doAuthenticatedGetArgsForCall[i].accessToken
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedGetReturns(result1 io.ReadCloser, result2 int, result3 http.Header, result4 error) {
	fake.DoAuthenticatedGetStub = nil
	fake.doAuthenticatedGetReturns = struct {
		result1 io.ReadCloser
		result2 int
		result3 http.Header
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedGetReturnsOnCall(i int, result1 io.ReadCloser, result2 int, result3 http.Header, result4 error) {
	fake.DoAuthenticatedGetStub = nil
	if fake.doAuthenticatedGetReturnsOnCall == nil {
		fake.doAuthenticatedGetReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 int
			result3 http.Header
			result4 error
		})
	}
	fake.doAuthenticatedGetReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 int
		result3 http.Header
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedDelete(url string, accessToken string) (int, error) {
	fake.doAuthenticatedDeleteMutex.Lock()
	ret, specificReturn := fake.doAuthenticatedDeleteReturnsOnCall[len(fake.doAuthenticatedDeleteArgsForCall)]
	fake.doAuthenticatedDeleteArgsForCall = append(fake.doAuthenticatedDeleteArgsForCall, struct {
		url         string
		accessToken string
	}{url, accessToken})
	fake.recordInvocation("DoAuthenticatedDelete", []interface{}{url, accessToken})
	fake.doAuthenticatedDeleteMutex.Unlock()
	if fake.DoAuthenticatedDeleteStub != nil {
		return fake.DoAuthenticatedDeleteStub(url, accessToken)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.doAuthenticatedDeleteReturns.result1, fake.doAuthenticatedDeleteReturns.result2
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedDeleteCallCount() int {
	fake.doAuthenticatedDeleteMutex.RLock()
	defer fake.doAuthenticatedDeleteMutex.RUnlock()
	return len(fake.doAuthenticatedDeleteArgsForCall)
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedDeleteArgsForCall(i int) (string, string) {
	fake.doAuthenticatedDeleteMutex.RLock()
	defer fake.doAuthenticatedDeleteMutex.RUnlock()
	return fake.doAuthenticatedDeleteArgsForCall[i].url, fake.doAuthenticatedDeleteArgsForCall[i].accessToken
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedDeleteReturns(result1 int, result2 error) {
	fake.DoAuthenticatedDeleteStub = nil
	fake.doAuthenticatedDeleteReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedDeleteReturnsOnCall(i int, result1 int, result2 error) {
	fake.DoAuthenticatedDeleteStub = nil
	if fake.doAuthenticatedDeleteReturnsOnCall == nil {
		fake.doAuthenticatedDeleteReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.doAuthenticatedDeleteReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedPost(url string, bodyType string, body string, accessToken string) (io.ReadCloser, int, error) {
	fake.doAuthenticatedPostMutex.Lock()
	ret, specificReturn := fake.doAuthenticatedPostReturnsOnCall[len(fake.doAuthenticatedPostArgsForCall)]
	fake.doAuthenticatedPostArgsForCall = append(fake.doAuthenticatedPostArgsForCall, struct {
		url         string
		bodyType    string
		body        string
		accessToken string
	}{url, bodyType, body, accessToken})
	fake.recordInvocation("DoAuthenticatedPost", []interface{}{url, bodyType, body, accessToken})
	fake.doAuthenticatedPostMutex.Unlock()
	if fake.DoAuthenticatedPostStub != nil {
		return fake.DoAuthenticatedPostStub(url, bodyType, body, accessToken)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.doAuthenticatedPostReturns.result1, fake.doAuthenticatedPostReturns.result2, fake.doAuthenticatedPostReturns.result3
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedPostCallCount() int {
	fake.doAuthenticatedPostMutex.RLock()
	defer fake.doAuthenticatedPostMutex.RUnlock()
	return len(fake.doAuthenticatedPostArgsForCall)
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedPostArgsForCall(i int) (string, string, string, string) {
	fake.doAuthenticatedPostMutex.RLock()
	defer fake.doAuthenticatedPostMutex.RUnlock()
	return fake.doAuthenticatedPostArgsForCall[i].url, fake.doAuthenticatedPostArgsForCall[i].bodyType, fake.doAuthenticatedPostArgsForCall[i].body, fake.doAuthenticatedPostArgsForCall[i].accessToken
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedPostReturns(result1 io.ReadCloser, result2 int, result3 error) {
	fake.DoAuthenticatedPostStub = nil
	fake.doAuthenticatedPostReturns = struct {
		result1 io.ReadCloser
		result2 int
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedPostReturnsOnCall(i int, result1 io.ReadCloser, result2 int, result3 error) {
	fake.DoAuthenticatedPostStub = nil
	if fake.doAuthenticatedPostReturnsOnCall == nil {
		fake.doAuthenticatedPostReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 int
			result3 error
		})
	}
	fake.doAuthenticatedPostReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 int
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedPut(url string, accessToken string) (int, error) {
	fake.doAuthenticatedPutMutex.Lock()
	ret, specificReturn := fake.doAuthenticatedPutReturnsOnCall[len(fake.doAuthenticatedPutArgsForCall)]
	fake.doAuthenticatedPutArgsForCall = append(fake.doAuthenticatedPutArgsForCall, struct {
		url         string
		accessToken string
	}{url, accessToken})
	fake.recordInvocation("DoAuthenticatedPut", []interface{}{url, accessToken})
	fake.doAuthenticatedPutMutex.Unlock()
	if fake.DoAuthenticatedPutStub != nil {
		return fake.DoAuthenticatedPutStub(url, accessToken)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.doAuthenticatedPutReturns.result1, fake.doAuthenticatedPutReturns.result2
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedPutCallCount() int {
	fake.doAuthenticatedPutMutex.RLock()
	defer fake.doAuthenticatedPutMutex.RUnlock()
	return len(fake.doAuthenticatedPutArgsForCall)
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedPutArgsForCall(i int) (string, string) {
	fake.doAuthenticatedPutMutex.RLock()
	defer fake.doAuthenticatedPutMutex.RUnlock()
	return fake.doAuthenticatedPutArgsForCall[i].url, fake.doAuthenticatedPutArgsForCall[i].accessToken
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedPutReturns(result1 int, result2 error) {
	fake.DoAuthenticatedPutStub = nil
	fake.doAuthenticatedPutReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedPutReturnsOnCall(i int, result1 int, result2 error) {
	fake.DoAuthenticatedPutStub = nil
	if fake.doAuthenticatedPutReturnsOnCall == nil {
		fake.doAuthenticatedPutReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.doAuthenticatedPutReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthenticatedClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doAuthenticatedGetMutex.RLock()
	defer fake.doAuthenticatedGetMutex.RUnlock()
	fake.doAuthenticatedDeleteMutex.RLock()
	defer fake.doAuthenticatedDeleteMutex.RUnlock()
	fake.doAuthenticatedPostMutex.RLock()
	defer fake.doAuthenticatedPostMutex.RUnlock()
	fake.doAuthenticatedPutMutex.RLock()
	defer fake.doAuthenticatedPutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthenticatedClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ httpclient.AuthenticatedClient = new(FakeAuthenticatedClient)
