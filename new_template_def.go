/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package talon

type NewTemplateDef struct {

	// Campaigner-friendly name for the template that will be shown in the rule editor.
	Title string `json:"title"`

	// A short description of the template that will be shown in the rule editor.
	Description string `json:"description,omitempty"`

	// Extended help text for the template.
	Help string `json:"help,omitempty"`

	// Used for grouping templates in the rule editor sidebar.
	Category string `json:"category"`

	// A Talang expression that contains variable bindings referring to args.
	Expr []interface{} `json:"expr"`

	// An array of argument definitions.
	Args []TemplateArgDef `json:"args"`

	// A flag to control exposure in Rule Builder.
	Expose bool `json:"expose,omitempty"`
}
