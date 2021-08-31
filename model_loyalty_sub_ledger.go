/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// LoyaltySubLedger Ledger of Balance in Loyalty Program for a Customer.
type LoyaltySubLedger struct {
	// ⚠️ Deprecated: Use 'totalActivePoints' property instead. Total amount of currently active and available points in the customer's balance.
	Total float32 `json:"total"`
	// Total amount of currently active and available points in the customer's balance.
	TotalActivePoints float32 `json:"totalActivePoints"`
	// Total amount of pending points, which are not active yet but will become active in the future.
	TotalPendingPoints float32 `json:"totalPendingPoints"`
	// Total amount of points already spent by this customer.
	TotalSpentPoints float32 `json:"totalSpentPoints"`
	// Total amount of points, that expired without ever being spent.
	TotalExpiredPoints float32 `json:"totalExpiredPoints"`
	// List of all events that have happened such as additions, subtractions and expiries.
	Transactions *[]LoyaltyLedgerEntry `json:"transactions,omitempty"`
	// List of all points that will expire.
	ExpiringPoints *[]LoyaltyLedgerEntry `json:"expiringPoints,omitempty"`
	// List of all currently active points.
	ActivePoints *[]LoyaltyLedgerEntry `json:"activePoints,omitempty"`
	// List of all points pending activation.
	PendingPoints *[]LoyaltyLedgerEntry `json:"pendingPoints,omitempty"`
	// List of expired points.
	ExpiredPoints *[]LoyaltyLedgerEntry `json:"expiredPoints,omitempty"`
	CurrentTier   *Tier                 `json:"currentTier,omitempty"`
}

// GetTotal returns the Total field value
func (o *LoyaltySubLedger) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// SetTotal sets field value
func (o *LoyaltySubLedger) SetTotal(v float32) {
	o.Total = v
}

// GetTotalActivePoints returns the TotalActivePoints field value
func (o *LoyaltySubLedger) GetTotalActivePoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalActivePoints
}

// SetTotalActivePoints sets field value
func (o *LoyaltySubLedger) SetTotalActivePoints(v float32) {
	o.TotalActivePoints = v
}

// GetTotalPendingPoints returns the TotalPendingPoints field value
func (o *LoyaltySubLedger) GetTotalPendingPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPendingPoints
}

// SetTotalPendingPoints sets field value
func (o *LoyaltySubLedger) SetTotalPendingPoints(v float32) {
	o.TotalPendingPoints = v
}

// GetTotalSpentPoints returns the TotalSpentPoints field value
func (o *LoyaltySubLedger) GetTotalSpentPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalSpentPoints
}

// SetTotalSpentPoints sets field value
func (o *LoyaltySubLedger) SetTotalSpentPoints(v float32) {
	o.TotalSpentPoints = v
}

// GetTotalExpiredPoints returns the TotalExpiredPoints field value
func (o *LoyaltySubLedger) GetTotalExpiredPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalExpiredPoints
}

// SetTotalExpiredPoints sets field value
func (o *LoyaltySubLedger) SetTotalExpiredPoints(v float32) {
	o.TotalExpiredPoints = v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *LoyaltySubLedger) GetTransactions() []LoyaltyLedgerEntry {
	if o == nil || o.Transactions == nil {
		var ret []LoyaltyLedgerEntry
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltySubLedger) GetTransactionsOk() ([]LoyaltyLedgerEntry, bool) {
	if o == nil || o.Transactions == nil {
		var ret []LoyaltyLedgerEntry
		return ret, false
	}
	return *o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *LoyaltySubLedger) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []LoyaltyLedgerEntry and assigns it to the Transactions field.
func (o *LoyaltySubLedger) SetTransactions(v []LoyaltyLedgerEntry) {
	o.Transactions = &v
}

// GetExpiringPoints returns the ExpiringPoints field value if set, zero value otherwise.
func (o *LoyaltySubLedger) GetExpiringPoints() []LoyaltyLedgerEntry {
	if o == nil || o.ExpiringPoints == nil {
		var ret []LoyaltyLedgerEntry
		return ret
	}
	return *o.ExpiringPoints
}

// GetExpiringPointsOk returns a tuple with the ExpiringPoints field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltySubLedger) GetExpiringPointsOk() ([]LoyaltyLedgerEntry, bool) {
	if o == nil || o.ExpiringPoints == nil {
		var ret []LoyaltyLedgerEntry
		return ret, false
	}
	return *o.ExpiringPoints, true
}

// HasExpiringPoints returns a boolean if a field has been set.
func (o *LoyaltySubLedger) HasExpiringPoints() bool {
	if o != nil && o.ExpiringPoints != nil {
		return true
	}

	return false
}

// SetExpiringPoints gets a reference to the given []LoyaltyLedgerEntry and assigns it to the ExpiringPoints field.
func (o *LoyaltySubLedger) SetExpiringPoints(v []LoyaltyLedgerEntry) {
	o.ExpiringPoints = &v
}

// GetActivePoints returns the ActivePoints field value if set, zero value otherwise.
func (o *LoyaltySubLedger) GetActivePoints() []LoyaltyLedgerEntry {
	if o == nil || o.ActivePoints == nil {
		var ret []LoyaltyLedgerEntry
		return ret
	}
	return *o.ActivePoints
}

// GetActivePointsOk returns a tuple with the ActivePoints field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltySubLedger) GetActivePointsOk() ([]LoyaltyLedgerEntry, bool) {
	if o == nil || o.ActivePoints == nil {
		var ret []LoyaltyLedgerEntry
		return ret, false
	}
	return *o.ActivePoints, true
}

// HasActivePoints returns a boolean if a field has been set.
func (o *LoyaltySubLedger) HasActivePoints() bool {
	if o != nil && o.ActivePoints != nil {
		return true
	}

	return false
}

// SetActivePoints gets a reference to the given []LoyaltyLedgerEntry and assigns it to the ActivePoints field.
func (o *LoyaltySubLedger) SetActivePoints(v []LoyaltyLedgerEntry) {
	o.ActivePoints = &v
}

// GetPendingPoints returns the PendingPoints field value if set, zero value otherwise.
func (o *LoyaltySubLedger) GetPendingPoints() []LoyaltyLedgerEntry {
	if o == nil || o.PendingPoints == nil {
		var ret []LoyaltyLedgerEntry
		return ret
	}
	return *o.PendingPoints
}

// GetPendingPointsOk returns a tuple with the PendingPoints field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltySubLedger) GetPendingPointsOk() ([]LoyaltyLedgerEntry, bool) {
	if o == nil || o.PendingPoints == nil {
		var ret []LoyaltyLedgerEntry
		return ret, false
	}
	return *o.PendingPoints, true
}

// HasPendingPoints returns a boolean if a field has been set.
func (o *LoyaltySubLedger) HasPendingPoints() bool {
	if o != nil && o.PendingPoints != nil {
		return true
	}

	return false
}

// SetPendingPoints gets a reference to the given []LoyaltyLedgerEntry and assigns it to the PendingPoints field.
func (o *LoyaltySubLedger) SetPendingPoints(v []LoyaltyLedgerEntry) {
	o.PendingPoints = &v
}

// GetExpiredPoints returns the ExpiredPoints field value if set, zero value otherwise.
func (o *LoyaltySubLedger) GetExpiredPoints() []LoyaltyLedgerEntry {
	if o == nil || o.ExpiredPoints == nil {
		var ret []LoyaltyLedgerEntry
		return ret
	}
	return *o.ExpiredPoints
}

// GetExpiredPointsOk returns a tuple with the ExpiredPoints field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltySubLedger) GetExpiredPointsOk() ([]LoyaltyLedgerEntry, bool) {
	if o == nil || o.ExpiredPoints == nil {
		var ret []LoyaltyLedgerEntry
		return ret, false
	}
	return *o.ExpiredPoints, true
}

// HasExpiredPoints returns a boolean if a field has been set.
func (o *LoyaltySubLedger) HasExpiredPoints() bool {
	if o != nil && o.ExpiredPoints != nil {
		return true
	}

	return false
}

// SetExpiredPoints gets a reference to the given []LoyaltyLedgerEntry and assigns it to the ExpiredPoints field.
func (o *LoyaltySubLedger) SetExpiredPoints(v []LoyaltyLedgerEntry) {
	o.ExpiredPoints = &v
}

// GetCurrentTier returns the CurrentTier field value if set, zero value otherwise.
func (o *LoyaltySubLedger) GetCurrentTier() Tier {
	if o == nil || o.CurrentTier == nil {
		var ret Tier
		return ret
	}
	return *o.CurrentTier
}

// GetCurrentTierOk returns a tuple with the CurrentTier field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltySubLedger) GetCurrentTierOk() (Tier, bool) {
	if o == nil || o.CurrentTier == nil {
		var ret Tier
		return ret, false
	}
	return *o.CurrentTier, true
}

// HasCurrentTier returns a boolean if a field has been set.
func (o *LoyaltySubLedger) HasCurrentTier() bool {
	if o != nil && o.CurrentTier != nil {
		return true
	}

	return false
}

// SetCurrentTier gets a reference to the given Tier and assigns it to the CurrentTier field.
func (o *LoyaltySubLedger) SetCurrentTier(v Tier) {
	o.CurrentTier = &v
}

type NullableLoyaltySubLedger struct {
	Value        LoyaltySubLedger
	ExplicitNull bool
}

func (v NullableLoyaltySubLedger) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltySubLedger) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
