// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package minipanic

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// GetterMock implements Getter
type GetterMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGet          func() (err error)
	inspectFuncGet   func()
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mGetterMockGet

	funcGet2          func()
	inspectFuncGet2   func()
	afterGet2Counter  uint64
	beforeGet2Counter uint64
	Get2Mock          mGetterMockGet2
}

// NewGetterMock returns a mock for Getter
func NewGetterMock(t minimock.Tester) *GetterMock {
	m := &GetterMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetMock = mGetterMockGet{mock: m}

	m.Get2Mock = mGetterMockGet2{mock: m}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mGetterMockGet struct {
	mock               *GetterMock
	defaultExpectation *GetterMockGetExpectation
	expectations       []*GetterMockGetExpectation
}

// GetterMockGetExpectation specifies expectation struct of the Getter.Get
type GetterMockGetExpectation struct {
	mock *GetterMock

	results *GetterMockGetResults
	Counter uint64
}

// GetterMockGetResults contains results of the Getter.Get
type GetterMockGetResults struct {
	err error
}

// Expect sets up expected params for Getter.Get
func (mmGet *mGetterMockGet) Expect() *mGetterMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("GetterMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &GetterMockGetExpectation{}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the Getter.Get
func (mmGet *mGetterMockGet) Inspect(f func()) *mGetterMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for GetterMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by Getter.Get
func (mmGet *mGetterMockGet) Return(err error) *GetterMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("GetterMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &GetterMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &GetterMockGetResults{err}
	return mmGet.mock
}

// Set uses given function f to mock the Getter.Get method
func (mmGet *mGetterMockGet) Set(f func() (err error)) *GetterMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the Getter.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the Getter.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// Get implements Getter
func (mmGet *GetterMock) Get() (err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet()
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the GetterMock.Get")
		}
		return (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet()
	}
	mmGet.t.Fatalf("Unexpected call to GetterMock.Get.")
	return
}

// GetAfterCounter returns a count of finished GetterMock.Get invocations
func (mmGet *GetterMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of GetterMock.Get invocations
func (mmGet *GetterMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *GetterMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *GetterMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to GetterMock.Get")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to GetterMock.Get")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to GetterMock.Get")
	}
}

type mGetterMockGet2 struct {
	mock               *GetterMock
	defaultExpectation *GetterMockGet2Expectation
	expectations       []*GetterMockGet2Expectation
}

// GetterMockGet2Expectation specifies expectation struct of the Getter.Get2
type GetterMockGet2Expectation struct {
	mock *GetterMock

	Counter uint64
}

// Expect sets up expected params for Getter.Get2
func (mmGet2 *mGetterMockGet2) Expect() *mGetterMockGet2 {
	if mmGet2.mock.funcGet2 != nil {
		mmGet2.mock.t.Fatalf("GetterMock.Get2 mock is already set by Set")
	}

	if mmGet2.defaultExpectation == nil {
		mmGet2.defaultExpectation = &GetterMockGet2Expectation{}
	}

	return mmGet2
}

// Inspect accepts an inspector function that has same arguments as the Getter.Get2
func (mmGet2 *mGetterMockGet2) Inspect(f func()) *mGetterMockGet2 {
	if mmGet2.mock.inspectFuncGet2 != nil {
		mmGet2.mock.t.Fatalf("Inspect function is already set for GetterMock.Get2")
	}

	mmGet2.mock.inspectFuncGet2 = f

	return mmGet2
}

// Return sets up results that will be returned by Getter.Get2
func (mmGet2 *mGetterMockGet2) Return() *GetterMock {
	if mmGet2.mock.funcGet2 != nil {
		mmGet2.mock.t.Fatalf("GetterMock.Get2 mock is already set by Set")
	}

	if mmGet2.defaultExpectation == nil {
		mmGet2.defaultExpectation = &GetterMockGet2Expectation{mock: mmGet2.mock}
	}

	return mmGet2.mock
}

// Set uses given function f to mock the Getter.Get2 method
func (mmGet2 *mGetterMockGet2) Set(f func()) *GetterMock {
	if mmGet2.defaultExpectation != nil {
		mmGet2.mock.t.Fatalf("Default expectation is already set for the Getter.Get2 method")
	}

	if len(mmGet2.expectations) > 0 {
		mmGet2.mock.t.Fatalf("Some expectations are already set for the Getter.Get2 method")
	}

	mmGet2.mock.funcGet2 = f
	return mmGet2.mock
}

// Get2 implements Getter
func (mmGet2 *GetterMock) Get2() {
	mm_atomic.AddUint64(&mmGet2.beforeGet2Counter, 1)
	defer mm_atomic.AddUint64(&mmGet2.afterGet2Counter, 1)

	if mmGet2.inspectFuncGet2 != nil {
		mmGet2.inspectFuncGet2()
	}

	if mmGet2.Get2Mock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet2.Get2Mock.defaultExpectation.Counter, 1)

		return

	}
	if mmGet2.funcGet2 != nil {
		mmGet2.funcGet2()
		return
	}
	mmGet2.t.Fatalf("Unexpected call to GetterMock.Get2.")

}

// Get2AfterCounter returns a count of finished GetterMock.Get2 invocations
func (mmGet2 *GetterMock) Get2AfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet2.afterGet2Counter)
}

// Get2BeforeCounter returns a count of GetterMock.Get2 invocations
func (mmGet2 *GetterMock) Get2BeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet2.beforeGet2Counter)
}

// MinimockGet2Done returns true if the count of the Get2 invocations corresponds
// the number of defined expectations
func (m *GetterMock) MinimockGet2Done() bool {
	for _, e := range m.Get2Mock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.Get2Mock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGet2Counter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet2 != nil && mm_atomic.LoadUint64(&m.afterGet2Counter) < 1 {
		return false
	}
	return true
}

// MinimockGet2Inspect logs each unmet expectation
func (m *GetterMock) MinimockGet2Inspect() {
	for _, e := range m.Get2Mock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to GetterMock.Get2")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.Get2Mock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGet2Counter) < 1 {
		m.t.Error("Expected call to GetterMock.Get2")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet2 != nil && mm_atomic.LoadUint64(&m.afterGet2Counter) < 1 {
		m.t.Error("Expected call to GetterMock.Get2")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *GetterMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetInspect()

			m.MinimockGet2Inspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *GetterMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *GetterMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetDone() &&
		m.MinimockGet2Done()
}