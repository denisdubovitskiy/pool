// Code generated by http://github.com/gojuno/minimock (v3.4.3). DO NOT EDIT.

package pool

//go:generate minimock -i github.com/denisdubovitskiy/pool.OnDialHook -o zzz_on_dial_hook_mock_test.go -n OnDialHookMock -p pool

import (
	"sync"
	mm_atomic "sync/atomic"
	"time"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// OnDialHookMock implements OnDialHook
type OnDialHookMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcOnDial          func(duration time.Duration, err error)
	funcOnDialOrigin    string
	inspectFuncOnDial   func(duration time.Duration, err error)
	afterOnDialCounter  uint64
	beforeOnDialCounter uint64
	OnDialMock          mOnDialHookMockOnDial
}

// NewOnDialHookMock returns a mock for OnDialHook
func NewOnDialHookMock(t minimock.Tester) *OnDialHookMock {
	m := &OnDialHookMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.OnDialMock = mOnDialHookMockOnDial{mock: m}
	m.OnDialMock.callArgs = []*OnDialHookMockOnDialParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mOnDialHookMockOnDial struct {
	optional           bool
	mock               *OnDialHookMock
	defaultExpectation *OnDialHookMockOnDialExpectation
	expectations       []*OnDialHookMockOnDialExpectation

	callArgs []*OnDialHookMockOnDialParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// OnDialHookMockOnDialExpectation specifies expectation struct of the OnDialHook.OnDial
type OnDialHookMockOnDialExpectation struct {
	mock               *OnDialHookMock
	params             *OnDialHookMockOnDialParams
	paramPtrs          *OnDialHookMockOnDialParamPtrs
	expectationOrigins OnDialHookMockOnDialExpectationOrigins

	returnOrigin string
	Counter      uint64
}

// OnDialHookMockOnDialParams contains parameters of the OnDialHook.OnDial
type OnDialHookMockOnDialParams struct {
	duration time.Duration
	err      error
}

// OnDialHookMockOnDialParamPtrs contains pointers to parameters of the OnDialHook.OnDial
type OnDialHookMockOnDialParamPtrs struct {
	duration *time.Duration
	err      *error
}

// OnDialHookMockOnDialOrigins contains origins of expectations of the OnDialHook.OnDial
type OnDialHookMockOnDialExpectationOrigins struct {
	origin         string
	originDuration string
	originErr      string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmOnDial *mOnDialHookMockOnDial) Optional() *mOnDialHookMockOnDial {
	mmOnDial.optional = true
	return mmOnDial
}

// Expect sets up expected params for OnDialHook.OnDial
func (mmOnDial *mOnDialHookMockOnDial) Expect(duration time.Duration, err error) *mOnDialHookMockOnDial {
	if mmOnDial.mock.funcOnDial != nil {
		mmOnDial.mock.t.Fatalf("OnDialHookMock.OnDial mock is already set by Set")
	}

	if mmOnDial.defaultExpectation == nil {
		mmOnDial.defaultExpectation = &OnDialHookMockOnDialExpectation{}
	}

	if mmOnDial.defaultExpectation.paramPtrs != nil {
		mmOnDial.mock.t.Fatalf("OnDialHookMock.OnDial mock is already set by ExpectParams functions")
	}

	mmOnDial.defaultExpectation.params = &OnDialHookMockOnDialParams{duration, err}
	mmOnDial.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmOnDial.expectations {
		if minimock.Equal(e.params, mmOnDial.defaultExpectation.params) {
			mmOnDial.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmOnDial.defaultExpectation.params)
		}
	}

	return mmOnDial
}

// ExpectDurationParam1 sets up expected param duration for OnDialHook.OnDial
func (mmOnDial *mOnDialHookMockOnDial) ExpectDurationParam1(duration time.Duration) *mOnDialHookMockOnDial {
	if mmOnDial.mock.funcOnDial != nil {
		mmOnDial.mock.t.Fatalf("OnDialHookMock.OnDial mock is already set by Set")
	}

	if mmOnDial.defaultExpectation == nil {
		mmOnDial.defaultExpectation = &OnDialHookMockOnDialExpectation{}
	}

	if mmOnDial.defaultExpectation.params != nil {
		mmOnDial.mock.t.Fatalf("OnDialHookMock.OnDial mock is already set by Expect")
	}

	if mmOnDial.defaultExpectation.paramPtrs == nil {
		mmOnDial.defaultExpectation.paramPtrs = &OnDialHookMockOnDialParamPtrs{}
	}
	mmOnDial.defaultExpectation.paramPtrs.duration = &duration
	mmOnDial.defaultExpectation.expectationOrigins.originDuration = minimock.CallerInfo(1)

	return mmOnDial
}

// ExpectErrParam2 sets up expected param err for OnDialHook.OnDial
func (mmOnDial *mOnDialHookMockOnDial) ExpectErrParam2(err error) *mOnDialHookMockOnDial {
	if mmOnDial.mock.funcOnDial != nil {
		mmOnDial.mock.t.Fatalf("OnDialHookMock.OnDial mock is already set by Set")
	}

	if mmOnDial.defaultExpectation == nil {
		mmOnDial.defaultExpectation = &OnDialHookMockOnDialExpectation{}
	}

	if mmOnDial.defaultExpectation.params != nil {
		mmOnDial.mock.t.Fatalf("OnDialHookMock.OnDial mock is already set by Expect")
	}

	if mmOnDial.defaultExpectation.paramPtrs == nil {
		mmOnDial.defaultExpectation.paramPtrs = &OnDialHookMockOnDialParamPtrs{}
	}
	mmOnDial.defaultExpectation.paramPtrs.err = &err
	mmOnDial.defaultExpectation.expectationOrigins.originErr = minimock.CallerInfo(1)

	return mmOnDial
}

// Inspect accepts an inspector function that has same arguments as the OnDialHook.OnDial
func (mmOnDial *mOnDialHookMockOnDial) Inspect(f func(duration time.Duration, err error)) *mOnDialHookMockOnDial {
	if mmOnDial.mock.inspectFuncOnDial != nil {
		mmOnDial.mock.t.Fatalf("Inspect function is already set for OnDialHookMock.OnDial")
	}

	mmOnDial.mock.inspectFuncOnDial = f

	return mmOnDial
}

// Return sets up results that will be returned by OnDialHook.OnDial
func (mmOnDial *mOnDialHookMockOnDial) Return() *OnDialHookMock {
	if mmOnDial.mock.funcOnDial != nil {
		mmOnDial.mock.t.Fatalf("OnDialHookMock.OnDial mock is already set by Set")
	}

	if mmOnDial.defaultExpectation == nil {
		mmOnDial.defaultExpectation = &OnDialHookMockOnDialExpectation{mock: mmOnDial.mock}
	}

	mmOnDial.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmOnDial.mock
}

// Set uses given function f to mock the OnDialHook.OnDial method
func (mmOnDial *mOnDialHookMockOnDial) Set(f func(duration time.Duration, err error)) *OnDialHookMock {
	if mmOnDial.defaultExpectation != nil {
		mmOnDial.mock.t.Fatalf("Default expectation is already set for the OnDialHook.OnDial method")
	}

	if len(mmOnDial.expectations) > 0 {
		mmOnDial.mock.t.Fatalf("Some expectations are already set for the OnDialHook.OnDial method")
	}

	mmOnDial.mock.funcOnDial = f
	mmOnDial.mock.funcOnDialOrigin = minimock.CallerInfo(1)
	return mmOnDial.mock
}

// Times sets number of times OnDialHook.OnDial should be invoked
func (mmOnDial *mOnDialHookMockOnDial) Times(n uint64) *mOnDialHookMockOnDial {
	if n == 0 {
		mmOnDial.mock.t.Fatalf("Times of OnDialHookMock.OnDial mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmOnDial.expectedInvocations, n)
	mmOnDial.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmOnDial
}

func (mmOnDial *mOnDialHookMockOnDial) invocationsDone() bool {
	if len(mmOnDial.expectations) == 0 && mmOnDial.defaultExpectation == nil && mmOnDial.mock.funcOnDial == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmOnDial.mock.afterOnDialCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmOnDial.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// OnDial implements OnDialHook
func (mmOnDial *OnDialHookMock) OnDial(duration time.Duration, err error) {
	mm_atomic.AddUint64(&mmOnDial.beforeOnDialCounter, 1)
	defer mm_atomic.AddUint64(&mmOnDial.afterOnDialCounter, 1)

	mmOnDial.t.Helper()

	if mmOnDial.inspectFuncOnDial != nil {
		mmOnDial.inspectFuncOnDial(duration, err)
	}

	mm_params := OnDialHookMockOnDialParams{duration, err}

	// Record call args
	mmOnDial.OnDialMock.mutex.Lock()
	mmOnDial.OnDialMock.callArgs = append(mmOnDial.OnDialMock.callArgs, &mm_params)
	mmOnDial.OnDialMock.mutex.Unlock()

	for _, e := range mmOnDial.OnDialMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmOnDial.OnDialMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmOnDial.OnDialMock.defaultExpectation.Counter, 1)
		mm_want := mmOnDial.OnDialMock.defaultExpectation.params
		mm_want_ptrs := mmOnDial.OnDialMock.defaultExpectation.paramPtrs

		mm_got := OnDialHookMockOnDialParams{duration, err}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.duration != nil && !minimock.Equal(*mm_want_ptrs.duration, mm_got.duration) {
				mmOnDial.t.Errorf("OnDialHookMock.OnDial got unexpected parameter duration, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmOnDial.OnDialMock.defaultExpectation.expectationOrigins.originDuration, *mm_want_ptrs.duration, mm_got.duration, minimock.Diff(*mm_want_ptrs.duration, mm_got.duration))
			}

			if mm_want_ptrs.err != nil && !minimock.Equal(*mm_want_ptrs.err, mm_got.err) {
				mmOnDial.t.Errorf("OnDialHookMock.OnDial got unexpected parameter err, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmOnDial.OnDialMock.defaultExpectation.expectationOrigins.originErr, *mm_want_ptrs.err, mm_got.err, minimock.Diff(*mm_want_ptrs.err, mm_got.err))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmOnDial.t.Errorf("OnDialHookMock.OnDial got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmOnDial.OnDialMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmOnDial.funcOnDial != nil {
		mmOnDial.funcOnDial(duration, err)
		return
	}
	mmOnDial.t.Fatalf("Unexpected call to OnDialHookMock.OnDial. %v %v", duration, err)

}

// OnDialAfterCounter returns a count of finished OnDialHookMock.OnDial invocations
func (mmOnDial *OnDialHookMock) OnDialAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmOnDial.afterOnDialCounter)
}

// OnDialBeforeCounter returns a count of OnDialHookMock.OnDial invocations
func (mmOnDial *OnDialHookMock) OnDialBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmOnDial.beforeOnDialCounter)
}

// Calls returns a list of arguments used in each call to OnDialHookMock.OnDial.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmOnDial *mOnDialHookMockOnDial) Calls() []*OnDialHookMockOnDialParams {
	mmOnDial.mutex.RLock()

	argCopy := make([]*OnDialHookMockOnDialParams, len(mmOnDial.callArgs))
	copy(argCopy, mmOnDial.callArgs)

	mmOnDial.mutex.RUnlock()

	return argCopy
}

// MinimockOnDialDone returns true if the count of the OnDial invocations corresponds
// the number of defined expectations
func (m *OnDialHookMock) MinimockOnDialDone() bool {
	if m.OnDialMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.OnDialMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.OnDialMock.invocationsDone()
}

// MinimockOnDialInspect logs each unmet expectation
func (m *OnDialHookMock) MinimockOnDialInspect() {
	for _, e := range m.OnDialMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OnDialHookMock.OnDial at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterOnDialCounter := mm_atomic.LoadUint64(&m.afterOnDialCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.OnDialMock.defaultExpectation != nil && afterOnDialCounter < 1 {
		if m.OnDialMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to OnDialHookMock.OnDial at\n%s", m.OnDialMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to OnDialHookMock.OnDial at\n%s with params: %#v", m.OnDialMock.defaultExpectation.expectationOrigins.origin, *m.OnDialMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcOnDial != nil && afterOnDialCounter < 1 {
		m.t.Errorf("Expected call to OnDialHookMock.OnDial at\n%s", m.funcOnDialOrigin)
	}

	if !m.OnDialMock.invocationsDone() && afterOnDialCounter > 0 {
		m.t.Errorf("Expected %d calls to OnDialHookMock.OnDial at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.OnDialMock.expectedInvocations), m.OnDialMock.expectedInvocationsOrigin, afterOnDialCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *OnDialHookMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockOnDialInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *OnDialHookMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *OnDialHookMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockOnDialDone()
}
