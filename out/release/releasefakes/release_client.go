// This file was generated by counterfeiter
package releasefakes

import (
	"sync"

	"github.com/pivotal-cf/go-pivnet"
)

type ReleaseClient struct {
	EULAsStub        func() ([]pivnet.EULA, error)
	eULAsMutex       sync.RWMutex
	eULAsArgsForCall []struct{}
	eULAsReturns     struct {
		result1 []pivnet.EULA
		result2 error
	}
	ReleaseTypesStub        func() ([]pivnet.ReleaseType, error)
	releaseTypesMutex       sync.RWMutex
	releaseTypesArgsForCall []struct{}
	releaseTypesReturns     struct {
		result1 []pivnet.ReleaseType
		result2 error
	}
	ReleasesForProductSlugStub        func(string) ([]pivnet.Release, error)
	releasesForProductSlugMutex       sync.RWMutex
	releasesForProductSlugArgsForCall []struct {
		arg1 string
	}
	releasesForProductSlugReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	CreateReleaseStub        func(pivnet.CreateReleaseConfig) (pivnet.Release, error)
	createReleaseMutex       sync.RWMutex
	createReleaseArgsForCall []struct {
		arg1 pivnet.CreateReleaseConfig
	}
	createReleaseReturns struct {
		result1 pivnet.Release
		result2 error
	}
	ProductVersionsStub        func(productSlug string, releases []pivnet.Release) ([]string, error)
	productVersionsMutex       sync.RWMutex
	productVersionsArgsForCall []struct {
		productSlug string
		releases    []pivnet.Release
	}
	productVersionsReturns struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseClient) EULAs() ([]pivnet.EULA, error) {
	fake.eULAsMutex.Lock()
	fake.eULAsArgsForCall = append(fake.eULAsArgsForCall, struct{}{})
	fake.recordInvocation("EULAs", []interface{}{})
	fake.eULAsMutex.Unlock()
	if fake.EULAsStub != nil {
		return fake.EULAsStub()
	} else {
		return fake.eULAsReturns.result1, fake.eULAsReturns.result2
	}
}

func (fake *ReleaseClient) EULAsCallCount() int {
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	return len(fake.eULAsArgsForCall)
}

func (fake *ReleaseClient) EULAsReturns(result1 []pivnet.EULA, result2 error) {
	fake.EULAsStub = nil
	fake.eULAsReturns = struct {
		result1 []pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *ReleaseClient) ReleaseTypes() ([]pivnet.ReleaseType, error) {
	fake.releaseTypesMutex.Lock()
	fake.releaseTypesArgsForCall = append(fake.releaseTypesArgsForCall, struct{}{})
	fake.recordInvocation("ReleaseTypes", []interface{}{})
	fake.releaseTypesMutex.Unlock()
	if fake.ReleaseTypesStub != nil {
		return fake.ReleaseTypesStub()
	} else {
		return fake.releaseTypesReturns.result1, fake.releaseTypesReturns.result2
	}
}

func (fake *ReleaseClient) ReleaseTypesCallCount() int {
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	return len(fake.releaseTypesArgsForCall)
}

func (fake *ReleaseClient) ReleaseTypesReturns(result1 []pivnet.ReleaseType, result2 error) {
	fake.ReleaseTypesStub = nil
	fake.releaseTypesReturns = struct {
		result1 []pivnet.ReleaseType
		result2 error
	}{result1, result2}
}

func (fake *ReleaseClient) ReleasesForProductSlug(arg1 string) ([]pivnet.Release, error) {
	fake.releasesForProductSlugMutex.Lock()
	fake.releasesForProductSlugArgsForCall = append(fake.releasesForProductSlugArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReleasesForProductSlug", []interface{}{arg1})
	fake.releasesForProductSlugMutex.Unlock()
	if fake.ReleasesForProductSlugStub != nil {
		return fake.ReleasesForProductSlugStub(arg1)
	} else {
		return fake.releasesForProductSlugReturns.result1, fake.releasesForProductSlugReturns.result2
	}
}

func (fake *ReleaseClient) ReleasesForProductSlugCallCount() int {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return len(fake.releasesForProductSlugArgsForCall)
}

func (fake *ReleaseClient) ReleasesForProductSlugArgsForCall(i int) string {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return fake.releasesForProductSlugArgsForCall[i].arg1
}

func (fake *ReleaseClient) ReleasesForProductSlugReturns(result1 []pivnet.Release, result2 error) {
	fake.ReleasesForProductSlugStub = nil
	fake.releasesForProductSlugReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *ReleaseClient) CreateRelease(arg1 pivnet.CreateReleaseConfig) (pivnet.Release, error) {
	fake.createReleaseMutex.Lock()
	fake.createReleaseArgsForCall = append(fake.createReleaseArgsForCall, struct {
		arg1 pivnet.CreateReleaseConfig
	}{arg1})
	fake.recordInvocation("CreateRelease", []interface{}{arg1})
	fake.createReleaseMutex.Unlock()
	if fake.CreateReleaseStub != nil {
		return fake.CreateReleaseStub(arg1)
	} else {
		return fake.createReleaseReturns.result1, fake.createReleaseReturns.result2
	}
}

func (fake *ReleaseClient) CreateReleaseCallCount() int {
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	return len(fake.createReleaseArgsForCall)
}

func (fake *ReleaseClient) CreateReleaseArgsForCall(i int) pivnet.CreateReleaseConfig {
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	return fake.createReleaseArgsForCall[i].arg1
}

func (fake *ReleaseClient) CreateReleaseReturns(result1 pivnet.Release, result2 error) {
	fake.CreateReleaseStub = nil
	fake.createReleaseReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *ReleaseClient) ProductVersions(productSlug string, releases []pivnet.Release) ([]string, error) {
	var releasesCopy []pivnet.Release
	if releases != nil {
		releasesCopy = make([]pivnet.Release, len(releases))
		copy(releasesCopy, releases)
	}
	fake.productVersionsMutex.Lock()
	fake.productVersionsArgsForCall = append(fake.productVersionsArgsForCall, struct {
		productSlug string
		releases    []pivnet.Release
	}{productSlug, releasesCopy})
	fake.recordInvocation("ProductVersions", []interface{}{productSlug, releasesCopy})
	fake.productVersionsMutex.Unlock()
	if fake.ProductVersionsStub != nil {
		return fake.ProductVersionsStub(productSlug, releases)
	} else {
		return fake.productVersionsReturns.result1, fake.productVersionsReturns.result2
	}
}

func (fake *ReleaseClient) ProductVersionsCallCount() int {
	fake.productVersionsMutex.RLock()
	defer fake.productVersionsMutex.RUnlock()
	return len(fake.productVersionsArgsForCall)
}

func (fake *ReleaseClient) ProductVersionsArgsForCall(i int) (string, []pivnet.Release) {
	fake.productVersionsMutex.RLock()
	defer fake.productVersionsMutex.RUnlock()
	return fake.productVersionsArgsForCall[i].productSlug, fake.productVersionsArgsForCall[i].releases
}

func (fake *ReleaseClient) ProductVersionsReturns(result1 []string, result2 error) {
	fake.ProductVersionsStub = nil
	fake.productVersionsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *ReleaseClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	fake.productVersionsMutex.RLock()
	defer fake.productVersionsMutex.RUnlock()
	return fake.invocations
}

func (fake *ReleaseClient) recordInvocation(key string, args []interface{}) {
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
