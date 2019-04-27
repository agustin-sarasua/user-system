package endpoints

import (
	"github.com/agustin-sarasua/user-system/src/model"
	"github.com/agustin-sarasua/user-system/src/service"
	"github.com/gin-gonic/gin"
)

func RegisterTenant(gc *gin.Context) {
	var t model.Tenant
	if err := gc.BindJSON(&t); err == nil {
		tenantExists := service.TenantService.TenantExists(&t)
		if !tenantExists {
			service.TenantService.RegisterTenantAdmin(&t)
		}
	}

}
