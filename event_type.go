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
type EventType struct {

	// Unique ID for this entity.
	Id int32 `json:"id"`

	// The exact moment this entity was created.
	Created time.Time `json:"created"`

	// The IDs of the applications that are related to this entity.
	ApplicationIds []int32 `json:"applicationIds"`

	// The human-friendly display name for this event type. Use a short, past-tense, description of the event.
	Title string `json:"title"`

	// The machine-friendly canonical name for this event type. This will be used in URLs, and cannot be changed after an event type has been created.
	Name string `json:"name"`

	// An explanation of when the event type is triggered. Write this with a campaign manager in mind. For example:  > The \"Payment Accepted\" event is triggered after successful processing of a payment by our payment gateway. 
	Description string `json:"description"`

	// This defines how the request payload will be parsed before your handler code is run.
	MimeType string `json:"mimeType"`

	// It is often helpful to include an example payload with the event type definition for documentation purposes.
	ExamplePayload string `json:"examplePayload,omitempty"`

	// It is strongly recommended to define a JSON schema that will be used to perform structural validation of request payloads after parsing. 
	Schema *interface{} `json:"schema,omitempty"`

	// The language of the handler code. Currently only `\"talang\"` is supported.
	HandlerLanguage string `json:"handlerLanguage,omitempty"`

	// Code that will be run after successful parsing & validation of the payload for this event. This code _may_ choose to evaluate campaign rules. 
	Handler string `json:"handler"`

	// The version of this event type. When updating an existing event type this must be **exactly** `currentVersion + 1`. 
	Version int32 `json:"version"`
}
