package model

type User struct {
	UserPoolID string
	UserName   string
	Email      string
	TenantID   string
	Tier       string
	FirstName  string
	LastName   string
	Role       string
}

type Tenant struct {
	ID                  string
	Username            string
	UserPoolId          string
	IdentityPoolId      string
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
	UserName            string
	Role                string
	FirstName           string
	LastName            string
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
