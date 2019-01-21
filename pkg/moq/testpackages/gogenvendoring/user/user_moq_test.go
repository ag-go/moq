// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package user

import (
	"github.com/matryer/somerepo"
	"sync"
)

var (
	lockServiceMockDoSomething sync.RWMutex
)

// ServiceMock is a mock implementation of Service.
//
//     func TestSomethingThatUsesService(t *testing.T) {
//
//         // make and configure a mocked Service
//         mockedService := &ServiceMock{
//             DoSomethingFunc: func(in1 somerepo.SomeType) error {
// 	               panic("mock out the DoSomething method")
//             },
//         }
//
//         // use mockedService in code that requires Service
//         // and then make assertions.
//
//     }
type ServiceMock struct {
	// DoSomethingFunc mocks the DoSomething method.
	DoSomethingFunc func(in1 somerepo.SomeType) error

	// calls tracks calls to the methods.
	calls struct {
		// DoSomething holds details about calls to the DoSomething method.
		DoSomething []struct {
			// In1 is the in1 argument value.
			In1 somerepo.SomeType
		}
	}
}

// DoSomething calls DoSomethingFunc.
func (mock *ServiceMock) DoSomething(in1 somerepo.SomeType) error {
	if mock.DoSomethingFunc == nil {
		panic("ServiceMock.DoSomethingFunc: method is nil but Service.DoSomething was just called")
	}
	callInfo := struct {
		In1 somerepo.SomeType
	}{
		In1: in1,
	}
	lockServiceMockDoSomething.Lock()
	mock.calls.DoSomething = append(mock.calls.DoSomething, callInfo)
	lockServiceMockDoSomething.Unlock()
	return mock.DoSomethingFunc(in1)
}

// DoSomethingCalls gets all the calls that were made to DoSomething.
// Check the length with:
//     len(mockedService.DoSomethingCalls())
func (mock *ServiceMock) DoSomethingCalls() []struct {
	In1 somerepo.SomeType
} {
	var calls []struct {
		In1 somerepo.SomeType
	}
	lockServiceMockDoSomething.RLock()
	calls = mock.calls.DoSomething
	lockServiceMockDoSomething.RUnlock()
	return calls
}