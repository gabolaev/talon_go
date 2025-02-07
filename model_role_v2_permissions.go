/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}`
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// RoleV2Permissions struct for RoleV2Permissions
type RoleV2Permissions struct {
	// List of grouped operation IDs to use as a reference in the roles section. Each group of operation IDs has a name.
	PermissionSets *[]RoleV2PermissionSet  `json:"permissionSets,omitempty"`
	Roles          *RoleV2PermissionsRoles `json:"roles,omitempty"`
}

// GetPermissionSets returns the PermissionSets field value if set, zero value otherwise.
func (o *RoleV2Permissions) GetPermissionSets() []RoleV2PermissionSet {
	if o == nil || o.PermissionSets == nil {
		var ret []RoleV2PermissionSet
		return ret
	}
	return *o.PermissionSets
}

// GetPermissionSetsOk returns a tuple with the PermissionSets field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RoleV2Permissions) GetPermissionSetsOk() ([]RoleV2PermissionSet, bool) {
	if o == nil || o.PermissionSets == nil {
		var ret []RoleV2PermissionSet
		return ret, false
	}
	return *o.PermissionSets, true
}

// HasPermissionSets returns a boolean if a field has been set.
func (o *RoleV2Permissions) HasPermissionSets() bool {
	if o != nil && o.PermissionSets != nil {
		return true
	}

	return false
}

// SetPermissionSets gets a reference to the given []RoleV2PermissionSet and assigns it to the PermissionSets field.
func (o *RoleV2Permissions) SetPermissionSets(v []RoleV2PermissionSet) {
	o.PermissionSets = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *RoleV2Permissions) GetRoles() RoleV2PermissionsRoles {
	if o == nil || o.Roles == nil {
		var ret RoleV2PermissionsRoles
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RoleV2Permissions) GetRolesOk() (RoleV2PermissionsRoles, bool) {
	if o == nil || o.Roles == nil {
		var ret RoleV2PermissionsRoles
		return ret, false
	}
	return *o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *RoleV2Permissions) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given RoleV2PermissionsRoles and assigns it to the Roles field.
func (o *RoleV2Permissions) SetRoles(v RoleV2PermissionsRoles) {
	o.Roles = &v
}

type NullableRoleV2Permissions struct {
	Value        RoleV2Permissions
	ExplicitNull bool
}

func (v NullableRoleV2Permissions) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRoleV2Permissions) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
