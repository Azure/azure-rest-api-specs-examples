package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/455d20a5e76d8184f7cff960501a57e1f88986b7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/JWTAuthenticators_Create_Or_Update.json
func ExampleJWTAuthenticatorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJWTAuthenticatorsClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", "jwt1", armcontainerservice.JWTAuthenticator{
		Properties: &armcontainerservice.JWTAuthenticatorProperties{
			ClaimMappings: &armcontainerservice.JWTAuthenticatorClaimMappings{
				Extra: []*armcontainerservice.JWTAuthenticatorExtraClaimMappingExpression{
					{
						Key:             to.Ptr("example.com/extrakey"),
						ValueExpression: to.Ptr("claims.customfield"),
					}},
				Groups: &armcontainerservice.JWTAuthenticatorClaimMappingExpression{
					Expression: to.Ptr("claims.groups.split(',').map(group, 'aks:jwt:' + group)"),
				},
				Username: &armcontainerservice.JWTAuthenticatorClaimMappingExpression{
					Expression: to.Ptr("'aks:jwt:' + claims.sub"),
				},
			},
			ClaimValidationRules: []*armcontainerservice.JWTAuthenticatorValidationRule{
				{
					Expression: to.Ptr("has(claims.sub)"),
					Message:    to.Ptr("Sub is required"),
				},
				{
					Expression: to.Ptr("claims.sub != ''"),
					Message:    to.Ptr("Sub cannot be empty"),
				}},
			Issuer: &armcontainerservice.JWTAuthenticatorIssuer{
				Audiences: []*string{
					to.Ptr("https://example.com/audience1"),
					to.Ptr("https://example.com/audience2")},
				URL: to.Ptr("https://example.com"),
			},
			UserValidationRules: []*armcontainerservice.JWTAuthenticatorValidationRule{
				{
					Expression: to.Ptr("user.groups.all(group, group.startsWith('aks:jwt:admin:'))"),
					Message:    to.Ptr("Must be in admin user group"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JWTAuthenticator = armcontainerservice.JWTAuthenticator{
	// 	Name: to.Ptr("jwt1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/jwtAuthenticators"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/jwtAuthenticators/jwt1"),
	// 	Properties: &armcontainerservice.JWTAuthenticatorProperties{
	// 		ClaimMappings: &armcontainerservice.JWTAuthenticatorClaimMappings{
	// 			Extra: []*armcontainerservice.JWTAuthenticatorExtraClaimMappingExpression{
	// 				{
	// 					Key: to.Ptr("example.com/extrakey"),
	// 					ValueExpression: to.Ptr("claims.customfield"),
	// 			}},
	// 			Groups: &armcontainerservice.JWTAuthenticatorClaimMappingExpression{
	// 				Expression: to.Ptr("claims.groups.split(',').map(group, 'aks:jwt:' + group)"),
	// 			},
	// 			Username: &armcontainerservice.JWTAuthenticatorClaimMappingExpression{
	// 				Expression: to.Ptr("'aks:jwt:' + claims.sub"),
	// 			},
	// 		},
	// 		ClaimValidationRules: []*armcontainerservice.JWTAuthenticatorValidationRule{
	// 			{
	// 				Expression: to.Ptr("has(claims.sub)"),
	// 				Message: to.Ptr("Sub is required"),
	// 			},
	// 			{
	// 				Expression: to.Ptr("claims.sub != ''"),
	// 				Message: to.Ptr("Sub cannot be empty"),
	// 		}},
	// 		Issuer: &armcontainerservice.JWTAuthenticatorIssuer{
	// 			Audiences: []*string{
	// 				to.Ptr("https://example.com/audience1"),
	// 				to.Ptr("https://example.com/audience2")},
	// 				URL: to.Ptr("https://example.com"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerservice.JWTAuthenticatorProvisioningStateSucceeded),
	// 			UserValidationRules: []*armcontainerservice.JWTAuthenticatorValidationRule{
	// 				{
	// 					Expression: to.Ptr("user.groups.all(group, group.startsWith('aks:jwt:admin:'))"),
	// 					Message: to.Ptr("Must be in admin user group"),
	// 			}},
	// 		},
	// 	}
}
