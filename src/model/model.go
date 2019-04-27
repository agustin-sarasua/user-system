package model

import (
	ci "github.com/aws/aws-sdk-go/service/cognitoidentity"
	cip "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

//ddb "github.com/aws/aws-sdk-go/service/dynamodb")

type User struct {
	ID             string
	UserPoolID     *string
	IdentityPoolID *string
	ClientID       *string
	Username       string
	Email          string
	TenantID       *string
	Tier           string
	FirstName      string
	LastName       string
	Role           string
	Sub            *string
}

type Tenant struct {
	ID                  string
	Username            string
	UserPoolID          string
	IdentityPoolID      string
	SystemAdminRole     string
	SystemSupportRole   string
	TrustRole           string
	SystemAdminPolicy   string
	SystemSupportPolicy string
	CompanyName         string
	AccountName         string
	OwnerName           string
	Tier                string
	Email               string
	Role                string
	FirstName           string
	LastName            string
	Sub                 string
}

type ProvisionAdminUserWithRolesOutput struct {
	Pool              *cip.CreateUserPoolOutput
	UserPoolClient    *cip.CreateUserPoolClientOutput
	IdentityPool      *ci.IdentityPool
	Role              map[string]*string
	Policy            map[string]*string
	AddRoleToIdentity *ci.SetIdentityPoolRolesOutput
}
