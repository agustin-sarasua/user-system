package cognito

import (
	"github.com/agustin-sarasua/user-system/src/config"
	"github.com/agustin-sarasua/user-system/src/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	ci "github.com/aws/aws-sdk-go/service/cognitoidentity"
	cip "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

/**
 * Create a Cognito Identity Pool with the supplied params
 * @param clientConfigParams The client config params
 */
func CreateIdentityPool(clientID string, userPoolID string, name string) *ci.IdentityPool {
	cfg := config.Configure("DEVELOPMENT")
	// Create Session with MaxRetry configuration to be shared by multiple
	// service clients.
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries: aws.Int(3),
	}))
	// Create a CognitoIdentityProvider client with additional configuration
	svc := ci.New(sess, aws.NewConfig().WithRegion(cfg.AwsRegion))

	provider := "cognito-idp." + cfg.AwsRegion + ".amazonaws.com/" + userPoolID
	input := &ci.CreateIdentityPoolInput{}
	p := &ci.Provider{}
	p = p.SetClientId(clientID).SetProviderName(provider).SetServerSideTokenCheck(true)
	input = input.SetAllowUnauthenticatedIdentities(false).SetIdentityPoolName(name).SetCognitoIdentityProviders([]*ci.Provider{p})

	out, err := svc.CreateIdentityPool(input)
	if err != nil {
	}
	return out
}

/**
 * Create a user pool client for a new tenant
 */
func CreateUserPoolClient(clientName *string, userPoolID *string) *cip.CreateUserPoolClientOutput {
	cfg := config.Configure("DEVELOPMENT")
	// Create Session with MaxRetry configuration to be shared by multiple
	// service clients.
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries: aws.Int(3),
	}))
	// Create a CognitoIdentityProvider client with additional configuration
	svc := cip.New(sess, aws.NewConfig().WithRegion(cfg.AwsRegion))

	// config the client parameters
	input := &cip.CreateUserPoolClientInput{
		ClientName: clientName,
		UserPoolId: userPoolID,
	}
	input = input.SetGenerateSecret(false).SetReadAttributes(buildStringPtrArray([]string{
		"email",
		"family_name",
		"given_name",
		"phone_number",
		"preferred_username",
		"custom:tier",
		"custom:tenant_id",
		"custom:company_name",
		"custom:account_name",
		"custom:role"})).SetRefreshTokenValidity(0).SetWriteAttributes(buildStringPtrArray([]string{
		"email",
		"family_name",
		"given_name",
		"phone_number",
		"preferred_username",
		"custom:tier",
		"custom:role"}))
	out, err := svc.CreateUserPoolClient(input)
	if err != nil {
		logger.Error("", err)
	}
	return out
}

func buildStringPtrArray(strs []string) []*string {
	ss := make([]*string, len(strs))
	for i, str := range strs {
		ss[i] = &str
	}
	return ss
}

func strPtr(v string) *string {
	return &v
}

/**
 * Create a new User Pool for a new tenant
 * @param tenantId The ID of the new tenant
 * @param callback Callback with created tenant results
 */
func CreateUserPool(tenantID string) *cip.CreateUserPoolOutput {
	cfg := config.Configure("DEVELOPMENT")
	// Create Session with MaxRetry configuration to be shared by multiple
	// service clients.
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries: aws.Int(3),
	}))
	// Create a CognitoIdentityProvider client with additional configuration
	svc := cip.New(sess, aws.NewConfig().WithRegion(cfg.AwsRegion))

	snsArn := cfg.Role.Sns

	input := &cip.CreateUserPoolInput{}
	input.SetPoolName(tenantID)
	input.SetAdminCreateUserConfig(createDefaultAdminCreateUserConfig())
	phoneNumber := "phone_number"
	email := "email"
	input.SetAliasAttributes([]*string{&phoneNumber})
	input.SetAutoVerifiedAttributes([]*string{&email, &phoneNumber})
	input.SetMfaConfiguration("OFF")
	input.SetPolicies(createDefaultPolicies())
	input.SetSchema(createDefaultSchemas())
	input.SetSmsConfiguration(createDefaultSmsConfig(snsArn))
	input.SetUserPoolTags(map[string]*string{"someKey": &tenantID})

	output, err := svc.CreateUserPool(input)
	if err != nil {
		logger.Error("", err)
	}
	return output
}

func createDefaultAdminCreateUserConfig() *cip.AdminCreateUserConfigType {
	inviteMessage := `<img src="https://d0.awsstatic.com/partner-network/logo_apn.png" alt="AWSPartner"> <br><br>Welcome to the Jungle <br><br>Login to the Multi-Tenant Identity Reference Architecture. <br><br>Username: {username} <br><br>Password: {####}`
	emailSubject := "UserSystem-SaaS-Identity-Cognito"
	adminCreateUserConfig := &cip.AdminCreateUserConfigType{}
	adminCreateUserConfig.SetAllowAdminCreateUserOnly(true)
	adminCreateUserConfig.SetUnusedAccountValidityDays(12)
	adminCreateUserConfig.SetInviteMessageTemplate(createDefaultMessageTemplate(inviteMessage, emailSubject))
	return adminCreateUserConfig
}

func createDefaultMessageTemplate(inviteMessage string, emailSubject string) *cip.MessageTemplateType {
	msgTemaplateType := &cip.MessageTemplateType{}
	msgTemaplateType.SetEmailMessage(inviteMessage)
	msgTemaplateType.SetEmailSubject(emailSubject)
	return msgTemaplateType
}

func createDefaultSmsConfig(snsArn string) *cip.SmsConfigurationType {
	v := cip.SmsConfigurationType{}
	v.SetSnsCallerArn(snsArn)
	v.SetExternalId("QuickStartTest")
	return &v
}

func createDefaultPolicies() *cip.UserPoolPolicyType {
	pt := &cip.UserPoolPolicyType{}
	ppt := &cip.PasswordPolicyType{}
	ppt.SetMinimumLength(8)
	ppt.SetRequireLowercase(true)
	ppt.SetRequireNumbers(true)
	ppt.SetRequireSymbols(false)
	ppt.SetRequireUppercase(true)
	pt.SetPasswordPolicy(ppt)
	return pt
}

func createDefaultSchemas() []*cip.SchemaAttributeType {
	// tenant_id
	tenantIDSchema := createDefaultSchema("tenant_id", false, false)
	// tier
	tierSchema := createDefaultSchema("tier", true, false)
	// email
	emailSchema := &cip.SchemaAttributeType{}
	emailSchema.SetName("email")
	emailSchema.SetRequired(false)
	// company_name
	companyNameSchema := createDefaultSchema("company_name", true, false)
	// role
	roleSchema := createDefaultSchema("role", true, false)
	//account_name
	accountNameSchema := createDefaultSchema("account_name", true, false)
	schema := []*cip.SchemaAttributeType{tenantIDSchema, tierSchema, emailSchema, companyNameSchema, roleSchema, accountNameSchema}
	return schema
}

func createDefaultSchema(name string, mutable bool, required bool) *cip.SchemaAttributeType {
	sct1 := cip.SchemaAttributeType{}
	sct1.SetAttributeDataType("String")
	sct1.SetDeveloperOnlyAttribute(false)
	sct1.SetMutable(mutable)
	sct1.SetName(name)
	n := &cip.NumberAttributeConstraintsType{}
	n.SetMaxValue("256")
	n.SetMinValue("1")
	sct1.SetNumberAttributeConstraints(n)
	sct1.SetRequired(required)
	s := &cip.StringAttributeConstraintsType{}
	s.SetMaxLength("256")
	s.SetMinLength("1")
	sct1.SetStringAttributeConstraints(s)

	return &sct1
}
