/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package talon

import (
	"time"
)

// 
type Application struct {

	// Unique ID for this entity.
	Id int32 `json:"id"`

	// The exact moment this entity was created.
	Created time.Time `json:"created"`

	// The exact moment this entity was last modified.
	Modified time.Time `json:"modified"`

	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`

	// The name of this application.
	Name string `json:"name"`

	// A longer description of the application.
	Description string `json:"description,omitempty"`

	// Hex key for HMAC-signing API calls as coming from this application (16 hex digits)
	Key string `json:"key"`

	// A string containing an IANA timezone descriptor.
	Timezone string `json:"timezone"`

	// A string describing a default currency for new customer sessions.
	Currency string `json:"currency"`

	// A string indicating how should campaigns in this application deal with case sensitivity on coupon codes.
	CaseSensitivity string `json:"caseSensitivity,omitempty"`

	// Arbitrary properties associated with this campaign
	Attributes *interface{} `json:"attributes,omitempty"`

	// An array containing all the loyalty programs to which this application is subscribed
	LoyaltyPrograms []LoyaltyProgram `json:"loyaltyPrograms"`
}
