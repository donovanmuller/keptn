// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"github.com/keptn/keptn/secret-service/pkg/model"
	"github.com/keptn/keptn/secret-service/pkg/repository"
	"sync"
)

// Ensure, that ScopesRepositoryMock does implement repository.ScopesRepository.
// If this is not the case, regenerate this file with moq.
var _ repository.ScopesRepository = &ScopesRepositoryMock{}

// ScopesRepositoryMock is a mock implementation of repository.ScopesRepository.
//
// 	func TestSomethingThatUsesScopesRepository(t *testing.T) {
//
// 		// make and configure a mocked repository.ScopesRepository
// 		mockedScopesRepository := &ScopesRepositoryMock{
// 			ReadFunc: func() (model.Scopes, error) {
// 				panic("mock out the Read method")
// 			},
// 		}
//
// 		// use mockedScopesRepository in code that requires repository.ScopesRepository
// 		// and then make assertions.
//
// 	}
type ScopesRepositoryMock struct {
	// ReadFunc mocks the Read method.
	ReadFunc func() (model.Scopes, error)

	// calls tracks calls to the methods.
	calls struct {
		// Read holds details about calls to the Read method.
		Read []struct {
		}
	}
	lockRead sync.RWMutex
}

// Read calls ReadFunc.
func (mock *ScopesRepositoryMock) Read() (model.Scopes, error) {
	if mock.ReadFunc == nil {
		panic("ScopesRepositoryMock.ReadFunc: method is nil but ScopesRepository.Read was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRead.Lock()
	mock.calls.Read = append(mock.calls.Read, callInfo)
	mock.lockRead.Unlock()
	return mock.ReadFunc()
}

// ReadCalls gets all the calls that were made to Read.
// Check the length with:
//     len(mockedScopesRepository.ReadCalls())
func (mock *ScopesRepositoryMock) ReadCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRead.RLock()
	calls = mock.calls.Read
	mock.lockRead.RUnlock()
	return calls
}
