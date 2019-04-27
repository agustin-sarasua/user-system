package model

import "time"

type User struct {
	ID             string
	UserPoolID     *string
	IdentityPoolID *string
	ClientID       *string
	Username       string
	Email          string
	TenantID       string
	Tier           string
	FirstName      string
	LastName       string
	Role           string
	Sub            string
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

type Credentials struct {
	Claims *Claims
}

type Claims struct {
	SessionToken string
	AccessKeyID  string
	SecretKey    string
	Expiration   time.Time
}

// type TenantAdminData struct {
// 	TenantID    string `json:"tenant_id,omitempty"`
// 	CompanyName string `json:"companyName,omitempty"`
// 	AccountName string `json:"accountName,omitempty"`
// 	OwnerName   string `json:"ownerName,omitempty"`
// 	Tier        string `json:"tier,omitempty"`
// 	Email       string `json:"email,omitempty"`
// 	UserName    string `json:"userName,omitempty"`
// 	Role        string `json:"role,omitempty"`
// 	FirstName   string `json:"firstName,omitempty"`
// 	LastName    string `json:"lastName,omitempty"`
// }
