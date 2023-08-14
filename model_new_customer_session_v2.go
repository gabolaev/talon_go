/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}`
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// NewCustomerSessionV2 The representation of the customer session.
type NewCustomerSessionV2 struct {
	// ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known `profileId`, we recommend you use a guest `profileId`.
	ProfileId *string `json:"profileId,omitempty"`
	// When using the `dry` query parameter, use this property to list the campaign to be evaluated by the Rule Engine.  These campaigns will be evaluated, even if they are disabled, allowing you to test specific campaigns before activating them.
	EvaluableCampaignIds *[]int32 `json:"evaluableCampaignIds,omitempty"`
	// Any coupon codes entered.  **Important**: If you [create a coupon budget](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets/#budget-types) for your campaign, ensure the session contains a coupon code by the time you close it.
	CouponCodes *[]string `json:"couponCodes,omitempty"`
	// Any referral code entered.  **Important**: If you [create a referral budget](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets/#budget-types) for your campaign, ensure the session contains a referral code by the time you close it.
	ReferralCode *string `json:"referralCode,omitempty"`
	// Any loyalty cards used.
	LoyaltyCards *[]string `json:"loyaltyCards,omitempty"`
	// Indicates the current state of the session. Sessions can be created as `open` or `closed`. The state transitions are:  1. `open` → `closed` 2. `open` → `cancelled` 3. Either:    - `closed` → `cancelled` (**only** via [Update customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2)) or    - `closed` → `partially_returned` (**only** via [Return cart items](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/returnCartItems))    - `closed` → `open` (**only** via [Reopen customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/reopenCustomerSession)) 4. `partially_returned` → `cancelled`  For more information, see [Customer session states](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions).
	State *string `json:"state,omitempty"`
	// The items to add to this session. **Do not exceed 1000 items** and ensure the sum of all cart item's `quantity` **does not exceed 10.000** per request.
	CartItems *[]CartItem `json:"cartItems,omitempty"`
	// Use this property to set a value for the additional costs of this session, such as a shipping cost.  They must be created in the Campaign Manager before you set them with this property. See [Managing additional costs](https://docs.talon.one/docs/product/account/dev-tools/managing-additional-costs).
	AdditionalCosts *map[string]AdditionalCost `json:"additionalCosts,omitempty"`
	// Session custom identifiers that you can set limits on or use inside your rules.  For example, you can use IP addresses as identifiers to potentially identify devices and limit discounts abuse in case of customers creating multiple accounts. See the [tutorial](https://docs.talon.one/docs/dev/tutorials/using-identifiers).  **Important**: Ensure the session contains an identifier by the time you close it if: - You [create a unique identifier budget](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets/#budget-types) for your campaign. - Your campaign has [coupons](https://docs.talon.one/docs/product/campaigns/coupons/coupon-page-overview).
	Identifiers *[]string `json:"identifiers,omitempty"`
	// Use this property to set a value for the attributes of your choice. Attributes represent any information to attach to your session, like the shipping city.  You can use [built-in attributes](https://docs.talon.one/docs/dev/concepts/attributes#built-in-attributes) or [custom ones](https://docs.talon.one/docs/dev/concepts/attributes#custom-attributes). Custom attributes must be created in the Campaign Manager before you set them with this property.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetProfileIdOk() (string, bool) {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret, false
	}
	return *o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *NewCustomerSessionV2) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetEvaluableCampaignIds returns the EvaluableCampaignIds field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetEvaluableCampaignIds() []int32 {
	if o == nil || o.EvaluableCampaignIds == nil {
		var ret []int32
		return ret
	}
	return *o.EvaluableCampaignIds
}

// GetEvaluableCampaignIdsOk returns a tuple with the EvaluableCampaignIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetEvaluableCampaignIdsOk() ([]int32, bool) {
	if o == nil || o.EvaluableCampaignIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.EvaluableCampaignIds, true
}

// HasEvaluableCampaignIds returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasEvaluableCampaignIds() bool {
	if o != nil && o.EvaluableCampaignIds != nil {
		return true
	}

	return false
}

// SetEvaluableCampaignIds gets a reference to the given []int32 and assigns it to the EvaluableCampaignIds field.
func (o *NewCustomerSessionV2) SetEvaluableCampaignIds(v []int32) {
	o.EvaluableCampaignIds = &v
}

// GetCouponCodes returns the CouponCodes field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetCouponCodes() []string {
	if o == nil || o.CouponCodes == nil {
		var ret []string
		return ret
	}
	return *o.CouponCodes
}

// GetCouponCodesOk returns a tuple with the CouponCodes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetCouponCodesOk() ([]string, bool) {
	if o == nil || o.CouponCodes == nil {
		var ret []string
		return ret, false
	}
	return *o.CouponCodes, true
}

// HasCouponCodes returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasCouponCodes() bool {
	if o != nil && o.CouponCodes != nil {
		return true
	}

	return false
}

// SetCouponCodes gets a reference to the given []string and assigns it to the CouponCodes field.
func (o *NewCustomerSessionV2) SetCouponCodes(v []string) {
	o.CouponCodes = &v
}

// GetReferralCode returns the ReferralCode field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetReferralCode() string {
	if o == nil || o.ReferralCode == nil {
		var ret string
		return ret
	}
	return *o.ReferralCode
}

// GetReferralCodeOk returns a tuple with the ReferralCode field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetReferralCodeOk() (string, bool) {
	if o == nil || o.ReferralCode == nil {
		var ret string
		return ret, false
	}
	return *o.ReferralCode, true
}

// HasReferralCode returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasReferralCode() bool {
	if o != nil && o.ReferralCode != nil {
		return true
	}

	return false
}

// SetReferralCode gets a reference to the given string and assigns it to the ReferralCode field.
func (o *NewCustomerSessionV2) SetReferralCode(v string) {
	o.ReferralCode = &v
}

// GetLoyaltyCards returns the LoyaltyCards field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetLoyaltyCards() []string {
	if o == nil || o.LoyaltyCards == nil {
		var ret []string
		return ret
	}
	return *o.LoyaltyCards
}

// GetLoyaltyCardsOk returns a tuple with the LoyaltyCards field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetLoyaltyCardsOk() ([]string, bool) {
	if o == nil || o.LoyaltyCards == nil {
		var ret []string
		return ret, false
	}
	return *o.LoyaltyCards, true
}

// HasLoyaltyCards returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasLoyaltyCards() bool {
	if o != nil && o.LoyaltyCards != nil {
		return true
	}

	return false
}

// SetLoyaltyCards gets a reference to the given []string and assigns it to the LoyaltyCards field.
func (o *NewCustomerSessionV2) SetLoyaltyCards(v []string) {
	o.LoyaltyCards = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NewCustomerSessionV2) SetState(v string) {
	o.State = &v
}

// GetCartItems returns the CartItems field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetCartItems() []CartItem {
	if o == nil || o.CartItems == nil {
		var ret []CartItem
		return ret
	}
	return *o.CartItems
}

// GetCartItemsOk returns a tuple with the CartItems field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetCartItemsOk() ([]CartItem, bool) {
	if o == nil || o.CartItems == nil {
		var ret []CartItem
		return ret, false
	}
	return *o.CartItems, true
}

// HasCartItems returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasCartItems() bool {
	if o != nil && o.CartItems != nil {
		return true
	}

	return false
}

// SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.
func (o *NewCustomerSessionV2) SetCartItems(v []CartItem) {
	o.CartItems = &v
}

// GetAdditionalCosts returns the AdditionalCosts field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetAdditionalCosts() map[string]AdditionalCost {
	if o == nil || o.AdditionalCosts == nil {
		var ret map[string]AdditionalCost
		return ret
	}
	return *o.AdditionalCosts
}

// GetAdditionalCostsOk returns a tuple with the AdditionalCosts field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetAdditionalCostsOk() (map[string]AdditionalCost, bool) {
	if o == nil || o.AdditionalCosts == nil {
		var ret map[string]AdditionalCost
		return ret, false
	}
	return *o.AdditionalCosts, true
}

// HasAdditionalCosts returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasAdditionalCosts() bool {
	if o != nil && o.AdditionalCosts != nil {
		return true
	}

	return false
}

// SetAdditionalCosts gets a reference to the given map[string]AdditionalCost and assigns it to the AdditionalCosts field.
func (o *NewCustomerSessionV2) SetAdditionalCosts(v map[string]AdditionalCost) {
	o.AdditionalCosts = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetIdentifiers() []string {
	if o == nil || o.Identifiers == nil {
		var ret []string
		return ret
	}
	return *o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetIdentifiersOk() ([]string, bool) {
	if o == nil || o.Identifiers == nil {
		var ret []string
		return ret, false
	}
	return *o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []string and assigns it to the Identifiers field.
func (o *NewCustomerSessionV2) SetIdentifiers(v []string) {
	o.Identifiers = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NewCustomerSessionV2) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSessionV2) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NewCustomerSessionV2) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *NewCustomerSessionV2) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

type NullableNewCustomerSessionV2 struct {
	Value        NewCustomerSessionV2
	ExplicitNull bool
}

func (v NullableNewCustomerSessionV2) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewCustomerSessionV2) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
