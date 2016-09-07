// This file was generated by counterfeiter
package outfakes

import (
	"sync"

	"github.com/pivotal-cf/go-pivnet"
)

type Uploader struct {
	UploadStub        func(release pivnet.Release, exactGlobs []string) error
	uploadMutex       sync.RWMutex
	uploadArgsForCall []struct {
		release    pivnet.Release
		exactGlobs []string
	}
	uploadReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Uploader) Upload(release pivnet.Release, exactGlobs []string) error {
	var exactGlobsCopy []string
	if exactGlobs != nil {
		exactGlobsCopy = make([]string, len(exactGlobs))
		copy(exactGlobsCopy, exactGlobs)
	}
	fake.uploadMutex.Lock()
	fake.uploadArgsForCall = append(fake.uploadArgsForCall, struct {
		release    pivnet.Release
		exactGlobs []string
	}{release, exactGlobsCopy})
	fake.recordInvocation("Upload", []interface{}{release, exactGlobsCopy})
	fake.uploadMutex.Unlock()
	if fake.UploadStub != nil {
		return fake.UploadStub(release, exactGlobs)
	} else {
		return fake.uploadReturns.result1
	}
}

func (fake *Uploader) UploadCallCount() int {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return len(fake.uploadArgsForCall)
}

func (fake *Uploader) UploadArgsForCall(i int) (pivnet.Release, []string) {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return fake.uploadArgsForCall[i].release, fake.uploadArgsForCall[i].exactGlobs
}

func (fake *Uploader) UploadReturns(result1 error) {
	fake.UploadStub = nil
	fake.uploadReturns = struct {
		result1 error
	}{result1}
}

func (fake *Uploader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return fake.invocations
}

func (fake *Uploader) recordInvocation(key string, args []interface{}) {
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
