package anchors

import (
	"context"
	"time"

	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("anchorRepository")

// AnchorRepository defines a set of functions that enable
// implementations of any type that stores and retrieves the anchoring, and pre-anchoring details.
type AnchorRepository interface {

	// PreCommitAnchor calls a transaction's PreCommit on the smart contract, to pre-commit a document update.
	PreCommitAnchor(ctx context.Context, anchorID AnchorID, signingRoot DocumentRoot) (confirmations chan bool, err error)

	// CommitAnchor sends a Commit transaction to Ethereum.
	CommitAnchor(ctx context.Context, anchorID AnchorID, documentRoot DocumentRoot, proof [32]byte) (chan bool, error)

	// GetAnchorData takes an anchorID and returns the corresponding documentRoot from the chain.
	GetAnchorData(anchorID AnchorID) (docRoot DocumentRoot, anchoredTime time.Time, err error)

	// HasValidPreCommit checks if the given anchorID has a valid pre-commit.
	HasValidPreCommit(anchorID AnchorID) bool
}
