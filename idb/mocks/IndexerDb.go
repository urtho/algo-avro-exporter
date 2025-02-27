// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	bookkeeping "github.com/algorand/go-algorand/data/bookkeeping"

	generated "github.com/algorand/indexer/api/generated/v2"

	idb "github.com/algorand/indexer/idb"

	mock "github.com/stretchr/testify/mock"

	transactions "github.com/algorand/go-algorand/data/transactions"
)

// IndexerDb is an autogenerated mock type for the IndexerDb type
type IndexerDb struct {
	mock.Mock
}

// AddBlock provides a mock function with given fields: block
func (_m *IndexerDb) AddBlock(block *bookkeeping.Block) error {
	ret := _m.Called(block)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bookkeeping.Block) error); ok {
		r0 = rf(block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Applications provides a mock function with given fields: ctx, filter
func (_m *IndexerDb) Applications(ctx context.Context, filter *generated.SearchForApplicationsParams) (<-chan idb.ApplicationRow, uint64) {
	ret := _m.Called(ctx, filter)

	var r0 <-chan idb.ApplicationRow
	if rf, ok := ret.Get(0).(func(context.Context, *generated.SearchForApplicationsParams) <-chan idb.ApplicationRow); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan idb.ApplicationRow)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, *generated.SearchForApplicationsParams) uint64); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}

// AssetBalances provides a mock function with given fields: ctx, abq
func (_m *IndexerDb) AssetBalances(ctx context.Context, abq idb.AssetBalanceQuery) (<-chan idb.AssetBalanceRow, uint64) {
	ret := _m.Called(ctx, abq)

	var r0 <-chan idb.AssetBalanceRow
	if rf, ok := ret.Get(0).(func(context.Context, idb.AssetBalanceQuery) <-chan idb.AssetBalanceRow); ok {
		r0 = rf(ctx, abq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan idb.AssetBalanceRow)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, idb.AssetBalanceQuery) uint64); ok {
		r1 = rf(ctx, abq)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}

// Assets provides a mock function with given fields: ctx, filter
func (_m *IndexerDb) Assets(ctx context.Context, filter idb.AssetsQuery) (<-chan idb.AssetRow, uint64) {
	ret := _m.Called(ctx, filter)

	var r0 <-chan idb.AssetRow
	if rf, ok := ret.Get(0).(func(context.Context, idb.AssetsQuery) <-chan idb.AssetRow); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan idb.AssetRow)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, idb.AssetsQuery) uint64); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *IndexerDb) Close() {
	_m.Called()
}

// GetAccounts provides a mock function with given fields: ctx, opts
func (_m *IndexerDb) GetAccounts(ctx context.Context, opts idb.AccountQueryOptions) (<-chan idb.AccountRow, uint64) {
	ret := _m.Called(ctx, opts)

	var r0 <-chan idb.AccountRow
	if rf, ok := ret.Get(0).(func(context.Context, idb.AccountQueryOptions) <-chan idb.AccountRow); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan idb.AccountRow)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, idb.AccountQueryOptions) uint64); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}

// GetBlock provides a mock function with given fields: ctx, round, options
func (_m *IndexerDb) GetBlock(ctx context.Context, round uint64, options idb.GetBlockOptions) (bookkeeping.BlockHeader, []idb.TxnRow, error) {
	ret := _m.Called(ctx, round, options)

	var r0 bookkeeping.BlockHeader
	if rf, ok := ret.Get(0).(func(context.Context, uint64, idb.GetBlockOptions) bookkeeping.BlockHeader); ok {
		r0 = rf(ctx, round, options)
	} else {
		r0 = ret.Get(0).(bookkeeping.BlockHeader)
	}

	var r1 []idb.TxnRow
	if rf, ok := ret.Get(1).(func(context.Context, uint64, idb.GetBlockOptions) []idb.TxnRow); ok {
		r1 = rf(ctx, round, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]idb.TxnRow)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, uint64, idb.GetBlockOptions) error); ok {
		r2 = rf(ctx, round, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetNextRoundToAccount provides a mock function with given fields:
func (_m *IndexerDb) GetNextRoundToAccount() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSpecialAccounts provides a mock function with given fields:
func (_m *IndexerDb) GetSpecialAccounts() (transactions.SpecialAddresses, error) {
	ret := _m.Called()

	var r0 transactions.SpecialAddresses
	if rf, ok := ret.Get(0).(func() transactions.SpecialAddresses); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(transactions.SpecialAddresses)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Health provides a mock function with given fields:
func (_m *IndexerDb) Health() (idb.Health, error) {
	ret := _m.Called()

	var r0 idb.Health
	if rf, ok := ret.Get(0).(func() idb.Health); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(idb.Health)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadGenesis provides a mock function with given fields: genesis
func (_m *IndexerDb) LoadGenesis(genesis bookkeeping.Genesis) error {
	ret := _m.Called(genesis)

	var r0 error
	if rf, ok := ret.Get(0).(func(bookkeeping.Genesis) error); ok {
		r0 = rf(genesis)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transactions provides a mock function with given fields: ctx, tf
func (_m *IndexerDb) Transactions(ctx context.Context, tf idb.TransactionFilter) (<-chan idb.TxnRow, uint64) {
	ret := _m.Called(ctx, tf)

	var r0 <-chan idb.TxnRow
	if rf, ok := ret.Get(0).(func(context.Context, idb.TransactionFilter) <-chan idb.TxnRow); ok {
		r0 = rf(ctx, tf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan idb.TxnRow)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, idb.TransactionFilter) uint64); ok {
		r1 = rf(ctx, tf)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}
