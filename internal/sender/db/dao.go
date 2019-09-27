package db

import (
	"math/big"

	"github.com/void616/gm-mint-sender/internal/sender/db/types"
	sumuslib "github.com/void616/gm-sumuslib"
)

// DAO is a DB interface
type DAO interface {
	Available() bool
	DuplicateError(err error) bool
	MaxPacketError(err error) bool
	Migrate() error

	// PutWallet saves sending wallet to track it's outgoing transactions later
	PutWallet(v *types.Wallet) error
	// ListWallets get s list of all known senders
	ListWallets() ([]*types.Wallet, error)

	// PutSending adds sending request
	PutSending(v *types.Sending) error
	// ListEnqueuedSendings gets a list of enqueued sending requests
	ListEnqueuedSendings(max uint16) ([]*types.Sending, error)
	// ListStaleSendings gets a list of stale posted requests
	ListStaleSendings(elderThanBlockID *big.Int, max uint16) ([]*types.Sending, error)
	// ListUnnotifiedSendings gets a list of requests without notification of requestor
	ListUnnotifiedSendings(max uint16) ([]*types.Sending, error)
	// UpdateSending updates sending
	UpdateSending(v *types.Sending) error
	// SetSendingConfirmed updates sending
	SetSendingConfirmed(d sumuslib.Digest, from sumuslib.PublicKey, block *big.Int) error

	// EarliestBlock finds a minimal block ID at which a transaction has been sent
	EarliestBlock() (*big.Int, bool, error)
	// LatestSenderNonce gets max used nonce for specified sender or zero
	LatestSenderNonce(sender sumuslib.PublicKey) (uint64, error)
}
