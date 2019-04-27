package service

import (
	"github.com/agustin-sarasua/user-system/src/model"
)

const (
	TenantAdminUserRol = "TenantAdmin"
	TenantUserUserRol  = "TenantUser"
)

type tenantServiceImpl struct {
}

func NewTenantService() *tenantServiceImpl {
	return &tenantServiceImpl{}
}

func (service *tenantServiceImpl) TenantExists(tenant *model.Tenant) bool {
	creds, _ := UserSvc.GetSystemCredentials()
	user, _ := UserSvc.LookupUserPoolData(creds, tenant.Username, nil, true)
	return user != nil
}

/**
 * Register a new tenant user and provision policies for that user
 * @param tenant The new tenant data
 * @returns {Promise} Results of tenant provisioning
 */
func (service *tenantServiceImpl) RegisterTenantAdmin(tenant *model.Tenant) *model.Tenant {
	// Call user service funcion
	creds, err := UserSvc.GetSystemCredentials()
	if err != nil {
	}
	u := &model.User{
		TenantID:   &tenant.ID,
		Email:      tenant.Email,
		FirstName:  tenant.FirstName,
		LastName:   tenant.LastName,
		Role:       tenant.Role,
		Tier:       tenant.Tier,
		UserPoolID: &tenant.UserPoolID,
		Username:   tenant.Username,
	}
	err = UserSvc.ProvisionAdminUserWithRoles(u, creds, TenantAdminUserRol, TenantUserUserRol)
	if err != nil {

	}
	return nil
}

func (service *tenantServiceImpl) saveTenantData(tenant model.Tenant) {

}
