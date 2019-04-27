package service

import (
	"github.com/agustin-sarasua/user-system/src/model"
	"github.com/agustin-sarasua/user-system/src/service/cognito"
	"github.com/agustin-sarasua/user-system/src/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	ci "github.com/aws/aws-sdk-go/service/cognitoidentity"
	cip "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	ddb "github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/gin-gonic/gin"
)

var UserSvc UserService
var TenantSvc TenantService
var CognitoSvc CognitoService
var CognitoUserSvc CognitoUserService
var DynamoHelper DynamoDBHelper

type UserService interface {
	LookupUserPoolData(credentials *credentials.Credentials, userID string, tenantID *string, isSystemContext bool) (*model.User, error)
	CreateSystemAdminUser(gc *gin.Context)
	GetSystemCredentials() (*credentials.Credentials, error)
	ProvisionAdminUserWithRoles(user *model.User, credentials *credentials.Credentials, adminPolicyName string, userPolicyName string) (*model.ProvisionAdminUserWithRolesOutput, error)
}

type TenantService interface {
	TenantExists(tenant *model.Tenant) bool
	RegisterTenantAdmin(tenant *model.Tenant) *model.Tenant
}

type CognitoService interface {
	CreateUserPool(tenantID *string) *cip.CreateUserPoolOutput
	CreateUserPoolClient(clientName *string, userPoolID *string) *cip.CreateUserPoolClientOutput
	CreateIdentityPool(clientID *string, userPoolID *string, name *string) *ci.IdentityPool
	GetPolicyTemplate(tenantID *string, policyType string, userPoolID *string) string
	GetTrustPolicy(identityPoolID *string) string
	CreatePolicy(policyName string, policyDocument string) (*iam.CreatePolicyOutput, error)
	CreateRole(policyDocument string, roleName string) (*iam.CreateRoleOutput, error)
	AddPolicyToRole(policyArn *string, roleName *string) (*iam.AttachRolePolicyOutput, error)
	AddRoleToIdentity(identityPoolID *string, trustAuthRole *string, roleSystem *string, roleSupportOnly *string, clientID *string, provider *string, adminRoleName *string, userRoleName *string) (*ci.SetIdentityPoolRolesOutput, error)
}

type CognitoUserService interface {
	CreateUser(creds *credentials.Credentials, user *model.User) *cip.AdminCreateUserOutput
}

type DynamoDBHelper interface {
	PutItem(user *model.User, creds *credentials.Credentials) *ddb.PutItemOutput
}

func CreateServices() {
	UserSvc = NewUserService()
	TenantSvc = NewTenantService()
	CognitoSvc = cognito.NewCognitoService()
	CognitoUserSvc = cognito.NewCognitoUserService()
	DynamoHelper = dynamodb.NewDynamoDBHelper()
}
