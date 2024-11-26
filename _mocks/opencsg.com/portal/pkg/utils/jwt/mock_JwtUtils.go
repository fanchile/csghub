// Code generated by mockery v2.47.0. DO NOT EDIT.

package jwt

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"

	models "opencsg.com/portal/internal/models"
)

// MockJwtUtils is an autogenerated mock type for the JwtUtils type
type MockJwtUtils struct {
	mock.Mock
}

type MockJwtUtils_Expecter struct {
	mock *mock.Mock
}

func (_m *MockJwtUtils) EXPECT() *MockJwtUtils_Expecter {
	return &MockJwtUtils_Expecter{mock: &_m.Mock}
}

// GetCurrentUser provides a mock function with given fields: ctx
func (_m *MockJwtUtils) GetCurrentUser(ctx *gin.Context) *models.User {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentUser")
	}

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(*gin.Context) *models.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	return r0
}

// MockJwtUtils_GetCurrentUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrentUser'
type MockJwtUtils_GetCurrentUser_Call struct {
	*mock.Call
}

// GetCurrentUser is a helper method to define mock.On call
//   - ctx *gin.Context
func (_e *MockJwtUtils_Expecter) GetCurrentUser(ctx interface{}) *MockJwtUtils_GetCurrentUser_Call {
	return &MockJwtUtils_GetCurrentUser_Call{Call: _e.mock.On("GetCurrentUser", ctx)}
}

func (_c *MockJwtUtils_GetCurrentUser_Call) Run(run func(ctx *gin.Context)) *MockJwtUtils_GetCurrentUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockJwtUtils_GetCurrentUser_Call) Return(_a0 *models.User) *MockJwtUtils_GetCurrentUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJwtUtils_GetCurrentUser_Call) RunAndReturn(run func(*gin.Context) *models.User) *MockJwtUtils_GetCurrentUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockJwtUtils creates a new instance of MockJwtUtils. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockJwtUtils(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockJwtUtils {
	mock := &MockJwtUtils{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}