package routers

import (
	"github.com/agustin-sarasua/user-system/src/service"
	"github.com/gin-gonic/gin"
)

func configureMappings(router *gin.RouterGroup) {
	// System Registration Management
	// Create System Admin

	// Tenant Registration
	// Register new Tenant -> Calls Tenant Manager
	router.GET("reg", service.LookupUserPool)

	// Tenant Manager -> Calls DynamoDB

	// User Management -> Calls Cognito
	// Lookup UserPool for any user - no user data returned
	router.GET("user/pool/:id", service.LookupUserPool)
	// Get user attributes
	router.GET("user/:id", service.GetUserAttributes)
	// Get a list of users using a tenant id to scope the list
	router.GET("users", service.GetUsers)
	// Create a new user
	router.POST("user", service.CreateUser)
	// Provision a new system admin user
	router.POST("user/system", service.CreateSystemAdminUser)
	// Provision a new tenant admin user
	router.POST("user", service.CreateTenantAdminUser)
	// Enable a user that is currently disabled
	router.PUT("user/enable", service.UpdateUserEnabledStatus)
	// Disable a user that is currently enabled
	router.PUT("user/disable", service.UpdateUserDisabledStatus)
	// Update a user's attributes
	router.PUT("user", service.UpdateUserAttributes)
	// Delete a user
	router.DELETE("user", service.DeleteUser)

	// router.POST("domain_discovery/category_predictor_wrapper/:site_id/predict", service.PredictLikePredictorPost)
	// router.GET("domain_discovery/category_predictor_wrapper/:site_id/predict", service.PredictLikePredictorGet)
	// router.GET("domain_discovery/config/:scope", service.GetKvsConfigForScope)
	// router.PUT("domain_discovery/config/:scope", service.PutKvsConfigForScope)
	// router.DELETE("domain_discovery/config/:scope", service.DelKvsConfigForScope)
	// router.POST("domain_discovery/update_deleted_categories", service.UpdateDeletedCategoriesDump)

	// // predictor endpoints
	// router.POST("/sites/:site_id/category_predictor/predict", service.PredictLikePredictorPost)
	// router.GET("/sites/:site_id/category_predictor/predict", service.PredictLikePredictorGet)
}
