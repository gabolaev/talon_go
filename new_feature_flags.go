/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package talon

type NewFeatureFlags struct {

	// Whether the account has access to the loyalty features or not
	Loyalty bool `json:"loyalty,omitempty"`

	// Whether the account queries coupons with or without total result size
	CouponsWithoutCount bool `json:"coupons_without_count,omitempty"`

	// Whether the account can test beta effects or not
	BetaEffects bool `json:"betaEffects,omitempty"`
}
