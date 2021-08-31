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

// NewApplication
type NewApplication struct {
	// The name of this application.
	Name string `json:"name"`
	// A longer description of the application.
	Description *string `json:"description,omitempty"`
	// A string containing an IANA timezone descriptor.
	Timezone string `json:"timezone"`
	// A string describing a default currency for new customer sessions.
	Currency string `json:"currency"`
	// A string indicating how should campaigns in this application deal with case sensitivity on coupon codes.
	CaseSensitivity *string `json:"caseSensitivity,omitempty"`
	// Arbitrary properties associated with this campaign
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// Default limits for campaigns created in this application
	Limits *[]LimitConfig `json:"limits,omitempty"`
	// Default priority for campaigns created in this application, can be one of (universal, stackable, exclusive). If no value is provided, this is set to \"universal\"
	CampaignPriority *string `json:"campaignPriority,omitempty"`
	// The strategy used when choosing exclusive campaigns for evaluation, can be one of (listOrder, lowestDiscount, highestDiscount). If no value is provided, this is set to \"listOrder\"
	ExclusiveCampaignsStrategy *string `json:"exclusiveCampaignsStrategy,omitempty"`
	// The default scope to apply \"setDiscount\" effects on if no scope was provided with the effect.
	DefaultDiscountScope *string `json:"defaultDiscountScope,omitempty"`
	// Indicates if discounts should cascade for this application
	EnableCascadingDiscounts *bool `json:"enableCascadingDiscounts,omitempty"`
	// Indicates if cart items of quantity larger than one should be separated into different items of quantity one
	EnableFlattenedCartItems *bool               `json:"enableFlattenedCartItems,omitempty"`
	AttributesSettings       *AttributesSettings `json:"attributesSettings,omitempty"`
	// Indicates if this is a live or sandbox application
	Sandbox *bool `json:"sandbox,omitempty"`
	// Hex key for HMAC-signing API calls as coming from this application (16 hex digits)
	Key *string `json:"key,omitempty"`
}

// GetName returns the Name field value
func (o *NewApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *NewApplication) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewApplication) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewApplication) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewApplication) SetDescription(v string) {
	o.Description = &v
}

// GetTimezone returns the Timezone field value
func (o *NewApplication) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// SetTimezone sets field value
func (o *NewApplication) SetTimezone(v string) {
	o.Timezone = v
}

// GetCurrency returns the Currency field value
func (o *NewApplication) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// SetCurrency sets field value
func (o *NewApplication) SetCurrency(v string) {
	o.Currency = v
}

// GetCaseSensitivity returns the CaseSensitivity field value if set, zero value otherwise.
func (o *NewApplication) GetCaseSensitivity() string {
	if o == nil || o.CaseSensitivity == nil {
		var ret string
		return ret
	}
	return *o.CaseSensitivity
}

// GetCaseSensitivityOk returns a tuple with the CaseSensitivity field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetCaseSensitivityOk() (string, bool) {
	if o == nil || o.CaseSensitivity == nil {
		var ret string
		return ret, false
	}
	return *o.CaseSensitivity, true
}

// HasCaseSensitivity returns a boolean if a field has been set.
func (o *NewApplication) HasCaseSensitivity() bool {
	if o != nil && o.CaseSensitivity != nil {
		return true
	}

	return false
}

// SetCaseSensitivity gets a reference to the given string and assigns it to the CaseSensitivity field.
func (o *NewApplication) SetCaseSensitivity(v string) {
	o.CaseSensitivity = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NewApplication) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NewApplication) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *NewApplication) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *NewApplication) GetLimits() []LimitConfig {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetLimitsOk() ([]LimitConfig, bool) {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret, false
	}
	return *o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *NewApplication) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.
func (o *NewApplication) SetLimits(v []LimitConfig) {
	o.Limits = &v
}

// GetCampaignPriority returns the CampaignPriority field value if set, zero value otherwise.
func (o *NewApplication) GetCampaignPriority() string {
	if o == nil || o.CampaignPriority == nil {
		var ret string
		return ret
	}
	return *o.CampaignPriority
}

// GetCampaignPriorityOk returns a tuple with the CampaignPriority field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetCampaignPriorityOk() (string, bool) {
	if o == nil || o.CampaignPriority == nil {
		var ret string
		return ret, false
	}
	return *o.CampaignPriority, true
}

// HasCampaignPriority returns a boolean if a field has been set.
func (o *NewApplication) HasCampaignPriority() bool {
	if o != nil && o.CampaignPriority != nil {
		return true
	}

	return false
}

// SetCampaignPriority gets a reference to the given string and assigns it to the CampaignPriority field.
func (o *NewApplication) SetCampaignPriority(v string) {
	o.CampaignPriority = &v
}

// GetExclusiveCampaignsStrategy returns the ExclusiveCampaignsStrategy field value if set, zero value otherwise.
func (o *NewApplication) GetExclusiveCampaignsStrategy() string {
	if o == nil || o.ExclusiveCampaignsStrategy == nil {
		var ret string
		return ret
	}
	return *o.ExclusiveCampaignsStrategy
}

// GetExclusiveCampaignsStrategyOk returns a tuple with the ExclusiveCampaignsStrategy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetExclusiveCampaignsStrategyOk() (string, bool) {
	if o == nil || o.ExclusiveCampaignsStrategy == nil {
		var ret string
		return ret, false
	}
	return *o.ExclusiveCampaignsStrategy, true
}

// HasExclusiveCampaignsStrategy returns a boolean if a field has been set.
func (o *NewApplication) HasExclusiveCampaignsStrategy() bool {
	if o != nil && o.ExclusiveCampaignsStrategy != nil {
		return true
	}

	return false
}

// SetExclusiveCampaignsStrategy gets a reference to the given string and assigns it to the ExclusiveCampaignsStrategy field.
func (o *NewApplication) SetExclusiveCampaignsStrategy(v string) {
	o.ExclusiveCampaignsStrategy = &v
}

// GetDefaultDiscountScope returns the DefaultDiscountScope field value if set, zero value otherwise.
func (o *NewApplication) GetDefaultDiscountScope() string {
	if o == nil || o.DefaultDiscountScope == nil {
		var ret string
		return ret
	}
	return *o.DefaultDiscountScope
}

// GetDefaultDiscountScopeOk returns a tuple with the DefaultDiscountScope field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetDefaultDiscountScopeOk() (string, bool) {
	if o == nil || o.DefaultDiscountScope == nil {
		var ret string
		return ret, false
	}
	return *o.DefaultDiscountScope, true
}

// HasDefaultDiscountScope returns a boolean if a field has been set.
func (o *NewApplication) HasDefaultDiscountScope() bool {
	if o != nil && o.DefaultDiscountScope != nil {
		return true
	}

	return false
}

// SetDefaultDiscountScope gets a reference to the given string and assigns it to the DefaultDiscountScope field.
func (o *NewApplication) SetDefaultDiscountScope(v string) {
	o.DefaultDiscountScope = &v
}

// GetEnableCascadingDiscounts returns the EnableCascadingDiscounts field value if set, zero value otherwise.
func (o *NewApplication) GetEnableCascadingDiscounts() bool {
	if o == nil || o.EnableCascadingDiscounts == nil {
		var ret bool
		return ret
	}
	return *o.EnableCascadingDiscounts
}

// GetEnableCascadingDiscountsOk returns a tuple with the EnableCascadingDiscounts field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetEnableCascadingDiscountsOk() (bool, bool) {
	if o == nil || o.EnableCascadingDiscounts == nil {
		var ret bool
		return ret, false
	}
	return *o.EnableCascadingDiscounts, true
}

// HasEnableCascadingDiscounts returns a boolean if a field has been set.
func (o *NewApplication) HasEnableCascadingDiscounts() bool {
	if o != nil && o.EnableCascadingDiscounts != nil {
		return true
	}

	return false
}

// SetEnableCascadingDiscounts gets a reference to the given bool and assigns it to the EnableCascadingDiscounts field.
func (o *NewApplication) SetEnableCascadingDiscounts(v bool) {
	o.EnableCascadingDiscounts = &v
}

// GetEnableFlattenedCartItems returns the EnableFlattenedCartItems field value if set, zero value otherwise.
func (o *NewApplication) GetEnableFlattenedCartItems() bool {
	if o == nil || o.EnableFlattenedCartItems == nil {
		var ret bool
		return ret
	}
	return *o.EnableFlattenedCartItems
}

// GetEnableFlattenedCartItemsOk returns a tuple with the EnableFlattenedCartItems field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetEnableFlattenedCartItemsOk() (bool, bool) {
	if o == nil || o.EnableFlattenedCartItems == nil {
		var ret bool
		return ret, false
	}
	return *o.EnableFlattenedCartItems, true
}

// HasEnableFlattenedCartItems returns a boolean if a field has been set.
func (o *NewApplication) HasEnableFlattenedCartItems() bool {
	if o != nil && o.EnableFlattenedCartItems != nil {
		return true
	}

	return false
}

// SetEnableFlattenedCartItems gets a reference to the given bool and assigns it to the EnableFlattenedCartItems field.
func (o *NewApplication) SetEnableFlattenedCartItems(v bool) {
	o.EnableFlattenedCartItems = &v
}

// GetAttributesSettings returns the AttributesSettings field value if set, zero value otherwise.
func (o *NewApplication) GetAttributesSettings() AttributesSettings {
	if o == nil || o.AttributesSettings == nil {
		var ret AttributesSettings
		return ret
	}
	return *o.AttributesSettings
}

// GetAttributesSettingsOk returns a tuple with the AttributesSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetAttributesSettingsOk() (AttributesSettings, bool) {
	if o == nil || o.AttributesSettings == nil {
		var ret AttributesSettings
		return ret, false
	}
	return *o.AttributesSettings, true
}

// HasAttributesSettings returns a boolean if a field has been set.
func (o *NewApplication) HasAttributesSettings() bool {
	if o != nil && o.AttributesSettings != nil {
		return true
	}

	return false
}

// SetAttributesSettings gets a reference to the given AttributesSettings and assigns it to the AttributesSettings field.
func (o *NewApplication) SetAttributesSettings(v AttributesSettings) {
	o.AttributesSettings = &v
}

// GetSandbox returns the Sandbox field value if set, zero value otherwise.
func (o *NewApplication) GetSandbox() bool {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret
	}
	return *o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetSandboxOk() (bool, bool) {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret, false
	}
	return *o.Sandbox, true
}

// HasSandbox returns a boolean if a field has been set.
func (o *NewApplication) HasSandbox() bool {
	if o != nil && o.Sandbox != nil {
		return true
	}

	return false
}

// SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.
func (o *NewApplication) SetSandbox(v bool) {
	o.Sandbox = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *NewApplication) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplication) GetKeyOk() (string, bool) {
	if o == nil || o.Key == nil {
		var ret string
		return ret, false
	}
	return *o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *NewApplication) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *NewApplication) SetKey(v string) {
	o.Key = &v
}

type NullableNewApplication struct {
	Value        NewApplication
	ExplicitNull bool
}

func (v NullableNewApplication) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewApplication) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
