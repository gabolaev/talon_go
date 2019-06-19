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
type Attribute struct {

	// Unique ID for this entity.
	Id int32 `json:"id"`

	// The exact moment this entity was created.
	Created time.Time `json:"created"`

	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`

	// The name of the entity that can have this attribute. When creating or updating the entities of a given type, you can include an `attributes` object with keys corresponding to the `name` of the custom attributes for that type.
	Entity string `json:"entity"`

	EventType string `json:"eventType,omitempty"`

	// The attribute name that will be used in API requests and Talang. E.g. if `name == \"region\"` then you would set the region attribute by including an `attributes.region` property in your request payload. 
	Name string `json:"name"`

	// The human-readable name for the attribute that will be shown in the Campaign Manager. Like `name`, the combination of entity and title must also be unique.
	Title string `json:"title"`

	// The data type of the attribute, a `time` attribute must be sent as a string that conforms to the [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp format.
	Type_ string `json:"type"`

	// A description of this attribute.
	Description string `json:"description"`

	// A list of suggestions for the attribute.
	Suggestions []string `json:"suggestions"`

	// Whether or not this attribute can be edited.
	Editable bool `json:"editable"`

	// Indicates whether this attribute is in use. If in use only title can be changed and other operations are prohibited.
	Locked bool `json:"locked"`

	// array of rulesets where the attribute is used
	UsedAt []string `json:"usedAt"`
}
