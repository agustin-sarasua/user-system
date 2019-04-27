package service

import (
	"github.com/agustin-sarasua/user-system/src/model"
	"github.com/gin-gonic/gin"
)

func RegisterTenant(gc *gin.Context) {
	var t model.Tenant
	if err := gc.BindJSON(&t); err == nil {
		tenantExists := tenantExists(t)
		if !tenantExists {
			registerTenantAdmin(t)
		}
	}

}

func tenantExists(tenant model.Tenant) bool {
	creds, err := GetSystemCredentials()
	user, err := LookupUserPoolData(creds, tenant.Username, nil, true)
	return user != nil
}

/**
 * Register a new tenant user and provision policies for that user
 * @param tenant The new tenant data
 * @returns {Promise} Results of tenant provisioning
 */
func registerTenantAdmin(tenant model.Tenant) *model.Tenant {
	// Call user service funcion
	creds, err := GetSystemCredentials()
	if err != nil {
	}
	u := &model.User{
		TenantID:   tenant.ID,
		Email:      tenant.Email,
		FirstName:  tenant.FirstName,
		LastName:   tenant.LastName,
		Role:       tenant.Role,
		Tier:       tenant.Tier,
		UserPoolID: tenant.UserPoolId,
		UserName:   tenant.UserName,
	}
	err = ProvisionAdminUserWithRoles(u, creds, TenantAdminUserRol, TenantUserUserRol)
	if err != nil {

	}
	return nil
}

func saveTenantData(tenant model.Tenant) {

}
