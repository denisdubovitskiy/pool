// Code generated by http://github.com/gojuno/minimock (v3.4.3). DO NOT EDIT.

package pool

//go:generate minimock -i github.com/denisdubovitskiy/pool.Dialer -o zzz_dialer_mock_test.go -n DialerMock -p pool

import (
	"net"
	"sync"
	mm_atomic "sync/atomic"
	"time"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// DialerMock implements Dialer
type DialerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcDialTimeout          func(network string, address string, timeout time.Duration) (c1 net.Conn, err error)
	funcDialTimeoutOrigin    string
	inspectFuncDialTimeout   func(network string, address string, timeout time.Duration)
	afterDialTimeoutCounter  uint64
	beforeDialTimeoutCounter uint64
	DialTimeoutMock          mDialerMockDialTimeout
}

// NewDialerMock returns a mock for Dialer
func NewDialerMock(t minimock.Tester) *DialerMock {
	m := &DialerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DialTimeoutMock = mDialerMockDialTimeout{mock: m}
	m.DialTimeoutMock.callArgs = []*DialerMockDialTimeoutParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mDialerMockDialTimeout struct {
	optional           bool
	mock               *DialerMock
	defaultExpectation *DialerMockDialTimeoutExpectation
	expectations       []*DialerMockDialTimeoutExpectation

	callArgs []*DialerMockDialTimeoutParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// DialerMockDialTimeoutExpectation specifies expectation struct of the Dialer.DialTimeout
type DialerMockDialTimeoutExpectation struct {
	mock               *DialerMock
	params             *DialerMockDialTimeoutParams
	paramPtrs          *DialerMockDialTimeoutParamPtrs
	expectationOrigins DialerMockDialTimeoutExpectationOrigins
	results            *DialerMockDialTimeoutResults
	returnOrigin       string
	Counter            uint64
}

// DialerMockDialTimeoutParams contains parameters of the Dialer.DialTimeout
type DialerMockDialTimeoutParams struct {
	network string
	address string
	timeout time.Duration
}

// DialerMockDialTimeoutParamPtrs contains pointers to parameters of the Dialer.DialTimeout
type DialerMockDialTimeoutParamPtrs struct {
	network *string
	address *string
	timeout *time.Duration
}

// DialerMockDialTimeoutResults contains results of the Dialer.DialTimeout
type DialerMockDialTimeoutResults struct {
	c1  net.Conn
	err error
}

// DialerMockDialTimeoutOrigins contains origins of expectations of the Dialer.DialTimeout
type DialerMockDialTimeoutExpectationOrigins struct {
	origin        string
	originNetwork string
	originAddress string
	originTimeout string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmDialTimeout *mDialerMockDialTimeout) Optional() *mDialerMockDialTimeout {
	mmDialTimeout.optional = true
	return mmDialTimeout
}

// Expect sets up expected params for Dialer.DialTimeout
func (mmDialTimeout *mDialerMockDialTimeout) Expect(network string, address string, timeout time.Duration) *mDialerMockDialTimeout {
	if mmDialTimeout.mock.funcDialTimeout != nil {
		mmDialTimeout.mock.t.Fatalf("DialerMock.DialTimeout mock is already set by Set")
	}

	if mmDialTimeout.defaultExpectation == nil {
		mmDialTimeout.defaultExpectation = &DialerMockDialTimeoutExpectation{}
	}

	if mmDialTimeout.defaultExpectation.paramPtrs != nil {
		mmDialTimeout.mock.t.Fatalf("DialerMock.DialTimeout mock is already set by ExpectParams functions")
	}

	mmDialTimeout.defaultExpectation.params = &DialerMockDialTimeoutParams{network, address, timeout}
	mmDialTimeout.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmDialTimeout.expectations {
		if minimock.Equal(e.params, mmDialTimeout.defaultExpectation.params) {
			mmDialTimeout.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDialTimeout.defaultExpectation.params)
		}
	}

	return mmDialTimeout
}

// ExpectNetworkParam1 sets up expected param network for Dialer.DialTimeout
func (mmDialTimeout *mDialerMockDialTimeout) ExpectNetworkParam1(network string) *mDialerMockDialTimeout {
	if mmDialTimeout.mock.funcDialTimeout != nil {
		mmDialTimeout.mock.t.Fatalf("DialerMock.DialTimeout mock is already set by Set")
	}

	if mmDialTimeout.defaultExpectation == nil {
		mmDialTimeout.defaultExpectation = &DialerMockDialTimeoutExpectation{}
	}

	if mmDialTimeout.defaultExpectation.params != nil {
		mmDialTimeout.mock.t.Fatalf("DialerMock.DialTimeout mock is already set by Expect")
	}

	if mmDialTimeout.defaultExpectation.paramPtrs == nil {
		mmDialTimeout.defaultExpectation.paramPtrs = &DialerMockDialTimeoutParamPtrs{}
	}
	mmDialTimeout.defaultExpectation.paramPtrs.network = &network
	mmDialTimeout.defaultExpectation.expectationOrigins.originNetwork = minimock.CallerInfo(1)

	return mmDialTimeout
}

// ExpectAddressParam2 sets up expected param address for Dialer.DialTimeout
func (mmDialTimeout *mDialerMockDialTimeout) ExpectAddressParam2(address string) *mDialerMockDialTimeout {
	if mmDialTimeout.mock.funcDialTimeout != nil {
		mmDialTimeout.mock.t.Fatalf("DialerMock.DialTimeout mock is already set by Set")
	}

	if mmDialTimeout.defaultExpectation == nil {
		mmDialTimeout.defaultExpectation = &DialerMockDialTimeoutExpectation{}
	}

	if mmDialTimeout.defaultExpectation.params != nil {
		mmDialTimeout.mock.t.Fatalf("DialerMock.DialTimeout mock is already set by Expect")
	}

	if mmDialTimeout.defaultExpectation.paramPtrs == nil {
		mmDialTimeout.defaultExpectation.paramPtrs = &DialerMockDialTimeoutParamPtrs{}
	}
	mmDialTimeout.defaultExpectation.paramPtrs.address = &address
	mmDialTimeout.defaultExpectation.expectationOrigins.originAddress = minimock.CallerInfo(1)

	return mmDialTimeout
}

// ExpectTimeoutParam3 sets up expected param timeout for Dialer.DialTimeout
func (mmDialTimeout *mDialerMockDialTimeout) ExpectTimeoutParam3(timeout time.Duration) *mDialerMockDialTimeout {
	if mmDialTimeout.mock.funcDialTimeout != nil {
		mmDialTimeout.mock.t.Fatalf("DialerMock.DialTimeout mock is already set by Set")
	}

	if mmDialTimeout.defaultExpectation == nil {
		mmDialTimeout.defaultExpectation = &DialerMockDialTimeoutExpectation{}
	}

	if mmDialTimeout.defaultExpectation.params != nil {
		mmDialTimeout.mock.t.Fatalf("DialerMock.DialTimeout mock is already set by Expect")
	}

	if mmDialTimeout.defaultExpectation.paramPtrs == nil {
		mmDialTimeout.defaultExpectation.paramPtrs = &DialerMockDialTimeoutParamPtrs{}
	}
	mmDialTimeout.defaultExpectation.paramPtrs.timeout = &timeout
	mmDialTimeout.defaultExpectation.expectationOrigins.originTimeout = minimock.CallerInfo(1)

	return mmDialTimeout
}

// Inspect accepts an inspector function that has same arguments as the Dialer.DialTimeout
func (mmDialTimeout *mDialerMockDialTimeout) Inspect(f func(network string, address string, timeout time.Duration)) *mDialerMockDialTimeout {
	if mmDialTimeout.mock.inspectFuncDialTimeout != nil {
		mmDialTimeout.mock.t.Fatalf("Inspect function is already set for DialerMock.DialTimeout")
	}

	mmDialTimeout.mock.inspectFuncDialTimeout = f

	return mmDialTimeout
}

// Return sets up results that will be returned by Dialer.DialTimeout
func (mmDialTimeout *mDialerMockDialTimeout) Return(c1 net.Conn, err error) *DialerMock {
	if mmDialTimeout.mock.funcDialTimeout != nil {
		mmDialTimeout.mock.t.Fatalf("DialerMock.DialTimeout mock is already set by Set")
	}

	if mmDialTimeout.defaultExpectation == nil {
		mmDialTimeout.defaultExpectation = &DialerMockDialTimeoutExpectation{mock: mmDialTimeout.mock}
	}
	mmDialTimeout.defaultExpectation.results = &DialerMockDialTimeoutResults{c1, err}
	mmDialTimeout.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmDialTimeout.mock
}

// Set uses given function f to mock the Dialer.DialTimeout method
func (mmDialTimeout *mDialerMockDialTimeout) Set(f func(network string, address string, timeout time.Duration) (c1 net.Conn, err error)) *DialerMock {
	if mmDialTimeout.defaultExpectation != nil {
		mmDialTimeout.mock.t.Fatalf("Default expectation is already set for the Dialer.DialTimeout method")
	}

	if len(mmDialTimeout.expectations) > 0 {
		mmDialTimeout.mock.t.Fatalf("Some expectations are already set for the Dialer.DialTimeout method")
	}

	mmDialTimeout.mock.funcDialTimeout = f
	mmDialTimeout.mock.funcDialTimeoutOrigin = minimock.CallerInfo(1)
	return mmDialTimeout.mock
}

// When sets expectation for the Dialer.DialTimeout which will trigger the result defined by the following
// Then helper
func (mmDialTimeout *mDialerMockDialTimeout) When(network string, address string, timeout time.Duration) *DialerMockDialTimeoutExpectation {
	if mmDialTimeout.mock.funcDialTimeout != nil {
		mmDialTimeout.mock.t.Fatalf("DialerMock.DialTimeout mock is already set by Set")
	}

	expectation := &DialerMockDialTimeoutExpectation{
		mock:               mmDialTimeout.mock,
		params:             &DialerMockDialTimeoutParams{network, address, timeout},
		expectationOrigins: DialerMockDialTimeoutExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmDialTimeout.expectations = append(mmDialTimeout.expectations, expectation)
	return expectation
}

// Then sets up Dialer.DialTimeout return parameters for the expectation previously defined by the When method
func (e *DialerMockDialTimeoutExpectation) Then(c1 net.Conn, err error) *DialerMock {
	e.results = &DialerMockDialTimeoutResults{c1, err}
	return e.mock
}

// Times sets number of times Dialer.DialTimeout should be invoked
func (mmDialTimeout *mDialerMockDialTimeout) Times(n uint64) *mDialerMockDialTimeout {
	if n == 0 {
		mmDialTimeout.mock.t.Fatalf("Times of DialerMock.DialTimeout mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmDialTimeout.expectedInvocations, n)
	mmDialTimeout.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmDialTimeout
}

func (mmDialTimeout *mDialerMockDialTimeout) invocationsDone() bool {
	if len(mmDialTimeout.expectations) == 0 && mmDialTimeout.defaultExpectation == nil && mmDialTimeout.mock.funcDialTimeout == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmDialTimeout.mock.afterDialTimeoutCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmDialTimeout.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// DialTimeout implements Dialer
func (mmDialTimeout *DialerMock) DialTimeout(network string, address string, timeout time.Duration) (c1 net.Conn, err error) {
	mm_atomic.AddUint64(&mmDialTimeout.beforeDialTimeoutCounter, 1)
	defer mm_atomic.AddUint64(&mmDialTimeout.afterDialTimeoutCounter, 1)

	mmDialTimeout.t.Helper()

	if mmDialTimeout.inspectFuncDialTimeout != nil {
		mmDialTimeout.inspectFuncDialTimeout(network, address, timeout)
	}

	mm_params := DialerMockDialTimeoutParams{network, address, timeout}

	// Record call args
	mmDialTimeout.DialTimeoutMock.mutex.Lock()
	mmDialTimeout.DialTimeoutMock.callArgs = append(mmDialTimeout.DialTimeoutMock.callArgs, &mm_params)
	mmDialTimeout.DialTimeoutMock.mutex.Unlock()

	for _, e := range mmDialTimeout.DialTimeoutMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.c1, e.results.err
		}
	}

	if mmDialTimeout.DialTimeoutMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDialTimeout.DialTimeoutMock.defaultExpectation.Counter, 1)
		mm_want := mmDialTimeout.DialTimeoutMock.defaultExpectation.params
		mm_want_ptrs := mmDialTimeout.DialTimeoutMock.defaultExpectation.paramPtrs

		mm_got := DialerMockDialTimeoutParams{network, address, timeout}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.network != nil && !minimock.Equal(*mm_want_ptrs.network, mm_got.network) {
				mmDialTimeout.t.Errorf("DialerMock.DialTimeout got unexpected parameter network, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmDialTimeout.DialTimeoutMock.defaultExpectation.expectationOrigins.originNetwork, *mm_want_ptrs.network, mm_got.network, minimock.Diff(*mm_want_ptrs.network, mm_got.network))
			}

			if mm_want_ptrs.address != nil && !minimock.Equal(*mm_want_ptrs.address, mm_got.address) {
				mmDialTimeout.t.Errorf("DialerMock.DialTimeout got unexpected parameter address, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmDialTimeout.DialTimeoutMock.defaultExpectation.expectationOrigins.originAddress, *mm_want_ptrs.address, mm_got.address, minimock.Diff(*mm_want_ptrs.address, mm_got.address))
			}

			if mm_want_ptrs.timeout != nil && !minimock.Equal(*mm_want_ptrs.timeout, mm_got.timeout) {
				mmDialTimeout.t.Errorf("DialerMock.DialTimeout got unexpected parameter timeout, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmDialTimeout.DialTimeoutMock.defaultExpectation.expectationOrigins.originTimeout, *mm_want_ptrs.timeout, mm_got.timeout, minimock.Diff(*mm_want_ptrs.timeout, mm_got.timeout))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDialTimeout.t.Errorf("DialerMock.DialTimeout got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmDialTimeout.DialTimeoutMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDialTimeout.DialTimeoutMock.defaultExpectation.results
		if mm_results == nil {
			mmDialTimeout.t.Fatal("No results are set for the DialerMock.DialTimeout")
		}
		return (*mm_results).c1, (*mm_results).err
	}
	if mmDialTimeout.funcDialTimeout != nil {
		return mmDialTimeout.funcDialTimeout(network, address, timeout)
	}
	mmDialTimeout.t.Fatalf("Unexpected call to DialerMock.DialTimeout. %v %v %v", network, address, timeout)
	return
}

// DialTimeoutAfterCounter returns a count of finished DialerMock.DialTimeout invocations
func (mmDialTimeout *DialerMock) DialTimeoutAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDialTimeout.afterDialTimeoutCounter)
}

// DialTimeoutBeforeCounter returns a count of DialerMock.DialTimeout invocations
func (mmDialTimeout *DialerMock) DialTimeoutBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDialTimeout.beforeDialTimeoutCounter)
}

// Calls returns a list of arguments used in each call to DialerMock.DialTimeout.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDialTimeout *mDialerMockDialTimeout) Calls() []*DialerMockDialTimeoutParams {
	mmDialTimeout.mutex.RLock()

	argCopy := make([]*DialerMockDialTimeoutParams, len(mmDialTimeout.callArgs))
	copy(argCopy, mmDialTimeout.callArgs)

	mmDialTimeout.mutex.RUnlock()

	return argCopy
}

// MinimockDialTimeoutDone returns true if the count of the DialTimeout invocations corresponds
// the number of defined expectations
func (m *DialerMock) MinimockDialTimeoutDone() bool {
	if m.DialTimeoutMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.DialTimeoutMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.DialTimeoutMock.invocationsDone()
}

// MinimockDialTimeoutInspect logs each unmet expectation
func (m *DialerMock) MinimockDialTimeoutInspect() {
	for _, e := range m.DialTimeoutMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DialerMock.DialTimeout at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterDialTimeoutCounter := mm_atomic.LoadUint64(&m.afterDialTimeoutCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.DialTimeoutMock.defaultExpectation != nil && afterDialTimeoutCounter < 1 {
		if m.DialTimeoutMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to DialerMock.DialTimeout at\n%s", m.DialTimeoutMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to DialerMock.DialTimeout at\n%s with params: %#v", m.DialTimeoutMock.defaultExpectation.expectationOrigins.origin, *m.DialTimeoutMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDialTimeout != nil && afterDialTimeoutCounter < 1 {
		m.t.Errorf("Expected call to DialerMock.DialTimeout at\n%s", m.funcDialTimeoutOrigin)
	}

	if !m.DialTimeoutMock.invocationsDone() && afterDialTimeoutCounter > 0 {
		m.t.Errorf("Expected %d calls to DialerMock.DialTimeout at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.DialTimeoutMock.expectedInvocations), m.DialTimeoutMock.expectedInvocationsOrigin, afterDialTimeoutCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DialerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockDialTimeoutInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DialerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *DialerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDialTimeoutDone()
}
