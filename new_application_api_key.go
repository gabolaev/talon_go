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
type NewApplicationApiKey struct {

	// ID of the API Key
	Id int32 `json:"id"`

	// ID of user who created
	CreatedBy int32 `json:"createdBy"`

	// Title for API Key
	Title string `json:"title"`

	// ID of account the key is used for
	AccountID int32 `json:"accountID"`

	// ID of application the key is used for
	ApplicationID int32 `json:"applicationID"`

	// The date the API key was created
	Created time.Time `json:"created"`

	// The date the API key expired
	Expires time.Time `json:"expires"`

	// Raw API Key
	Key string `json:"key"`
}
