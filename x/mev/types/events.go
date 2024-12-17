package types

// Event types for MEV module
const (
	EventTypeBundleSubmitted = "bundle_submitted"
	EventTypeBundleProcessed = "bundle_processed"
	EventTypeBundleExecuted  = "bundle_executed"
	EventTypeBundleFailed    = "bundle_failed"

	// Attribute keys
	AttributeKeyBundleID    = "bundle_id"
	AttributeKeySender      = "sender"
	AttributeKeyBlockHeight = "block_height"
	AttributeKeyTxCount     = "tx_count"
	AttributeKeyBundleFee   = "bundle_fee"
	AttributeKeyPriority    = "priority"
	AttributeKeyError       = "error"
)
