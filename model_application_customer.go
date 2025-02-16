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
	"time"
)

// ApplicationCustomer
type ApplicationCustomer struct {
	// Internal ID of this entity. Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created. The time this entity was created. The time this entity was created. The time this entity was created.
	Created time.Time `json:"created"`
	// The integration ID set by your integration layer. The integration ID set by your integration layer.
	IntegrationId string `json:"integrationId"`
	// Arbitrary properties associated with this item.
	Attributes map[string]interface{} `json:"attributes"`
	// The ID of the Talon.One account that owns this profile. The ID of the Talon.One account that owns this profile.
	AccountId int32 `json:"accountId"`
	// The total amount of closed sessions by a customer. A closed session is a successful purchase.
	ClosedSessions int32 `json:"closedSessions"`
	// The total amount of money spent by the customer **before** discounts are applied.  The total sales amount excludes the following: - Cancelled or reopened sessions. - Returned items.
	TotalSales float32 `json:"totalSales"`
	// **DEPRECATED** A list of loyalty programs joined by the customer.
	LoyaltyMemberships *[]LoyaltyMembership `json:"loyaltyMemberships,omitempty"`
	// The audiences the customer belongs to.
	AudienceMemberships *[]AudienceMembership `json:"audienceMemberships,omitempty"`
	// Timestamp of the most recent event received from this customer. This field is updated on calls that trigger the rule-engine and that are not [dry requests](https://docs.talon.one/docs/dev/integration-api/dry-requests/#overlay).  For example, [reserving a coupon](https://docs.talon.one/integration-api#operation/createCouponReservation) for a customer doesn't impact this field.
	LastActivity time.Time `json:"lastActivity"`
	// Shows whether the customer is part of a sandbox or live Application. See the [docs](https://docs.talon.one/docs/product/applications/overview#application-environments).
	Sandbox *bool `json:"sandbox,omitempty"`
	// The Integration ID of the Customer Profile that referred this Customer in the Application.
	AdvocateIntegrationId *string `json:"advocateIntegrationId,omitempty"`
}

// GetId returns the Id field value
func (o *ApplicationCustomer) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *ApplicationCustomer) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *ApplicationCustomer) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *ApplicationCustomer) SetCreated(v time.Time) {
	o.Created = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *ApplicationCustomer) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// SetIntegrationId sets field value
func (o *ApplicationCustomer) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetAttributes returns the Attributes field value
func (o *ApplicationCustomer) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// SetAttributes sets field value
func (o *ApplicationCustomer) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetAccountId returns the AccountId field value
func (o *ApplicationCustomer) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *ApplicationCustomer) SetAccountId(v int32) {
	o.AccountId = v
}

// GetClosedSessions returns the ClosedSessions field value
func (o *ApplicationCustomer) GetClosedSessions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClosedSessions
}

// SetClosedSessions sets field value
func (o *ApplicationCustomer) SetClosedSessions(v int32) {
	o.ClosedSessions = v
}

// GetTotalSales returns the TotalSales field value
func (o *ApplicationCustomer) GetTotalSales() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalSales
}

// SetTotalSales sets field value
func (o *ApplicationCustomer) SetTotalSales(v float32) {
	o.TotalSales = v
}

// GetLoyaltyMemberships returns the LoyaltyMemberships field value if set, zero value otherwise.
func (o *ApplicationCustomer) GetLoyaltyMemberships() []LoyaltyMembership {
	if o == nil || o.LoyaltyMemberships == nil {
		var ret []LoyaltyMembership
		return ret
	}
	return *o.LoyaltyMemberships
}

// GetLoyaltyMembershipsOk returns a tuple with the LoyaltyMemberships field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCustomer) GetLoyaltyMembershipsOk() ([]LoyaltyMembership, bool) {
	if o == nil || o.LoyaltyMemberships == nil {
		var ret []LoyaltyMembership
		return ret, false
	}
	return *o.LoyaltyMemberships, true
}

// HasLoyaltyMemberships returns a boolean if a field has been set.
func (o *ApplicationCustomer) HasLoyaltyMemberships() bool {
	if o != nil && o.LoyaltyMemberships != nil {
		return true
	}

	return false
}

// SetLoyaltyMemberships gets a reference to the given []LoyaltyMembership and assigns it to the LoyaltyMemberships field.
func (o *ApplicationCustomer) SetLoyaltyMemberships(v []LoyaltyMembership) {
	o.LoyaltyMemberships = &v
}

// GetAudienceMemberships returns the AudienceMemberships field value if set, zero value otherwise.
func (o *ApplicationCustomer) GetAudienceMemberships() []AudienceMembership {
	if o == nil || o.AudienceMemberships == nil {
		var ret []AudienceMembership
		return ret
	}
	return *o.AudienceMemberships
}

// GetAudienceMembershipsOk returns a tuple with the AudienceMemberships field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCustomer) GetAudienceMembershipsOk() ([]AudienceMembership, bool) {
	if o == nil || o.AudienceMemberships == nil {
		var ret []AudienceMembership
		return ret, false
	}
	return *o.AudienceMemberships, true
}

// HasAudienceMemberships returns a boolean if a field has been set.
func (o *ApplicationCustomer) HasAudienceMemberships() bool {
	if o != nil && o.AudienceMemberships != nil {
		return true
	}

	return false
}

// SetAudienceMemberships gets a reference to the given []AudienceMembership and assigns it to the AudienceMemberships field.
func (o *ApplicationCustomer) SetAudienceMemberships(v []AudienceMembership) {
	o.AudienceMemberships = &v
}

// GetLastActivity returns the LastActivity field value
func (o *ApplicationCustomer) GetLastActivity() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastActivity
}

// SetLastActivity sets field value
func (o *ApplicationCustomer) SetLastActivity(v time.Time) {
	o.LastActivity = v
}

// GetSandbox returns the Sandbox field value if set, zero value otherwise.
func (o *ApplicationCustomer) GetSandbox() bool {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret
	}
	return *o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCustomer) GetSandboxOk() (bool, bool) {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret, false
	}
	return *o.Sandbox, true
}

// HasSandbox returns a boolean if a field has been set.
func (o *ApplicationCustomer) HasSandbox() bool {
	if o != nil && o.Sandbox != nil {
		return true
	}

	return false
}

// SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.
func (o *ApplicationCustomer) SetSandbox(v bool) {
	o.Sandbox = &v
}

// GetAdvocateIntegrationId returns the AdvocateIntegrationId field value if set, zero value otherwise.
func (o *ApplicationCustomer) GetAdvocateIntegrationId() string {
	if o == nil || o.AdvocateIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.AdvocateIntegrationId
}

// GetAdvocateIntegrationIdOk returns a tuple with the AdvocateIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCustomer) GetAdvocateIntegrationIdOk() (string, bool) {
	if o == nil || o.AdvocateIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.AdvocateIntegrationId, true
}

// HasAdvocateIntegrationId returns a boolean if a field has been set.
func (o *ApplicationCustomer) HasAdvocateIntegrationId() bool {
	if o != nil && o.AdvocateIntegrationId != nil {
		return true
	}

	return false
}

// SetAdvocateIntegrationId gets a reference to the given string and assigns it to the AdvocateIntegrationId field.
func (o *ApplicationCustomer) SetAdvocateIntegrationId(v string) {
	o.AdvocateIntegrationId = &v
}

type NullableApplicationCustomer struct {
	Value        ApplicationCustomer
	ExplicitNull bool
}

func (v NullableApplicationCustomer) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApplicationCustomer) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
