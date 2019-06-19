/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package talon

// 
type NewCustomerSession struct {

	// ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID.
	ProfileId string `json:"profileId,omitempty"`

	// Any coupon code entered.
	Coupon string `json:"coupon,omitempty"`

	// Any referral code entered.
	Referral string `json:"referral,omitempty"`

	// Indicates the current state of the session. All sessions must start in the \"open\" state, after which valid transitions are...  1. open -> closed 2. open -> cancelled 3. closed -> cancelled 
	State string `json:"state,omitempty"`

	// Serialized JSON representation.
	CartItems []CartItem `json:"cartItems,omitempty"`

	// The total sum of the cart in one session.
	Total float32 `json:"total,omitempty"`

	// A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings. 
	Attributes *interface{} `json:"attributes,omitempty"`
}
