package service

import (
	"time"

	"github.com/agustin-sarasua/user-system/src/config"
	"github.com/agustin-sarasua/user-system/src/model"
	"github.com/agustin-sarasua/user-system/src/service/cognito"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

const (
	SystemAdminUserRol = "SystemAdmin"
	SystemUserUserRol  = "SystemUser"
	TenantAdminUserRol = "TenantAdmin"
	TenantUserUserRol  = "TenantUser"
)

/**
 * Lookup a user's pool data in the user table
 * @param credentials The credentials used ben looking up the user
 * @param userId The id of the user being looked up
 * @param tenantId The id of the tenant (if this is not system context)
 * @param isSystemContext Is this being called in the context of a system user (registration, system user provisioning)
 * @param callback The results of the lookup
 */
func LookupUserPoolData(credentials *Credentials, userID string, tenantID string, isSystemContext bool) (error, *model.User) {
	return nil, nil
}

// func GetUserAttributes(gc *gin.Context) {

// }

// func GetUsers(gc *gin.Context) {

// }

// func CreateUser(gc *gin.Context) {

// }

// func CreateTenantAdminUser() {

// }

// func UpdateUserEnabledStatus(gc *gin.Context) {

// }

// func UpdateUserDisabledStatus(gc *gin.Context) {

// }

// func UpdateUserAttributes(gc *gin.Context) {

// }

// func DeleteUser(gc *gin.Context) {

// }

func CreateSystemAdminUser(gc *gin.Context) {
	var user model.User
	if err := gc.BindJSON(&user); err == nil {
		creds, err := GetSystemCredentials()
		if err != nil {
		}
		err = ProvisionAdminUserWithRoles(&user, creds, SystemAdminUserRol, SystemUserUserRol)
		if err != nil {

		}
	}

}

type Credentials struct {
	claims *Claims
}

type Claims struct {
	SessionToken string
	AccessKeyId  string
	SecretKey    string
	Expiration   time.Time
}

func GetSystemCredentials() (*Credentials, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("", "test-account"),
	})

	creds, err := sess.Config.Credentials.Get()
	if err != nil {
		return nil, err
	}

	expiresAt, err := sess.Config.Credentials.ExpiresAt()
	if err != nil {
		return nil, err
	}

	return &Credentials{claims: &Claims{
		SessionToken: creds.SessionToken,
		AccessKeyId:  creds.AccessKeyID,
		SecretKey:    creds.SecretAccessKey,
		Expiration:   expiresAt}}, nil
}

/**
 * Provision an admin user and the associated policies/roles
 * @param user The user being created
 * @param credentials Credentials to use for provisioning
 * @param adminPolicyName The name of of the admin policy to provisioned
 * @param userPolicyName The name of the user policy to be provisioned
 * @param callback Returns an object with the results of the provisioned items
 */
func ProvisionAdminUserWithRoles(user *model.User, credentials *Credentials, adminPolicyName string, userPolicyName string) error {

	cfg := config.Configure("DEVELOPMENT")

	// setup params for template generation
	type policyCreationParams struct {
		TenantID         string
		AccountID        string
		Region           string
		TenantTableName  string
		UserTableName    string
		ProductTableName string
		OrderTableName   string
		UserPoolID       string
	}

	// setup params for template generation
	t := policyCreationParams{
		TenantID:         user.TenantID,
		AccountID:        cfg.AwsAccount,
		Region:           cfg.AwsRegion,
		TenantTableName:  cfg.Table.Tenant,
		UserTableName:    cfg.Table.User,
		ProductTableName: cfg.Table.Product,
		OrderTableName:   cfg.Table.Order,
	}

	// init role based on admin policy name
	user.Role = adminPolicyName

	// see if this user is already in the system
	err, userPoolData := LookupUserPoolData(credentials, user.UserName, user.TenantID, true)
	if err != nil {

	} else {
		// create the new user
		upo := cognito.CreateUserPool(user.TenantID)
		upc := cognito.CreateUserPoolClient(upo.UserPool.Name, upo.UserPool.Id)
	}

	return nil
}
