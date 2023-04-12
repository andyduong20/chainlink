// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"
	mock "github.com/stretchr/testify/mock"

	time "time"

	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"

	types "github.com/smartcontractkit/chainlink/v2/common/types"

	uuid "github.com/satori/go.uuid"
)

// TxStore is an autogenerated mock type for the TxStore type
type TxStore[ADDR interface{}, CHAINID interface{}, TX_HASH types.Hashable[TX_HASH], BLOCK_HASH types.Hashable[BLOCK_HASH], NEWTX interface{}, R interface{}, TX interface{}, TXATTEMPT interface{}, TXID interface{}, TXMETA interface{}] struct {
	mock.Mock
}

// CheckEthTxQueueCapacity provides a mock function with given fields: fromAddress, maxQueuedTransactions, chainID, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) CheckEthTxQueueCapacity(fromAddress ADDR, maxQueuedTransactions uint64, chainID CHAINID, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fromAddress, maxQueuedTransactions, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(ADDR, uint64, CHAINID, ...pg.QOpt) error); ok {
		r0 = rf(fromAddress, maxQueuedTransactions, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) Close() {
	_m.Called()
}

// CountUnconfirmedTransactions provides a mock function with given fields: fromAddress, chainID, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) CountUnconfirmedTransactions(fromAddress ADDR, chainID CHAINID, qopts ...pg.QOpt) (uint32, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fromAddress, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, CHAINID, ...pg.QOpt) (uint32, error)); ok {
		return rf(fromAddress, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(ADDR, CHAINID, ...pg.QOpt) uint32); ok {
		r0 = rf(fromAddress, chainID, qopts...)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(ADDR, CHAINID, ...pg.QOpt) error); ok {
		r1 = rf(fromAddress, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountUnstartedTransactions provides a mock function with given fields: fromAddress, chainID, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) CountUnstartedTransactions(fromAddress ADDR, chainID CHAINID, qopts ...pg.QOpt) (uint32, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fromAddress, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, CHAINID, ...pg.QOpt) (uint32, error)); ok {
		return rf(fromAddress, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(ADDR, CHAINID, ...pg.QOpt) uint32); ok {
		r0 = rf(fromAddress, chainID, qopts...)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(ADDR, CHAINID, ...pg.QOpt) error); ok {
		r1 = rf(fromAddress, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEthTransaction provides a mock function with given fields: newTx, chainID, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) CreateEthTransaction(newTx NEWTX, chainID CHAINID, qopts ...pg.QOpt) (txmgrtypes.Transaction, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, newTx, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 txmgrtypes.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(NEWTX, CHAINID, ...pg.QOpt) (txmgrtypes.Transaction, error)); ok {
		return rf(newTx, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(NEWTX, CHAINID, ...pg.QOpt) txmgrtypes.Transaction); ok {
		r0 = rf(newTx, chainID, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(txmgrtypes.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(NEWTX, CHAINID, ...pg.QOpt) error); ok {
		r1 = rf(newTx, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteInProgressAttempt provides a mock function with given fields: ctx, attempt
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) DeleteInProgressAttempt(ctx context.Context, attempt TXATTEMPT) error {
	ret := _m.Called(ctx, attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, TXATTEMPT) error); ok {
		r0 = rf(ctx, attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EthTransactions provides a mock function with given fields: offset, limit
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) EthTransactions(offset int, limit int) ([]TX, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []TX
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]TX, int, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []TX); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TX)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EthTransactionsWithAttempts provides a mock function with given fields: offset, limit
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) EthTransactionsWithAttempts(offset int, limit int) ([]TX, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []TX
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]TX, int, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []TX); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TX)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EthTxAttempts provides a mock function with given fields: offset, limit
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) EthTxAttempts(offset int, limit int) ([]TXATTEMPT, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []TXATTEMPT
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]TXATTEMPT, int, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []TXATTEMPT); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TXATTEMPT)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindEthReceiptsPendingConfirmation provides a mock function with given fields: ctx, blockNum, chainID
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEthReceiptsPendingConfirmation(ctx context.Context, blockNum int64, chainID CHAINID) ([]txmgrtypes.ReceiptPlus[R], error) {
	ret := _m.Called(ctx, blockNum, chainID)

	var r0 []txmgrtypes.ReceiptPlus[R]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, CHAINID) ([]txmgrtypes.ReceiptPlus[R], error)); ok {
		return rf(ctx, blockNum, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, CHAINID) []txmgrtypes.ReceiptPlus[R]); ok {
		r0 = rf(ctx, blockNum, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgrtypes.ReceiptPlus[R])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, CHAINID) error); ok {
		r1 = rf(ctx, blockNum, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxAttempt provides a mock function with given fields: hash
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEthTxAttempt(hash TX_HASH) (*TXATTEMPT, error) {
	ret := _m.Called(hash)

	var r0 *TXATTEMPT
	var r1 error
	if rf, ok := ret.Get(0).(func(TX_HASH) (*TXATTEMPT, error)); ok {
		return rf(hash)
	}
	if rf, ok := ret.Get(0).(func(TX_HASH) *TXATTEMPT); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TXATTEMPT)
		}
	}

	if rf, ok := ret.Get(1).(func(TX_HASH) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxAttemptConfirmedByEthTxIDs provides a mock function with given fields: ids
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEthTxAttemptConfirmedByEthTxIDs(ids []TXID) ([]TXATTEMPT, error) {
	ret := _m.Called(ids)

	var r0 []TXATTEMPT
	var r1 error
	if rf, ok := ret.Get(0).(func([]TXID) ([]TXATTEMPT, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]TXID) []TXATTEMPT); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TXATTEMPT)
		}
	}

	if rf, ok := ret.Get(1).(func([]TXID) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxAttemptsByEthTxIDs provides a mock function with given fields: ids
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEthTxAttemptsByEthTxIDs(ids []TXID) ([]TXATTEMPT, error) {
	ret := _m.Called(ids)

	var r0 []TXATTEMPT
	var r1 error
	if rf, ok := ret.Get(0).(func([]TXID) ([]TXATTEMPT, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]TXID) []TXATTEMPT); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TXATTEMPT)
		}
	}

	if rf, ok := ret.Get(1).(func([]TXID) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxAttemptsRequiringReceiptFetch provides a mock function with given fields: chainID
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEthTxAttemptsRequiringReceiptFetch(chainID CHAINID) ([]TXATTEMPT, error) {
	ret := _m.Called(chainID)

	var r0 []TXATTEMPT
	var r1 error
	if rf, ok := ret.Get(0).(func(CHAINID) ([]TXATTEMPT, error)); ok {
		return rf(chainID)
	}
	if rf, ok := ret.Get(0).(func(CHAINID) []TXATTEMPT); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TXATTEMPT)
		}
	}

	if rf, ok := ret.Get(1).(func(CHAINID) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxAttemptsRequiringResend provides a mock function with given fields: olderThan, maxInFlightTransactions, chainID, address
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEthTxAttemptsRequiringResend(olderThan time.Time, maxInFlightTransactions uint32, chainID CHAINID, address ADDR) ([]TXATTEMPT, error) {
	ret := _m.Called(olderThan, maxInFlightTransactions, chainID, address)

	var r0 []TXATTEMPT
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, uint32, CHAINID, ADDR) ([]TXATTEMPT, error)); ok {
		return rf(olderThan, maxInFlightTransactions, chainID, address)
	}
	if rf, ok := ret.Get(0).(func(time.Time, uint32, CHAINID, ADDR) []TXATTEMPT); ok {
		r0 = rf(olderThan, maxInFlightTransactions, chainID, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TXATTEMPT)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, uint32, CHAINID, ADDR) error); ok {
		r1 = rf(olderThan, maxInFlightTransactions, chainID, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxByHash provides a mock function with given fields: hash
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEthTxByHash(hash TX_HASH) (*TX, error) {
	ret := _m.Called(hash)

	var r0 *TX
	var r1 error
	if rf, ok := ret.Get(0).(func(TX_HASH) (*TX, error)); ok {
		return rf(hash)
	}
	if rf, ok := ret.Get(0).(func(TX_HASH) *TX); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TX)
		}
	}

	if rf, ok := ret.Get(1).(func(TX_HASH) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxWithAttempts provides a mock function with given fields: etxID
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEthTxWithAttempts(etxID TXID) (TX, error) {
	ret := _m.Called(etxID)

	var r0 TX
	var r1 error
	if rf, ok := ret.Get(0).(func(TXID) (TX, error)); ok {
		return rf(etxID)
	}
	if rf, ok := ret.Get(0).(func(TXID) TX); ok {
		r0 = rf(etxID)
	} else {
		r0 = ret.Get(0).(TX)
	}

	if rf, ok := ret.Get(1).(func(TXID) error); ok {
		r1 = rf(etxID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxWithNonce provides a mock function with given fields: fromAddress, nonce
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEthTxWithNonce(fromAddress ADDR, nonce TXMETA) (*TX, error) {
	ret := _m.Called(fromAddress, nonce)

	var r0 *TX
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, TXMETA) (*TX, error)); ok {
		return rf(fromAddress, nonce)
	}
	if rf, ok := ret.Get(0).(func(ADDR, TXMETA) *TX); ok {
		r0 = rf(fromAddress, nonce)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TX)
		}
	}

	if rf, ok := ret.Get(1).(func(ADDR, TXMETA) error); ok {
		r1 = rf(fromAddress, nonce)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxsRequiringGasBump provides a mock function with given fields: ctx, address, blockNum, gasBumpThreshold, depth, chainID
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEthTxsRequiringGasBump(ctx context.Context, address ADDR, blockNum int64, gasBumpThreshold int64, depth int64, chainID CHAINID) ([]*TX, error) {
	ret := _m.Called(ctx, address, blockNum, gasBumpThreshold, depth, chainID)

	var r0 []*TX
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, int64, int64, int64, CHAINID) ([]*TX, error)); ok {
		return rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, int64, int64, int64, CHAINID) []*TX); ok {
		r0 = rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*TX)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, int64, int64, int64, CHAINID) error); ok {
		r1 = rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxsRequiringResubmissionDueToInsufficientEth provides a mock function with given fields: address, chainID, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEthTxsRequiringResubmissionDueToInsufficientEth(address ADDR, chainID CHAINID, qopts ...pg.QOpt) ([]*TX, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*TX
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, CHAINID, ...pg.QOpt) ([]*TX, error)); ok {
		return rf(address, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(ADDR, CHAINID, ...pg.QOpt) []*TX); ok {
		r0 = rf(address, chainID, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*TX)
		}
	}

	if rf, ok := ret.Get(1).(func(ADDR, CHAINID, ...pg.QOpt) error); ok {
		r1 = rf(address, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEtxAttemptsConfirmedMissingReceipt provides a mock function with given fields: chainID
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindEtxAttemptsConfirmedMissingReceipt(chainID CHAINID) ([]TXATTEMPT, error) {
	ret := _m.Called(chainID)

	var r0 []TXATTEMPT
	var r1 error
	if rf, ok := ret.Get(0).(func(CHAINID) ([]TXATTEMPT, error)); ok {
		return rf(chainID)
	}
	if rf, ok := ret.Get(0).(func(CHAINID) []TXATTEMPT); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TXATTEMPT)
		}
	}

	if rf, ok := ret.Get(1).(func(CHAINID) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNextUnstartedTransactionFromAddress provides a mock function with given fields: etx, fromAddress, chainID, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindNextUnstartedTransactionFromAddress(etx *TX, fromAddress ADDR, chainID CHAINID, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etx, fromAddress, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*TX, ADDR, CHAINID, ...pg.QOpt) error); ok {
		r0 = rf(etx, fromAddress, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindTransactionsConfirmedInBlockRange provides a mock function with given fields: highBlockNumber, lowBlockNumber, chainID
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) FindTransactionsConfirmedInBlockRange(highBlockNumber int64, lowBlockNumber int64, chainID CHAINID) ([]*TX, error) {
	ret := _m.Called(highBlockNumber, lowBlockNumber, chainID)

	var r0 []*TX
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64, CHAINID) ([]*TX, error)); ok {
		return rf(highBlockNumber, lowBlockNumber, chainID)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, CHAINID) []*TX); ok {
		r0 = rf(highBlockNumber, lowBlockNumber, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*TX)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, int64, CHAINID) error); ok {
		r1 = rf(highBlockNumber, lowBlockNumber, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEthTxInProgress provides a mock function with given fields: fromAddress, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) GetEthTxInProgress(fromAddress ADDR, qopts ...pg.QOpt) (*TX, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fromAddress)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *TX
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, ...pg.QOpt) (*TX, error)); ok {
		return rf(fromAddress, qopts...)
	}
	if rf, ok := ret.Get(0).(func(ADDR, ...pg.QOpt) *TX); ok {
		r0 = rf(fromAddress, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TX)
		}
	}

	if rf, ok := ret.Get(1).(func(ADDR, ...pg.QOpt) error); ok {
		r1 = rf(fromAddress, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInProgressEthTxAttempts provides a mock function with given fields: ctx, address, chainID
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) GetInProgressEthTxAttempts(ctx context.Context, address ADDR, chainID CHAINID) ([]TXATTEMPT, error) {
	ret := _m.Called(ctx, address, chainID)

	var r0 []TXATTEMPT
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAINID) ([]TXATTEMPT, error)); ok {
		return rf(ctx, address, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAINID) []TXATTEMPT); ok {
		r0 = rf(ctx, address, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TXATTEMPT)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, CHAINID) error); ok {
		r1 = rf(ctx, address, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasInProgressTransaction provides a mock function with given fields: account, chainID, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) HasInProgressTransaction(account ADDR, chainID CHAINID, qopts ...pg.QOpt) (bool, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, account, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, CHAINID, ...pg.QOpt) (bool, error)); ok {
		return rf(account, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(ADDR, CHAINID, ...pg.QOpt) bool); ok {
		r0 = rf(account, chainID, qopts...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(ADDR, CHAINID, ...pg.QOpt) error); ok {
		r1 = rf(account, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertEthReceipt provides a mock function with given fields: receipt
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) InsertEthReceipt(receipt *txmgrtypes.Receipt[R, TX_HASH, BLOCK_HASH]) error {
	ret := _m.Called(receipt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgrtypes.Receipt[R, TX_HASH, BLOCK_HASH]) error); ok {
		r0 = rf(receipt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertEthTx provides a mock function with given fields: etx
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) InsertEthTx(etx *TX) error {
	ret := _m.Called(etx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*TX) error); ok {
		r0 = rf(etx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertEthTxAttempt provides a mock function with given fields: attempt
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) InsertEthTxAttempt(attempt *TXATTEMPT) error {
	ret := _m.Called(attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*TXATTEMPT) error); ok {
		r0 = rf(attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadEthTxAttempts provides a mock function with given fields: etx, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) LoadEthTxAttempts(etx *TX, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*TX, ...pg.QOpt) error); ok {
		r0 = rf(etx, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadEthTxesAttempts provides a mock function with given fields: etxs, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) LoadEthTxesAttempts(etxs []*TX, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etxs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*TX, ...pg.QOpt) error); ok {
		r0 = rf(etxs, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkAllConfirmedMissingReceipt provides a mock function with given fields: chainID
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) MarkAllConfirmedMissingReceipt(chainID CHAINID) error {
	ret := _m.Called(chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(CHAINID) error); ok {
		r0 = rf(chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkOldTxesMissingReceiptAsErrored provides a mock function with given fields: blockNum, finalityDepth, chainID, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) MarkOldTxesMissingReceiptAsErrored(blockNum int64, finalityDepth uint32, chainID CHAINID, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, blockNum, finalityDepth, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, uint32, CHAINID, ...pg.QOpt) error); ok {
		r0 = rf(blockNum, finalityDepth, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PreloadEthTxes provides a mock function with given fields: attempts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) PreloadEthTxes(attempts []TXATTEMPT) error {
	ret := _m.Called(attempts)

	var r0 error
	if rf, ok := ret.Get(0).(func([]TXATTEMPT) error); ok {
		r0 = rf(attempts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PruneUnstartedTxQueue provides a mock function with given fields: queueSize, subject, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) PruneUnstartedTxQueue(queueSize uint32, subject uuid.UUID, qopts ...pg.QOpt) (int64, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, queueSize, subject)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint32, uuid.UUID, ...pg.QOpt) (int64, error)); ok {
		return rf(queueSize, subject, qopts...)
	}
	if rf, ok := ret.Get(0).(func(uint32, uuid.UUID, ...pg.QOpt) int64); ok {
		r0 = rf(queueSize, subject, qopts...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint32, uuid.UUID, ...pg.QOpt) error); ok {
		r1 = rf(queueSize, subject, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveConfirmedMissingReceiptAttempt provides a mock function with given fields: ctx, timeout, attempt, broadcastAt
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) SaveConfirmedMissingReceiptAttempt(ctx context.Context, timeout time.Duration, attempt *TXATTEMPT, broadcastAt time.Time) error {
	ret := _m.Called(ctx, timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration, *TXATTEMPT, time.Time) error); ok {
		r0 = rf(ctx, timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveFetchedReceipts provides a mock function with given fields: receipts, chainID
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) SaveFetchedReceipts(receipts []R, chainID CHAINID) error {
	ret := _m.Called(receipts, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func([]R, CHAINID) error); ok {
		r0 = rf(receipts, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveInProgressAttempt provides a mock function with given fields: attempt
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) SaveInProgressAttempt(attempt *TXATTEMPT) error {
	ret := _m.Called(attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*TXATTEMPT) error); ok {
		r0 = rf(attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveInsufficientEthAttempt provides a mock function with given fields: timeout, attempt, broadcastAt
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) SaveInsufficientEthAttempt(timeout time.Duration, attempt *TXATTEMPT, broadcastAt time.Time) error {
	ret := _m.Called(timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, *TXATTEMPT, time.Time) error); ok {
		r0 = rf(timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveReplacementInProgressAttempt provides a mock function with given fields: oldAttempt, replacementAttempt, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) SaveReplacementInProgressAttempt(oldAttempt TXATTEMPT, replacementAttempt *TXATTEMPT, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, oldAttempt, replacementAttempt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(TXATTEMPT, *TXATTEMPT, ...pg.QOpt) error); ok {
		r0 = rf(oldAttempt, replacementAttempt, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveSentAttempt provides a mock function with given fields: timeout, attempt, broadcastAt
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) SaveSentAttempt(timeout time.Duration, attempt *TXATTEMPT, broadcastAt time.Time) error {
	ret := _m.Called(timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, *TXATTEMPT, time.Time) error); ok {
		r0 = rf(timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetBroadcastBeforeBlockNum provides a mock function with given fields: blockNum, chainID
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) SetBroadcastBeforeBlockNum(blockNum int64, chainID CHAINID) error {
	ret := _m.Called(blockNum, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, CHAINID) error); ok {
		r0 = rf(blockNum, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBroadcastAts provides a mock function with given fields: now, etxIDs
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) UpdateBroadcastAts(now time.Time, etxIDs []TXID) error {
	ret := _m.Called(now, etxIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time, []TXID) error); ok {
		r0 = rf(now, etxIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEthKeyNextNonce provides a mock function with given fields: newNextNonce, currentNextNonce, address, chainID, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) UpdateEthKeyNextNonce(newNextNonce TXMETA, currentNextNonce TXMETA, address ADDR, chainID CHAINID, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, newNextNonce, currentNextNonce, address, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(TXMETA, TXMETA, ADDR, CHAINID, ...pg.QOpt) error); ok {
		r0 = rf(newNextNonce, currentNextNonce, address, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEthTxAttemptInProgressToBroadcast provides a mock function with given fields: etx, attempt, NewAttemptState, incrNextNonceCallback, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) UpdateEthTxAttemptInProgressToBroadcast(etx *TX, attempt TXATTEMPT, NewAttemptState txmgrtypes.TxAttemptState, incrNextNonceCallback func(pg.Queryer) error, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etx, attempt, NewAttemptState, incrNextNonceCallback)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*TX, TXATTEMPT, txmgrtypes.TxAttemptState, func(pg.Queryer) error, ...pg.QOpt) error); ok {
		r0 = rf(etx, attempt, NewAttemptState, incrNextNonceCallback, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEthTxFatalError provides a mock function with given fields: etx, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) UpdateEthTxFatalError(etx *TX, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*TX, ...pg.QOpt) error); ok {
		r0 = rf(etx, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEthTxForRebroadcast provides a mock function with given fields: etx, etxAttempt
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) UpdateEthTxForRebroadcast(etx TX, etxAttempt TXATTEMPT) error {
	ret := _m.Called(etx, etxAttempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(TX, TXATTEMPT) error); ok {
		r0 = rf(etx, etxAttempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEthTxUnstartedToInProgress provides a mock function with given fields: etx, attempt, qopts
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) UpdateEthTxUnstartedToInProgress(etx *TX, attempt *TXATTEMPT, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etx, attempt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*TX, *TXATTEMPT, ...pg.QOpt) error); ok {
		r0 = rf(etx, attempt, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEthTxsUnconfirmed provides a mock function with given fields: ids
func (_m *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]) UpdateEthTxsUnconfirmed(ids []int64) error {
	ret := _m.Called(ids)

	var r0 error
	if rf, ok := ret.Get(0).(func([]int64) error); ok {
		r0 = rf(ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTxStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewTxStore creates a new instance of TxStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxStore[ADDR interface{}, CHAINID interface{}, TX_HASH types.Hashable[TX_HASH], BLOCK_HASH types.Hashable[BLOCK_HASH], NEWTX interface{}, R interface{}, TX interface{}, TXATTEMPT interface{}, TXID interface{}, TXMETA interface{}](t mockConstructorTestingTNewTxStore) *TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA] {
	mock := &TxStore[ADDR, CHAINID, TX_HASH, BLOCK_HASH, NEWTX, R, TX, TXATTEMPT, TXID, TXMETA]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}