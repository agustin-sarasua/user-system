package cognito

import (
	"github.com/agustin-sarasua/user-system/src/config"
	"github.com/agustin-sarasua/user-system/src/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	cip "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type cognitoUserServiceImpl struct {
}

func NewCognitoUserService() *cognitoUserServiceImpl {
	return &cognitoUserServiceImpl{}
}

/**
 * Create a Cognito user with custom attributes
 * @param user User with attribute values
 * @param callback Callback with created user
 */
func (service *cognitoUserServiceImpl) CreateUser(creds *credentials.Credentials, user *model.User) *cip.AdminCreateUserOutput {
	cfg := config.Cfg
	// Create Session with MaxRetry configuration to be shared by multiple
	// service clients.
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries: aws.Int(3),
	}))
	// Create a CognitoIdentityProvider client with additional configuration
	svc := cip.New(sess, aws.NewConfig().WithRegion(cfg.AwsRegion))

	// config the client parameters
	input := &cip.AdminCreateUserInput{
		UserPoolId: user.UserPoolID,
	}
	input = input.SetUsername(user.Username).SetDesiredDeliveryMediums(buildStringPtrArray([]string{
		"email"})).SetForceAliasCreation(true).SetUserAttributes(createAttributeTypes(user))
	out, err := svc.AdminCreateUser(input)
	if err != nil {
	}
	return out
}

func createAttributeTypes(user *model.User) []*cip.AttributeType {
	return []*cip.AttributeType{
		createAttributeType("email", user.Email),
		createAttributeType("custom:tenant_id", *user.TenantID),
		createAttributeType("given_name", user.FirstName),
		createAttributeType("family_name", user.LastName),
		createAttributeType("custom:role", user.Role),
		createAttributeType("custom:tier", user.Tier),
	}
}

func createAttributeType(name string, value string) *cip.AttributeType {
	return &cip.AttributeType{
		Name:  &name,
		Value: &value,
	}
}
