// Code generated by mockery v2.53.3. DO NOT EDIT.

package datacoord

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockManager is an autogenerated mock type for the Manager type
type MockManager struct {
	mock.Mock
}

type MockManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockManager) EXPECT() *MockManager_Expecter {
	return &MockManager_Expecter{mock: &_m.Mock}
}

// AllocNewGrowingSegment provides a mock function with given fields: ctx, req
func (_m *MockManager) AllocNewGrowingSegment(ctx context.Context, req AllocNewGrowingSegmentRequest) (*SegmentInfo, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for AllocNewGrowingSegment")
	}

	var r0 *SegmentInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, AllocNewGrowingSegmentRequest) (*SegmentInfo, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, AllocNewGrowingSegmentRequest) *SegmentInfo); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SegmentInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, AllocNewGrowingSegmentRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockManager_AllocNewGrowingSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllocNewGrowingSegment'
type MockManager_AllocNewGrowingSegment_Call struct {
	*mock.Call
}

// AllocNewGrowingSegment is a helper method to define mock.On call
//   - ctx context.Context
//   - req AllocNewGrowingSegmentRequest
func (_e *MockManager_Expecter) AllocNewGrowingSegment(ctx interface{}, req interface{}) *MockManager_AllocNewGrowingSegment_Call {
	return &MockManager_AllocNewGrowingSegment_Call{Call: _e.mock.On("AllocNewGrowingSegment", ctx, req)}
}

func (_c *MockManager_AllocNewGrowingSegment_Call) Run(run func(ctx context.Context, req AllocNewGrowingSegmentRequest)) *MockManager_AllocNewGrowingSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(AllocNewGrowingSegmentRequest))
	})
	return _c
}

func (_c *MockManager_AllocNewGrowingSegment_Call) Return(_a0 *SegmentInfo, _a1 error) *MockManager_AllocNewGrowingSegment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_AllocNewGrowingSegment_Call) RunAndReturn(run func(context.Context, AllocNewGrowingSegmentRequest) (*SegmentInfo, error)) *MockManager_AllocNewGrowingSegment_Call {
	_c.Call.Return(run)
	return _c
}

// AllocSegment provides a mock function with given fields: ctx, collectionID, partitionID, channelName, requestRows, storageVersion
func (_m *MockManager) AllocSegment(ctx context.Context, collectionID int64, partitionID int64, channelName string, requestRows int64, storageVersion int64) ([]*Allocation, error) {
	ret := _m.Called(ctx, collectionID, partitionID, channelName, requestRows, storageVersion)

	if len(ret) == 0 {
		panic("no return value specified for AllocSegment")
	}

	var r0 []*Allocation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string, int64, int64) ([]*Allocation, error)); ok {
		return rf(ctx, collectionID, partitionID, channelName, requestRows, storageVersion)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string, int64, int64) []*Allocation); ok {
		r0 = rf(ctx, collectionID, partitionID, channelName, requestRows, storageVersion)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Allocation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, string, int64, int64) error); ok {
		r1 = rf(ctx, collectionID, partitionID, channelName, requestRows, storageVersion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockManager_AllocSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllocSegment'
type MockManager_AllocSegment_Call struct {
	*mock.Call
}

// AllocSegment is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
//   - partitionID int64
//   - channelName string
//   - requestRows int64
//   - storageVersion int64
func (_e *MockManager_Expecter) AllocSegment(ctx interface{}, collectionID interface{}, partitionID interface{}, channelName interface{}, requestRows interface{}, storageVersion interface{}) *MockManager_AllocSegment_Call {
	return &MockManager_AllocSegment_Call{Call: _e.mock.On("AllocSegment", ctx, collectionID, partitionID, channelName, requestRows, storageVersion)}
}

func (_c *MockManager_AllocSegment_Call) Run(run func(ctx context.Context, collectionID int64, partitionID int64, channelName string, requestRows int64, storageVersion int64)) *MockManager_AllocSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64), args[3].(string), args[4].(int64), args[5].(int64))
	})
	return _c
}

func (_c *MockManager_AllocSegment_Call) Return(_a0 []*Allocation, _a1 error) *MockManager_AllocSegment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_AllocSegment_Call) RunAndReturn(run func(context.Context, int64, int64, string, int64, int64) ([]*Allocation, error)) *MockManager_AllocSegment_Call {
	_c.Call.Return(run)
	return _c
}

// CleanZeroSealedSegmentsOfChannel provides a mock function with given fields: ctx, channel, cpTs
func (_m *MockManager) CleanZeroSealedSegmentsOfChannel(ctx context.Context, channel string, cpTs uint64) {
	_m.Called(ctx, channel, cpTs)
}

// MockManager_CleanZeroSealedSegmentsOfChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CleanZeroSealedSegmentsOfChannel'
type MockManager_CleanZeroSealedSegmentsOfChannel_Call struct {
	*mock.Call
}

// CleanZeroSealedSegmentsOfChannel is a helper method to define mock.On call
//   - ctx context.Context
//   - channel string
//   - cpTs uint64
func (_e *MockManager_Expecter) CleanZeroSealedSegmentsOfChannel(ctx interface{}, channel interface{}, cpTs interface{}) *MockManager_CleanZeroSealedSegmentsOfChannel_Call {
	return &MockManager_CleanZeroSealedSegmentsOfChannel_Call{Call: _e.mock.On("CleanZeroSealedSegmentsOfChannel", ctx, channel, cpTs)}
}

func (_c *MockManager_CleanZeroSealedSegmentsOfChannel_Call) Run(run func(ctx context.Context, channel string, cpTs uint64)) *MockManager_CleanZeroSealedSegmentsOfChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *MockManager_CleanZeroSealedSegmentsOfChannel_Call) Return() *MockManager_CleanZeroSealedSegmentsOfChannel_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_CleanZeroSealedSegmentsOfChannel_Call) RunAndReturn(run func(context.Context, string, uint64)) *MockManager_CleanZeroSealedSegmentsOfChannel_Call {
	_c.Run(run)
	return _c
}

// DropSegment provides a mock function with given fields: ctx, channel, segmentID
func (_m *MockManager) DropSegment(ctx context.Context, channel string, segmentID int64) {
	_m.Called(ctx, channel, segmentID)
}

// MockManager_DropSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropSegment'
type MockManager_DropSegment_Call struct {
	*mock.Call
}

// DropSegment is a helper method to define mock.On call
//   - ctx context.Context
//   - channel string
//   - segmentID int64
func (_e *MockManager_Expecter) DropSegment(ctx interface{}, channel interface{}, segmentID interface{}) *MockManager_DropSegment_Call {
	return &MockManager_DropSegment_Call{Call: _e.mock.On("DropSegment", ctx, channel, segmentID)}
}

func (_c *MockManager_DropSegment_Call) Run(run func(ctx context.Context, channel string, segmentID int64)) *MockManager_DropSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockManager_DropSegment_Call) Return() *MockManager_DropSegment_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_DropSegment_Call) RunAndReturn(run func(context.Context, string, int64)) *MockManager_DropSegment_Call {
	_c.Run(run)
	return _c
}

// DropSegmentsOfChannel provides a mock function with given fields: ctx, channel
func (_m *MockManager) DropSegmentsOfChannel(ctx context.Context, channel string) {
	_m.Called(ctx, channel)
}

// MockManager_DropSegmentsOfChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropSegmentsOfChannel'
type MockManager_DropSegmentsOfChannel_Call struct {
	*mock.Call
}

// DropSegmentsOfChannel is a helper method to define mock.On call
//   - ctx context.Context
//   - channel string
func (_e *MockManager_Expecter) DropSegmentsOfChannel(ctx interface{}, channel interface{}) *MockManager_DropSegmentsOfChannel_Call {
	return &MockManager_DropSegmentsOfChannel_Call{Call: _e.mock.On("DropSegmentsOfChannel", ctx, channel)}
}

func (_c *MockManager_DropSegmentsOfChannel_Call) Run(run func(ctx context.Context, channel string)) *MockManager_DropSegmentsOfChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockManager_DropSegmentsOfChannel_Call) Return() *MockManager_DropSegmentsOfChannel_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_DropSegmentsOfChannel_Call) RunAndReturn(run func(context.Context, string)) *MockManager_DropSegmentsOfChannel_Call {
	_c.Run(run)
	return _c
}

// ExpireAllocations provides a mock function with given fields: ctx, channel, ts
func (_m *MockManager) ExpireAllocations(ctx context.Context, channel string, ts uint64) {
	_m.Called(ctx, channel, ts)
}

// MockManager_ExpireAllocations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExpireAllocations'
type MockManager_ExpireAllocations_Call struct {
	*mock.Call
}

// ExpireAllocations is a helper method to define mock.On call
//   - ctx context.Context
//   - channel string
//   - ts uint64
func (_e *MockManager_Expecter) ExpireAllocations(ctx interface{}, channel interface{}, ts interface{}) *MockManager_ExpireAllocations_Call {
	return &MockManager_ExpireAllocations_Call{Call: _e.mock.On("ExpireAllocations", ctx, channel, ts)}
}

func (_c *MockManager_ExpireAllocations_Call) Run(run func(ctx context.Context, channel string, ts uint64)) *MockManager_ExpireAllocations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *MockManager_ExpireAllocations_Call) Return() *MockManager_ExpireAllocations_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_ExpireAllocations_Call) RunAndReturn(run func(context.Context, string, uint64)) *MockManager_ExpireAllocations_Call {
	_c.Run(run)
	return _c
}

// GetFlushableSegments provides a mock function with given fields: ctx, channel, ts
func (_m *MockManager) GetFlushableSegments(ctx context.Context, channel string, ts uint64) ([]int64, error) {
	ret := _m.Called(ctx, channel, ts)

	if len(ret) == 0 {
		panic("no return value specified for GetFlushableSegments")
	}

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) ([]int64, error)); ok {
		return rf(ctx, channel, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) []int64); ok {
		r0 = rf(ctx, channel, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, uint64) error); ok {
		r1 = rf(ctx, channel, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockManager_GetFlushableSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFlushableSegments'
type MockManager_GetFlushableSegments_Call struct {
	*mock.Call
}

// GetFlushableSegments is a helper method to define mock.On call
//   - ctx context.Context
//   - channel string
//   - ts uint64
func (_e *MockManager_Expecter) GetFlushableSegments(ctx interface{}, channel interface{}, ts interface{}) *MockManager_GetFlushableSegments_Call {
	return &MockManager_GetFlushableSegments_Call{Call: _e.mock.On("GetFlushableSegments", ctx, channel, ts)}
}

func (_c *MockManager_GetFlushableSegments_Call) Run(run func(ctx context.Context, channel string, ts uint64)) *MockManager_GetFlushableSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *MockManager_GetFlushableSegments_Call) Return(_a0 []int64, _a1 error) *MockManager_GetFlushableSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_GetFlushableSegments_Call) RunAndReturn(run func(context.Context, string, uint64) ([]int64, error)) *MockManager_GetFlushableSegments_Call {
	_c.Call.Return(run)
	return _c
}

// SealAllSegments provides a mock function with given fields: ctx, channel, segIDs
func (_m *MockManager) SealAllSegments(ctx context.Context, channel string, segIDs []int64) ([]int64, error) {
	ret := _m.Called(ctx, channel, segIDs)

	if len(ret) == 0 {
		panic("no return value specified for SealAllSegments")
	}

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []int64) ([]int64, error)); ok {
		return rf(ctx, channel, segIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []int64) []int64); ok {
		r0 = rf(ctx, channel, segIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []int64) error); ok {
		r1 = rf(ctx, channel, segIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockManager_SealAllSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SealAllSegments'
type MockManager_SealAllSegments_Call struct {
	*mock.Call
}

// SealAllSegments is a helper method to define mock.On call
//   - ctx context.Context
//   - channel string
//   - segIDs []int64
func (_e *MockManager_Expecter) SealAllSegments(ctx interface{}, channel interface{}, segIDs interface{}) *MockManager_SealAllSegments_Call {
	return &MockManager_SealAllSegments_Call{Call: _e.mock.On("SealAllSegments", ctx, channel, segIDs)}
}

func (_c *MockManager_SealAllSegments_Call) Run(run func(ctx context.Context, channel string, segIDs []int64)) *MockManager_SealAllSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]int64))
	})
	return _c
}

func (_c *MockManager_SealAllSegments_Call) Return(_a0 []int64, _a1 error) *MockManager_SealAllSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_SealAllSegments_Call) RunAndReturn(run func(context.Context, string, []int64) ([]int64, error)) *MockManager_SealAllSegments_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockManager creates a new instance of MockManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockManager {
	mock := &MockManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
