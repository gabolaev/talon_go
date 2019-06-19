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
type Campaign struct {

	// Unique ID for this entity.
	Id int32 `json:"id"`

	// The exact moment this entity was created.
	Created time.Time `json:"created"`

	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`

	// The ID of the account that owns this entity.
	UserId int32 `json:"userId"`

	// A friendly name for this campaign.
	Name string `json:"name"`

	// A detailed description of the campaign.
	Description string `json:"description"`

	// Datetime when the campaign will become active.
	StartTime time.Time `json:"startTime,omitempty"`

	// Datetime when the campaign will become in-active.
	EndTime time.Time `json:"endTime,omitempty"`

	// Arbitrary properties associated with this campaign
	Attributes *interface{} `json:"attributes,omitempty"`

	// A disabled or archived campaign is not evaluated for rules or coupons. 
	State string `json:"state"`

	// ID of Ruleset this campaign applies on customer session evaluation.
	ActiveRulesetId int32 `json:"activeRulesetId,omitempty"`

	// A list of tags for the campaign.
	Tags []string `json:"tags"`

	// A list of features for the campaign.
	Features []string `json:"features"`

	CouponSettings *CodeGeneratorSettings `json:"couponSettings,omitempty"`

	ReferralSettings *CodeGeneratorSettings `json:"referralSettings,omitempty"`

	// The set of limits that will operate for this campaign
	Limits []LimitConfig `json:"limits"`

	// Number of coupons redeemed in the campaign.
	CouponRedemptionCount int32 `json:"couponRedemptionCount,omitempty"`

	// Number of referral codes redeemed in the campaign.
	ReferralRedemptionCount int32 `json:"referralRedemptionCount,omitempty"`

	// Total amount of discounts redeemed in the campaign.
	DiscountCount int32 `json:"discountCount,omitempty"`

	// Timestamp of the most recent event received by this campaign.
	LastActivity time.Time `json:"lastActivity,omitempty"`

	// Timestamp of the most recent update to the campaign or any of its elements.
	Updated time.Time `json:"updated,omitempty"`

	// Name of the user who created this campaign if available.
	CreatedBy string `json:"createdBy,omitempty"`

	// Name of the user who last updated this campaign if available.
	UpdatedBy string `json:"updatedBy,omitempty"`
}
