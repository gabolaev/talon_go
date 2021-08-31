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

// InlineResponse20018 struct for InlineResponse20018
type InlineResponse20018 struct {
	HasMore         *bool      `json:"hasMore,omitempty"`
	TotalResultSize *int32     `json:"totalResultSize,omitempty"`
	Data            []Audience `json:"data"`
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *InlineResponse20018) GetHasMore() bool {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018) GetHasMoreOk() (bool, bool) {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret, false
	}
	return *o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *InlineResponse20018) HasHasMore() bool {
	if o != nil && o.HasMore != nil {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *InlineResponse20018) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResultSize returns the TotalResultSize field value if set, zero value otherwise.
func (o *InlineResponse20018) GetTotalResultSize() int32 {
	if o == nil || o.TotalResultSize == nil {
		var ret int32
		return ret
	}
	return *o.TotalResultSize
}

// GetTotalResultSizeOk returns a tuple with the TotalResultSize field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018) GetTotalResultSizeOk() (int32, bool) {
	if o == nil || o.TotalResultSize == nil {
		var ret int32
		return ret, false
	}
	return *o.TotalResultSize, true
}

// HasTotalResultSize returns a boolean if a field has been set.
func (o *InlineResponse20018) HasTotalResultSize() bool {
	if o != nil && o.TotalResultSize != nil {
		return true
	}

	return false
}

// SetTotalResultSize gets a reference to the given int32 and assigns it to the TotalResultSize field.
func (o *InlineResponse20018) SetTotalResultSize(v int32) {
	o.TotalResultSize = &v
}

// GetData returns the Data field value
func (o *InlineResponse20018) GetData() []Audience {
	if o == nil {
		var ret []Audience
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *InlineResponse20018) SetData(v []Audience) {
	o.Data = v
}

type NullableInlineResponse20018 struct {
	Value        InlineResponse20018
	ExplicitNull bool
}

func (v NullableInlineResponse20018) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineResponse20018) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
