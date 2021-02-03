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

// LoyaltyProgramBalance The balance in a Loyalty Program for some Customer.
type LoyaltyProgramBalance struct {
	// Sum of current active points amounts
	CurrentBalance float32 `json:"currentBalance"`
	// Sum of pending points amounts
	PendingBalance float32 `json:"pendingBalance"`
	// Sum of expired points amounts
	ExpiredBalance float32 `json:"expiredBalance"`
	// Sum of spent points amounts
	SpentBalance float32 `json:"spentBalance"`
	// Sum of current active points amounts, including additions and deductions on open sessions
	TentativeCurrentBalance float32 `json:"tentativeCurrentBalance"`
}

// GetCurrentBalance returns the CurrentBalance field value
func (o *LoyaltyProgramBalance) GetCurrentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentBalance
}

// SetCurrentBalance sets field value
func (o *LoyaltyProgramBalance) SetCurrentBalance(v float32) {
	o.CurrentBalance = v
}

// GetPendingBalance returns the PendingBalance field value
func (o *LoyaltyProgramBalance) GetPendingBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PendingBalance
}

// SetPendingBalance sets field value
func (o *LoyaltyProgramBalance) SetPendingBalance(v float32) {
	o.PendingBalance = v
}

// GetExpiredBalance returns the ExpiredBalance field value
func (o *LoyaltyProgramBalance) GetExpiredBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExpiredBalance
}

// SetExpiredBalance sets field value
func (o *LoyaltyProgramBalance) SetExpiredBalance(v float32) {
	o.ExpiredBalance = v
}

// GetSpentBalance returns the SpentBalance field value
func (o *LoyaltyProgramBalance) GetSpentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SpentBalance
}

// SetSpentBalance sets field value
func (o *LoyaltyProgramBalance) SetSpentBalance(v float32) {
	o.SpentBalance = v
}

// GetTentativeCurrentBalance returns the TentativeCurrentBalance field value
func (o *LoyaltyProgramBalance) GetTentativeCurrentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TentativeCurrentBalance
}

// SetTentativeCurrentBalance sets field value
func (o *LoyaltyProgramBalance) SetTentativeCurrentBalance(v float32) {
	o.TentativeCurrentBalance = v
}

type NullableLoyaltyProgramBalance struct {
	Value        LoyaltyProgramBalance
	ExplicitNull bool
}

func (v NullableLoyaltyProgramBalance) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyProgramBalance) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
