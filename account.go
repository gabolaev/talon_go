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
type Account struct {

	// Unique ID for this entity.
	Id int32 `json:"id"`

	// The exact moment this entity was created.
	Created time.Time `json:"created"`

	// The exact moment this entity was last modified.
	Modified time.Time `json:"modified"`

	CompanyName string `json:"companyName"`

	// Subdomain Name for yourcompany.talon.one
	DomainName string `json:"domainName"`

	// State of the account (trial, active, trial_expired)
	State string `json:"state"`

	// The billing email address associated with your company account.
	BillingEmail string `json:"billingEmail"`

	// The name of your booked plan.
	PlanName string `json:"planName,omitempty"`

	// The point in time at which your current plan expires.
	PlanExpires time.Time `json:"planExpires,omitempty"`

	// The maximum number of Applications covered by your plan.
	ApplicationLimit int32 `json:"applicationLimit,omitempty"`

	// The maximum number of Campaign Manager Users covered by your plan.
	UserLimit int32 `json:"userLimit,omitempty"`

	// The maximum number of Campaigns covered by your plan.
	CampaignLimit int32 `json:"campaignLimit,omitempty"`

	// The maximum number of Integration API calls covered by your plan per billing period.
	ApiLimit int32 `json:"apiLimit,omitempty"`

	// The current number of Applications in your account.
	ApplicationCount int32 `json:"applicationCount"`

	// The current number of Campaign Manager Users in your account.
	UserCount int32 `json:"userCount"`

	// The current number of active Campaigns in your account.
	CampaignsActiveCount int32 `json:"campaignsActiveCount"`

	// The current number of inactive Campaigns in your account.
	CampaignsInactiveCount int32 `json:"campaignsInactiveCount"`

	// Arbitrary properties associated with this campaign
	Attributes *interface{} `json:"attributes,omitempty"`
}
