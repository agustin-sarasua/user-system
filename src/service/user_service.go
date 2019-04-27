package service

import (
	"errors"
	"fmt"

	"github.com/agustin-sarasua/user-system/src/config"
	"github.com/agustin-sarasua/user-system/src/logger"
	"github.com/agustin-sarasua/user-system/src/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

const (
	SystemAdminUserRol = "SystemAdmin"
	SystemUserUserRol  = "SystemUser"
)

type userServiceImpl struct {
}

func NewUserService() *userServiceImpl {
	return &userServiceImpl{}
}

/**
 * Lookup a user's pool data in the user table
 * @param credentials The credentials used ben looking up the user
 * @param userId The id of the user being looked up
 * @param tenantId The id of the tenant (if this is not system context)
 * @param isSystemContext Is this being called in the context of a system user (registration, system user provisioning)
 */
func (service *userServiceImpl) LookupUserPoolData(credentials *credentials.Credentials, userID string, tenantID *string, isSystemContext bool) (*model.User, error) {
	return nil, nil
}

func (service *userServiceImpl) CreateSystemAdminUser(gc *gin.Context) {
	// var user model.User
	// if err := gc.BindJSON(&user); err == nil {
	// 	creds, err := service.GetSystemCredentials()
	// 	if err != nil {
	// 	}
	// 	out, err = service.ProvisionAdminUserWithRoles(&user, creds, SystemAdminUserRol, SystemUserUserRol)
	// 	if err != nil {

	// 	}
	// }

}

func (service *userServiceImpl) GetSystemCredentials() (*credentials.Credentials, error) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("", "test-account"),
	})
	return sess.Config.Credentials, nil
}

/**
 * Provision an admin user and the associated policies/roles
 * @param user The user being created
 * @param credentials Credentials to use for provisioning
 * @param adminPolicyName The name of of the admin policy to provisioned
 * @param userPolicyName The name of the user policy to be provisioned
 */
func (service *userServiceImpl) ProvisionAdminUserWithRoles(user *model.User, credentials *credentials.Credentials, adminPolicyName string, userPolicyName string) (*model.ProvisionAdminUserWithRolesOutput, error) {

	cfg := config.Cfg

	// init role based on admin policy name
	user.Role = adminPolicyName

	// see if this user is already in the system
	err, _ := service.LookupUserPoolData(credentials, user.Username, user.TenantID, true)
	if err == nil {
		logger.Info("User already exists", nil)
		return nil, errors.New("User already exists")
	} else {
		// create the new user
		createdUserPool := CognitoSvc.CreateUserPool(user.TenantID)
		createdUserPoolClient := CognitoSvc.CreateUserPoolClient(createdUserPool.UserPool.Name, createdUserPool.UserPool.Id)
		createdIdentityPool := CognitoSvc.CreateIdentityPool(createdUserPoolClient.UserPoolClient.ClientId, createdUserPoolClient.UserPoolClient.UserPoolId, createdUserPoolClient.UserPoolClient.ClientName)
		// create and populate policy templates
		trustPolicyTemplate := CognitoSvc.GetTrustPolicy(createdIdentityPool.IdentityPoolId)

		// get the admin policy template
		adminPolicyTemplate := CognitoSvc.GetPolicyTemplate(user.TenantID, adminPolicyName, createdUserPool.UserPool.Id)
		policyName := fmt.Sprintf("%s-%sPolicy", user.TenantID, adminPolicyName)
		createdAdminPolicy, _ := CognitoSvc.CreatePolicy(policyName, adminPolicyTemplate)
		service.createNewUser(credentials,
			createdUserPool.UserPool.Id,
			createdIdentityPool.IdentityPoolId,
			createdUserPoolClient.UserPoolClient.ClientId, user.TenantID, user)

		// get the user policy template
		userPolicyTemplate := CognitoSvc.GetPolicyTemplate(user.TenantID, userPolicyName, createdUserPool.UserPool.Id)
		policyName = fmt.Sprintf("%s-%sPolicy", user.TenantID, userPolicyName)
		createdUserPolicy, _ := CognitoSvc.CreatePolicy(policyName, userPolicyTemplate)

		// create Admin Role
		adminRoleName := fmt.Sprintf("%s-%s", user.TenantID, adminPolicyName)
		createdAdminRole, _ := CognitoSvc.CreateRole(trustPolicyTemplate, adminRoleName)

		// create User Role
		userRoleName := fmt.Sprintf("%s-%s", user.TenantID, adminPolicyName)
		createdUserRole, _ := CognitoSvc.CreateRole(trustPolicyTemplate, userRoleName)

		// create Trust Role
		trustPolicyRoleName := fmt.Sprintf("%s-Trust", user.TenantID)
		createdTrustPolicyRole, _ := CognitoSvc.CreateRole(trustPolicyTemplate, trustPolicyRoleName)

		CognitoSvc.AddPolicyToRole(createdAdminPolicy.Policy.Arn, createdAdminRole.Role.RoleName)

		CognitoSvc.AddPolicyToRole(createdUserPolicy.Policy.Arn, createdUserRole.Role.RoleName)

		out, _ := CognitoSvc.AddRoleToIdentity(createdIdentityPool.IdentityPoolId,
			createdTrustPolicyRole.Role.Arn,
			createdAdminRole.Role.Arn,
			createdUserRole.Role.Arn,
			createdUserPoolClient.UserPoolClient.ClientId,
			createdUserPoolClient.UserPoolClient.UserPoolId,
			&adminPolicyName, &userPolicyName)

		return &model.ProvisionAdminUserWithRolesOutput{
			Pool:           createdUserPool,
			UserPoolClient: createdUserPoolClient,
			IdentityPool:   createdIdentityPool,
			Role: map[string]*string{
				"systemAdminRole":   createdAdminRole.Role.RoleName,
				"systemSupportRole": createdUserRole.Role.RoleName,
				"trustRole":         createdTrustPolicyRole.Role.RoleName,
			},
			Policy: map[string]*string{
				"systemAdminPolicy":   createdAdminPolicy.Policy.Arn,
				"systemSupportPolicy": createdUserPolicy.Policy.Arn,
			},
			AddRoleToIdentity: out,
		}, nil
	}
}

/**
 * Create a new user using the supplied credentials/user
 * @param credentials The creds used for the user creation
 * @param userPoolId The user pool where the user will be added
 * @param identityPoolId the identityPoolId
 * @param clientId The client identifier
 * @param tenantId The tenant identifier
 * @param newUser The data fro the user being created
 */
func (service *userServiceImpl) createNewUser(creds *credentials.Credentials, userPoolID *string, identityPoolID *string, clientID *string, tenantID *string, user *model.User) {
	newUser := &model.User{
		UserPoolID: userPoolID,
		TenantID:   tenantID,
		Email:      user.Email,
	}
	cognitoUser := CognitoUserSvc.CreateUser(creds, newUser)
	newUser.ID = newUser.Username
	newUser.UserPoolID = userPoolID
	newUser.IdentityPoolID = identityPoolID
	newUser.ClientID = clientID
	newUser.TenantID = tenantID
	newUser.Sub = cognitoUser.User.Attributes[0].Value

	out := DynamoHelper.PutItem(newUser, creds)
	if out == nil {
		// TODO manage errors
	}
}
