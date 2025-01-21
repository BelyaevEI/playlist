// Code generated by http://github.com/gojuno/minimock (v3.4.3). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/BelyaevEI/playlist/internal/service/auth.AuthService -o auth_service_minimock.go -n AuthServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/BelyaevEI/playlist/internal/model"
	"github.com/gojuno/minimock/v3"
)

// AuthServiceMock implements mm_auth.AuthService
type AuthServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcLogin          func(ctx context.Context, userLogin *model.UserLogin) (s1 string, err error)
	funcLoginOrigin    string
	inspectFuncLogin   func(ctx context.Context, userLogin *model.UserLogin)
	afterLoginCounter  uint64
	beforeLoginCounter uint64
	LoginMock          mAuthServiceMockLogin

	funcRegistration          func(ctx context.Context, userRegistration *model.UserRegistration) (s1 string, err error)
	funcRegistrationOrigin    string
	inspectFuncRegistration   func(ctx context.Context, userRegistration *model.UserRegistration)
	afterRegistrationCounter  uint64
	beforeRegistrationCounter uint64
	RegistrationMock          mAuthServiceMockRegistration
}

// NewAuthServiceMock returns a mock for mm_auth.AuthService
func NewAuthServiceMock(t minimock.Tester) *AuthServiceMock {
	m := &AuthServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.LoginMock = mAuthServiceMockLogin{mock: m}
	m.LoginMock.callArgs = []*AuthServiceMockLoginParams{}

	m.RegistrationMock = mAuthServiceMockRegistration{mock: m}
	m.RegistrationMock.callArgs = []*AuthServiceMockRegistrationParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mAuthServiceMockLogin struct {
	optional           bool
	mock               *AuthServiceMock
	defaultExpectation *AuthServiceMockLoginExpectation
	expectations       []*AuthServiceMockLoginExpectation

	callArgs []*AuthServiceMockLoginParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// AuthServiceMockLoginExpectation specifies expectation struct of the AuthService.Login
type AuthServiceMockLoginExpectation struct {
	mock               *AuthServiceMock
	params             *AuthServiceMockLoginParams
	paramPtrs          *AuthServiceMockLoginParamPtrs
	expectationOrigins AuthServiceMockLoginExpectationOrigins
	results            *AuthServiceMockLoginResults
	returnOrigin       string
	Counter            uint64
}

// AuthServiceMockLoginParams contains parameters of the AuthService.Login
type AuthServiceMockLoginParams struct {
	ctx       context.Context
	userLogin *model.UserLogin
}

// AuthServiceMockLoginParamPtrs contains pointers to parameters of the AuthService.Login
type AuthServiceMockLoginParamPtrs struct {
	ctx       *context.Context
	userLogin **model.UserLogin
}

// AuthServiceMockLoginResults contains results of the AuthService.Login
type AuthServiceMockLoginResults struct {
	s1  string
	err error
}

// AuthServiceMockLoginOrigins contains origins of expectations of the AuthService.Login
type AuthServiceMockLoginExpectationOrigins struct {
	origin          string
	originCtx       string
	originUserLogin string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmLogin *mAuthServiceMockLogin) Optional() *mAuthServiceMockLogin {
	mmLogin.optional = true
	return mmLogin
}

// Expect sets up expected params for AuthService.Login
func (mmLogin *mAuthServiceMockLogin) Expect(ctx context.Context, userLogin *model.UserLogin) *mAuthServiceMockLogin {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("AuthServiceMock.Login mock is already set by Set")
	}

	if mmLogin.defaultExpectation == nil {
		mmLogin.defaultExpectation = &AuthServiceMockLoginExpectation{}
	}

	if mmLogin.defaultExpectation.paramPtrs != nil {
		mmLogin.mock.t.Fatalf("AuthServiceMock.Login mock is already set by ExpectParams functions")
	}

	mmLogin.defaultExpectation.params = &AuthServiceMockLoginParams{ctx, userLogin}
	mmLogin.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmLogin.expectations {
		if minimock.Equal(e.params, mmLogin.defaultExpectation.params) {
			mmLogin.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLogin.defaultExpectation.params)
		}
	}

	return mmLogin
}

// ExpectCtxParam1 sets up expected param ctx for AuthService.Login
func (mmLogin *mAuthServiceMockLogin) ExpectCtxParam1(ctx context.Context) *mAuthServiceMockLogin {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("AuthServiceMock.Login mock is already set by Set")
	}

	if mmLogin.defaultExpectation == nil {
		mmLogin.defaultExpectation = &AuthServiceMockLoginExpectation{}
	}

	if mmLogin.defaultExpectation.params != nil {
		mmLogin.mock.t.Fatalf("AuthServiceMock.Login mock is already set by Expect")
	}

	if mmLogin.defaultExpectation.paramPtrs == nil {
		mmLogin.defaultExpectation.paramPtrs = &AuthServiceMockLoginParamPtrs{}
	}
	mmLogin.defaultExpectation.paramPtrs.ctx = &ctx
	mmLogin.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmLogin
}

// ExpectUserLoginParam2 sets up expected param userLogin for AuthService.Login
func (mmLogin *mAuthServiceMockLogin) ExpectUserLoginParam2(userLogin *model.UserLogin) *mAuthServiceMockLogin {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("AuthServiceMock.Login mock is already set by Set")
	}

	if mmLogin.defaultExpectation == nil {
		mmLogin.defaultExpectation = &AuthServiceMockLoginExpectation{}
	}

	if mmLogin.defaultExpectation.params != nil {
		mmLogin.mock.t.Fatalf("AuthServiceMock.Login mock is already set by Expect")
	}

	if mmLogin.defaultExpectation.paramPtrs == nil {
		mmLogin.defaultExpectation.paramPtrs = &AuthServiceMockLoginParamPtrs{}
	}
	mmLogin.defaultExpectation.paramPtrs.userLogin = &userLogin
	mmLogin.defaultExpectation.expectationOrigins.originUserLogin = minimock.CallerInfo(1)

	return mmLogin
}

// Inspect accepts an inspector function that has same arguments as the AuthService.Login
func (mmLogin *mAuthServiceMockLogin) Inspect(f func(ctx context.Context, userLogin *model.UserLogin)) *mAuthServiceMockLogin {
	if mmLogin.mock.inspectFuncLogin != nil {
		mmLogin.mock.t.Fatalf("Inspect function is already set for AuthServiceMock.Login")
	}

	mmLogin.mock.inspectFuncLogin = f

	return mmLogin
}

// Return sets up results that will be returned by AuthService.Login
func (mmLogin *mAuthServiceMockLogin) Return(s1 string, err error) *AuthServiceMock {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("AuthServiceMock.Login mock is already set by Set")
	}

	if mmLogin.defaultExpectation == nil {
		mmLogin.defaultExpectation = &AuthServiceMockLoginExpectation{mock: mmLogin.mock}
	}
	mmLogin.defaultExpectation.results = &AuthServiceMockLoginResults{s1, err}
	mmLogin.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmLogin.mock
}

// Set uses given function f to mock the AuthService.Login method
func (mmLogin *mAuthServiceMockLogin) Set(f func(ctx context.Context, userLogin *model.UserLogin) (s1 string, err error)) *AuthServiceMock {
	if mmLogin.defaultExpectation != nil {
		mmLogin.mock.t.Fatalf("Default expectation is already set for the AuthService.Login method")
	}

	if len(mmLogin.expectations) > 0 {
		mmLogin.mock.t.Fatalf("Some expectations are already set for the AuthService.Login method")
	}

	mmLogin.mock.funcLogin = f
	mmLogin.mock.funcLoginOrigin = minimock.CallerInfo(1)
	return mmLogin.mock
}

// When sets expectation for the AuthService.Login which will trigger the result defined by the following
// Then helper
func (mmLogin *mAuthServiceMockLogin) When(ctx context.Context, userLogin *model.UserLogin) *AuthServiceMockLoginExpectation {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("AuthServiceMock.Login mock is already set by Set")
	}

	expectation := &AuthServiceMockLoginExpectation{
		mock:               mmLogin.mock,
		params:             &AuthServiceMockLoginParams{ctx, userLogin},
		expectationOrigins: AuthServiceMockLoginExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmLogin.expectations = append(mmLogin.expectations, expectation)
	return expectation
}

// Then sets up AuthService.Login return parameters for the expectation previously defined by the When method
func (e *AuthServiceMockLoginExpectation) Then(s1 string, err error) *AuthServiceMock {
	e.results = &AuthServiceMockLoginResults{s1, err}
	return e.mock
}

// Times sets number of times AuthService.Login should be invoked
func (mmLogin *mAuthServiceMockLogin) Times(n uint64) *mAuthServiceMockLogin {
	if n == 0 {
		mmLogin.mock.t.Fatalf("Times of AuthServiceMock.Login mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmLogin.expectedInvocations, n)
	mmLogin.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmLogin
}

func (mmLogin *mAuthServiceMockLogin) invocationsDone() bool {
	if len(mmLogin.expectations) == 0 && mmLogin.defaultExpectation == nil && mmLogin.mock.funcLogin == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmLogin.mock.afterLoginCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmLogin.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Login implements mm_auth.AuthService
func (mmLogin *AuthServiceMock) Login(ctx context.Context, userLogin *model.UserLogin) (s1 string, err error) {
	mm_atomic.AddUint64(&mmLogin.beforeLoginCounter, 1)
	defer mm_atomic.AddUint64(&mmLogin.afterLoginCounter, 1)

	mmLogin.t.Helper()

	if mmLogin.inspectFuncLogin != nil {
		mmLogin.inspectFuncLogin(ctx, userLogin)
	}

	mm_params := AuthServiceMockLoginParams{ctx, userLogin}

	// Record call args
	mmLogin.LoginMock.mutex.Lock()
	mmLogin.LoginMock.callArgs = append(mmLogin.LoginMock.callArgs, &mm_params)
	mmLogin.LoginMock.mutex.Unlock()

	for _, e := range mmLogin.LoginMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmLogin.LoginMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLogin.LoginMock.defaultExpectation.Counter, 1)
		mm_want := mmLogin.LoginMock.defaultExpectation.params
		mm_want_ptrs := mmLogin.LoginMock.defaultExpectation.paramPtrs

		mm_got := AuthServiceMockLoginParams{ctx, userLogin}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmLogin.t.Errorf("AuthServiceMock.Login got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmLogin.LoginMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.userLogin != nil && !minimock.Equal(*mm_want_ptrs.userLogin, mm_got.userLogin) {
				mmLogin.t.Errorf("AuthServiceMock.Login got unexpected parameter userLogin, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmLogin.LoginMock.defaultExpectation.expectationOrigins.originUserLogin, *mm_want_ptrs.userLogin, mm_got.userLogin, minimock.Diff(*mm_want_ptrs.userLogin, mm_got.userLogin))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmLogin.t.Errorf("AuthServiceMock.Login got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmLogin.LoginMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmLogin.LoginMock.defaultExpectation.results
		if mm_results == nil {
			mmLogin.t.Fatal("No results are set for the AuthServiceMock.Login")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmLogin.funcLogin != nil {
		return mmLogin.funcLogin(ctx, userLogin)
	}
	mmLogin.t.Fatalf("Unexpected call to AuthServiceMock.Login. %v %v", ctx, userLogin)
	return
}

// LoginAfterCounter returns a count of finished AuthServiceMock.Login invocations
func (mmLogin *AuthServiceMock) LoginAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLogin.afterLoginCounter)
}

// LoginBeforeCounter returns a count of AuthServiceMock.Login invocations
func (mmLogin *AuthServiceMock) LoginBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLogin.beforeLoginCounter)
}

// Calls returns a list of arguments used in each call to AuthServiceMock.Login.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLogin *mAuthServiceMockLogin) Calls() []*AuthServiceMockLoginParams {
	mmLogin.mutex.RLock()

	argCopy := make([]*AuthServiceMockLoginParams, len(mmLogin.callArgs))
	copy(argCopy, mmLogin.callArgs)

	mmLogin.mutex.RUnlock()

	return argCopy
}

// MinimockLoginDone returns true if the count of the Login invocations corresponds
// the number of defined expectations
func (m *AuthServiceMock) MinimockLoginDone() bool {
	if m.LoginMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.LoginMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.LoginMock.invocationsDone()
}

// MinimockLoginInspect logs each unmet expectation
func (m *AuthServiceMock) MinimockLoginInspect() {
	for _, e := range m.LoginMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AuthServiceMock.Login at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterLoginCounter := mm_atomic.LoadUint64(&m.afterLoginCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.LoginMock.defaultExpectation != nil && afterLoginCounter < 1 {
		if m.LoginMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to AuthServiceMock.Login at\n%s", m.LoginMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to AuthServiceMock.Login at\n%s with params: %#v", m.LoginMock.defaultExpectation.expectationOrigins.origin, *m.LoginMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLogin != nil && afterLoginCounter < 1 {
		m.t.Errorf("Expected call to AuthServiceMock.Login at\n%s", m.funcLoginOrigin)
	}

	if !m.LoginMock.invocationsDone() && afterLoginCounter > 0 {
		m.t.Errorf("Expected %d calls to AuthServiceMock.Login at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.LoginMock.expectedInvocations), m.LoginMock.expectedInvocationsOrigin, afterLoginCounter)
	}
}

type mAuthServiceMockRegistration struct {
	optional           bool
	mock               *AuthServiceMock
	defaultExpectation *AuthServiceMockRegistrationExpectation
	expectations       []*AuthServiceMockRegistrationExpectation

	callArgs []*AuthServiceMockRegistrationParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// AuthServiceMockRegistrationExpectation specifies expectation struct of the AuthService.Registration
type AuthServiceMockRegistrationExpectation struct {
	mock               *AuthServiceMock
	params             *AuthServiceMockRegistrationParams
	paramPtrs          *AuthServiceMockRegistrationParamPtrs
	expectationOrigins AuthServiceMockRegistrationExpectationOrigins
	results            *AuthServiceMockRegistrationResults
	returnOrigin       string
	Counter            uint64
}

// AuthServiceMockRegistrationParams contains parameters of the AuthService.Registration
type AuthServiceMockRegistrationParams struct {
	ctx              context.Context
	userRegistration *model.UserRegistration
}

// AuthServiceMockRegistrationParamPtrs contains pointers to parameters of the AuthService.Registration
type AuthServiceMockRegistrationParamPtrs struct {
	ctx              *context.Context
	userRegistration **model.UserRegistration
}

// AuthServiceMockRegistrationResults contains results of the AuthService.Registration
type AuthServiceMockRegistrationResults struct {
	s1  string
	err error
}

// AuthServiceMockRegistrationOrigins contains origins of expectations of the AuthService.Registration
type AuthServiceMockRegistrationExpectationOrigins struct {
	origin                 string
	originCtx              string
	originUserRegistration string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmRegistration *mAuthServiceMockRegistration) Optional() *mAuthServiceMockRegistration {
	mmRegistration.optional = true
	return mmRegistration
}

// Expect sets up expected params for AuthService.Registration
func (mmRegistration *mAuthServiceMockRegistration) Expect(ctx context.Context, userRegistration *model.UserRegistration) *mAuthServiceMockRegistration {
	if mmRegistration.mock.funcRegistration != nil {
		mmRegistration.mock.t.Fatalf("AuthServiceMock.Registration mock is already set by Set")
	}

	if mmRegistration.defaultExpectation == nil {
		mmRegistration.defaultExpectation = &AuthServiceMockRegistrationExpectation{}
	}

	if mmRegistration.defaultExpectation.paramPtrs != nil {
		mmRegistration.mock.t.Fatalf("AuthServiceMock.Registration mock is already set by ExpectParams functions")
	}

	mmRegistration.defaultExpectation.params = &AuthServiceMockRegistrationParams{ctx, userRegistration}
	mmRegistration.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmRegistration.expectations {
		if minimock.Equal(e.params, mmRegistration.defaultExpectation.params) {
			mmRegistration.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRegistration.defaultExpectation.params)
		}
	}

	return mmRegistration
}

// ExpectCtxParam1 sets up expected param ctx for AuthService.Registration
func (mmRegistration *mAuthServiceMockRegistration) ExpectCtxParam1(ctx context.Context) *mAuthServiceMockRegistration {
	if mmRegistration.mock.funcRegistration != nil {
		mmRegistration.mock.t.Fatalf("AuthServiceMock.Registration mock is already set by Set")
	}

	if mmRegistration.defaultExpectation == nil {
		mmRegistration.defaultExpectation = &AuthServiceMockRegistrationExpectation{}
	}

	if mmRegistration.defaultExpectation.params != nil {
		mmRegistration.mock.t.Fatalf("AuthServiceMock.Registration mock is already set by Expect")
	}

	if mmRegistration.defaultExpectation.paramPtrs == nil {
		mmRegistration.defaultExpectation.paramPtrs = &AuthServiceMockRegistrationParamPtrs{}
	}
	mmRegistration.defaultExpectation.paramPtrs.ctx = &ctx
	mmRegistration.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmRegistration
}

// ExpectUserRegistrationParam2 sets up expected param userRegistration for AuthService.Registration
func (mmRegistration *mAuthServiceMockRegistration) ExpectUserRegistrationParam2(userRegistration *model.UserRegistration) *mAuthServiceMockRegistration {
	if mmRegistration.mock.funcRegistration != nil {
		mmRegistration.mock.t.Fatalf("AuthServiceMock.Registration mock is already set by Set")
	}

	if mmRegistration.defaultExpectation == nil {
		mmRegistration.defaultExpectation = &AuthServiceMockRegistrationExpectation{}
	}

	if mmRegistration.defaultExpectation.params != nil {
		mmRegistration.mock.t.Fatalf("AuthServiceMock.Registration mock is already set by Expect")
	}

	if mmRegistration.defaultExpectation.paramPtrs == nil {
		mmRegistration.defaultExpectation.paramPtrs = &AuthServiceMockRegistrationParamPtrs{}
	}
	mmRegistration.defaultExpectation.paramPtrs.userRegistration = &userRegistration
	mmRegistration.defaultExpectation.expectationOrigins.originUserRegistration = minimock.CallerInfo(1)

	return mmRegistration
}

// Inspect accepts an inspector function that has same arguments as the AuthService.Registration
func (mmRegistration *mAuthServiceMockRegistration) Inspect(f func(ctx context.Context, userRegistration *model.UserRegistration)) *mAuthServiceMockRegistration {
	if mmRegistration.mock.inspectFuncRegistration != nil {
		mmRegistration.mock.t.Fatalf("Inspect function is already set for AuthServiceMock.Registration")
	}

	mmRegistration.mock.inspectFuncRegistration = f

	return mmRegistration
}

// Return sets up results that will be returned by AuthService.Registration
func (mmRegistration *mAuthServiceMockRegistration) Return(s1 string, err error) *AuthServiceMock {
	if mmRegistration.mock.funcRegistration != nil {
		mmRegistration.mock.t.Fatalf("AuthServiceMock.Registration mock is already set by Set")
	}

	if mmRegistration.defaultExpectation == nil {
		mmRegistration.defaultExpectation = &AuthServiceMockRegistrationExpectation{mock: mmRegistration.mock}
	}
	mmRegistration.defaultExpectation.results = &AuthServiceMockRegistrationResults{s1, err}
	mmRegistration.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmRegistration.mock
}

// Set uses given function f to mock the AuthService.Registration method
func (mmRegistration *mAuthServiceMockRegistration) Set(f func(ctx context.Context, userRegistration *model.UserRegistration) (s1 string, err error)) *AuthServiceMock {
	if mmRegistration.defaultExpectation != nil {
		mmRegistration.mock.t.Fatalf("Default expectation is already set for the AuthService.Registration method")
	}

	if len(mmRegistration.expectations) > 0 {
		mmRegistration.mock.t.Fatalf("Some expectations are already set for the AuthService.Registration method")
	}

	mmRegistration.mock.funcRegistration = f
	mmRegistration.mock.funcRegistrationOrigin = minimock.CallerInfo(1)
	return mmRegistration.mock
}

// When sets expectation for the AuthService.Registration which will trigger the result defined by the following
// Then helper
func (mmRegistration *mAuthServiceMockRegistration) When(ctx context.Context, userRegistration *model.UserRegistration) *AuthServiceMockRegistrationExpectation {
	if mmRegistration.mock.funcRegistration != nil {
		mmRegistration.mock.t.Fatalf("AuthServiceMock.Registration mock is already set by Set")
	}

	expectation := &AuthServiceMockRegistrationExpectation{
		mock:               mmRegistration.mock,
		params:             &AuthServiceMockRegistrationParams{ctx, userRegistration},
		expectationOrigins: AuthServiceMockRegistrationExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmRegistration.expectations = append(mmRegistration.expectations, expectation)
	return expectation
}

// Then sets up AuthService.Registration return parameters for the expectation previously defined by the When method
func (e *AuthServiceMockRegistrationExpectation) Then(s1 string, err error) *AuthServiceMock {
	e.results = &AuthServiceMockRegistrationResults{s1, err}
	return e.mock
}

// Times sets number of times AuthService.Registration should be invoked
func (mmRegistration *mAuthServiceMockRegistration) Times(n uint64) *mAuthServiceMockRegistration {
	if n == 0 {
		mmRegistration.mock.t.Fatalf("Times of AuthServiceMock.Registration mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmRegistration.expectedInvocations, n)
	mmRegistration.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmRegistration
}

func (mmRegistration *mAuthServiceMockRegistration) invocationsDone() bool {
	if len(mmRegistration.expectations) == 0 && mmRegistration.defaultExpectation == nil && mmRegistration.mock.funcRegistration == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmRegistration.mock.afterRegistrationCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmRegistration.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Registration implements mm_auth.AuthService
func (mmRegistration *AuthServiceMock) Registration(ctx context.Context, userRegistration *model.UserRegistration) (s1 string, err error) {
	mm_atomic.AddUint64(&mmRegistration.beforeRegistrationCounter, 1)
	defer mm_atomic.AddUint64(&mmRegistration.afterRegistrationCounter, 1)

	mmRegistration.t.Helper()

	if mmRegistration.inspectFuncRegistration != nil {
		mmRegistration.inspectFuncRegistration(ctx, userRegistration)
	}

	mm_params := AuthServiceMockRegistrationParams{ctx, userRegistration}

	// Record call args
	mmRegistration.RegistrationMock.mutex.Lock()
	mmRegistration.RegistrationMock.callArgs = append(mmRegistration.RegistrationMock.callArgs, &mm_params)
	mmRegistration.RegistrationMock.mutex.Unlock()

	for _, e := range mmRegistration.RegistrationMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmRegistration.RegistrationMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRegistration.RegistrationMock.defaultExpectation.Counter, 1)
		mm_want := mmRegistration.RegistrationMock.defaultExpectation.params
		mm_want_ptrs := mmRegistration.RegistrationMock.defaultExpectation.paramPtrs

		mm_got := AuthServiceMockRegistrationParams{ctx, userRegistration}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmRegistration.t.Errorf("AuthServiceMock.Registration got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmRegistration.RegistrationMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.userRegistration != nil && !minimock.Equal(*mm_want_ptrs.userRegistration, mm_got.userRegistration) {
				mmRegistration.t.Errorf("AuthServiceMock.Registration got unexpected parameter userRegistration, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmRegistration.RegistrationMock.defaultExpectation.expectationOrigins.originUserRegistration, *mm_want_ptrs.userRegistration, mm_got.userRegistration, minimock.Diff(*mm_want_ptrs.userRegistration, mm_got.userRegistration))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRegistration.t.Errorf("AuthServiceMock.Registration got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmRegistration.RegistrationMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRegistration.RegistrationMock.defaultExpectation.results
		if mm_results == nil {
			mmRegistration.t.Fatal("No results are set for the AuthServiceMock.Registration")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmRegistration.funcRegistration != nil {
		return mmRegistration.funcRegistration(ctx, userRegistration)
	}
	mmRegistration.t.Fatalf("Unexpected call to AuthServiceMock.Registration. %v %v", ctx, userRegistration)
	return
}

// RegistrationAfterCounter returns a count of finished AuthServiceMock.Registration invocations
func (mmRegistration *AuthServiceMock) RegistrationAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegistration.afterRegistrationCounter)
}

// RegistrationBeforeCounter returns a count of AuthServiceMock.Registration invocations
func (mmRegistration *AuthServiceMock) RegistrationBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegistration.beforeRegistrationCounter)
}

// Calls returns a list of arguments used in each call to AuthServiceMock.Registration.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRegistration *mAuthServiceMockRegistration) Calls() []*AuthServiceMockRegistrationParams {
	mmRegistration.mutex.RLock()

	argCopy := make([]*AuthServiceMockRegistrationParams, len(mmRegistration.callArgs))
	copy(argCopy, mmRegistration.callArgs)

	mmRegistration.mutex.RUnlock()

	return argCopy
}

// MinimockRegistrationDone returns true if the count of the Registration invocations corresponds
// the number of defined expectations
func (m *AuthServiceMock) MinimockRegistrationDone() bool {
	if m.RegistrationMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.RegistrationMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.RegistrationMock.invocationsDone()
}

// MinimockRegistrationInspect logs each unmet expectation
func (m *AuthServiceMock) MinimockRegistrationInspect() {
	for _, e := range m.RegistrationMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AuthServiceMock.Registration at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterRegistrationCounter := mm_atomic.LoadUint64(&m.afterRegistrationCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.RegistrationMock.defaultExpectation != nil && afterRegistrationCounter < 1 {
		if m.RegistrationMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to AuthServiceMock.Registration at\n%s", m.RegistrationMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to AuthServiceMock.Registration at\n%s with params: %#v", m.RegistrationMock.defaultExpectation.expectationOrigins.origin, *m.RegistrationMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRegistration != nil && afterRegistrationCounter < 1 {
		m.t.Errorf("Expected call to AuthServiceMock.Registration at\n%s", m.funcRegistrationOrigin)
	}

	if !m.RegistrationMock.invocationsDone() && afterRegistrationCounter > 0 {
		m.t.Errorf("Expected %d calls to AuthServiceMock.Registration at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.RegistrationMock.expectedInvocations), m.RegistrationMock.expectedInvocationsOrigin, afterRegistrationCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *AuthServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockLoginInspect()

			m.MinimockRegistrationInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *AuthServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *AuthServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockLoginDone() &&
		m.MinimockRegistrationDone()
}