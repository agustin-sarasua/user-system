package cognito

import "fmt"

func getSystemAdminPolicy(params *policyParams) {

}

func getSystemUserPolicy(params *policyParams) {

}

/**
 * Get the IAM policies for a Tenant Admin user
 * @param policyParams Dictionary with configuration parameters
 * @returns The populated tenant user policy template
 */
func getTenantUserPolicy(policyParams *policyParams) string {
	tenantUserPolicyTemplate := fmt.Sprintf(`{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "TenantReadOnlyUserTable",
                "Effect": "Allow",
                "Action": [
                    "dynamodb:GetItem",
                    "dynamodb:BatchGetItem",
                    "dynamodb:Query",
                    "dynamodb:DescribeTable",
                    "dynamodb:CreateTable"

                ],
                "Resource": [%s, %s],
                "Condition": {
                    "ForAllValues:StringEquals": {
                        "dynamodb:LeadingKeys": [%s]
                    }
                }

            },
            {
                "Sid": "ReadWriteOrderTable",
                "Effect": "Allow",
                "Action": [
                    "dynamodb:GetItem",
                    "dynamodb:BatchGetItem",
                    "dynamodb:Query",
                    "dynamodb:PutItem",
                    "dynamodb:UpdateItem",
                    "dynamodb:DeleteItem",
                    "dynamodb:BatchWriteItem",
                    "dynamodb:DescribeTable",
                    "dynamodb:CreateTable"

                ],
                "Resource": [%s],
                "Condition": {
                    "ForAllValues:StringEquals": {
                        "dynamodb:LeadingKeys": [%s]
                    }
                }
            },
            {
                "Sid": "TenantReadOnlyProductTable",
                "Effect": "Allow",
                "Action": [
                    "dynamodb:GetItem",
                    "dynamodb:BatchGetItem",
                    "dynamodb:Query",
                    "dynamodb:DescribeTable",
                    "dynamodb:CreateTable"

                ],
                "Resource": [%s],
                "Condition": {
                    "ForAllValues:StringEquals": {
                        "dynamodb:LeadingKeys": [%s]
                    }
                }
            },
            {
                "Sid": "TenantCognitoAccess",
                "Effect": "Allow",
                "Action": [
                    "cognito-idp:AdminGetUser",
                    "cognito-idp:ListUsers"
                ],
                "Resource": [%s]
            },
        ]
    }`, policyParams.UserTableArn, policyParams.UserTableArn+"/*", *policyParams.TenantID, policyParams.OrderTableArn, *policyParams.TenantID, policyParams.ProductTableArn, *policyParams.TenantID, policyParams.CognitoArn)

	return tenantUserPolicyTemplate
}

/**
 * Get the IAM policies for a Tenant Admin user
 * @param policyParams Dictionary with configuration parameters
 * @returns The populated system admin policy template
 */
func getTenantAdminPolicy(policyParams *policyParams) string {
	tenantAdminPolicyTemplate := fmt.Sprintf(`{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "TenantAdminUserTable",
                "Effect": "Allow",
                "Action": [
                    "dynamodb:GetItem",
                    "dynamodb:BatchGetItem",
                    "dynamodb:Query",
                    "dynamodb:PutItem",
                    "dynamodb:UpdateItem",
                    "dynamodb:DeleteItem",
                    "dynamodb:BatchWriteItem",
                    "dynamodb:DescribeTable",
                    "dynamodb:CreateTable"

                ],
                "Resource": [%s, %s],
                "Condition": {
                    "ForAllValues:StringEquals": {
                        "dynamodb:LeadingKeys": [%s]
                    }
                }
            },
            {
                "Sid": "TenantAdminOrderTable",
                "Effect": "Allow",
                "Action": [
                    "dynamodb:GetItem",
                    "dynamodb:BatchGetItem",
                    "dynamodb:Query",
                    "dynamodb:PutItem",
                    "dynamodb:UpdateItem",
                    "dynamodb:DeleteItem",
                    "dynamodb:BatchWriteItem",
                    "dynamodb:DescribeTable",
                    "dynamodb:CreateTable"

                ],
                "Resource": [%s],
                "Condition": {
                    "ForAllValues:StringEquals": {
                        "dynamodb:LeadingKeys": [%s]
                    }
                }
            },
            {
                "Sid": "TenantAdminProductTable",
                "Effect": "Allow",
                "Action": [
                    "dynamodb:GetItem",
                    "dynamodb:BatchGetItem",
                    "dynamodb:Query",
                    "dynamodb:PutItem",
                    "dynamodb:UpdateItem",
                    "dynamodb:DeleteItem",
                    "dynamodb:BatchWriteItem",
                    "dynamodb:DescribeTable",
                    "dynamodb:CreateTable"

                ],
                "Resource": [%s],
                "Condition": {
                    "ForAllValues:StringEquals": {
                        "dynamodb:LeadingKeys": [%s]
                    }
				}
            },
            {
                "Sid": "TenantCognitoAccess",
                "Effect": "Allow",
                "Action": [
                    "cognito-idp:AdminCreateUser",
                    "cognito-idp:AdminDeleteUser",
                    "cognito-idp:AdminDisableUser",
                    "cognito-idp:AdminEnableUser",
                    "cognito-idp:AdminGetUser",
                    "cognito-idp:ListUsers",
                    "cognito-idp:AdminUpdateUserAttributes"
                ],
                "Resource": [%s]
            },
        ]
	}`, policyParams.UserTableArn, policyParams.UserTableArn+"/*", *policyParams.TenantID, policyParams.OrderTableArn, *policyParams.TenantID, policyParams.ProductTableArn, *policyParams.TenantID, policyParams.CognitoArn)
	return tenantAdminPolicyTemplate
}
