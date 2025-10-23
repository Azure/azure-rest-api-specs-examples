package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3855ffb4be0cd4d227b130b67d874fa816736c04/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-08-02-preview/examples/JWTAuthenticators_List.json
func ExampleJWTAuthenticatorsClient_NewListByManagedClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJWTAuthenticatorsClient().NewListByManagedClusterPager("rg1", "clustername1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.JWTAuthenticatorListResult = armcontainerservice.JWTAuthenticatorListResult{
		// 	Value: []*armcontainerservice.JWTAuthenticator{
		// 		{
		// 			Name: to.Ptr("jwt1"),
		// 			Type: to.Ptr("Microsoft.ContainerService/managedClusters/jwtAuthenticators"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/jwtAuthenticators/jwt1"),
		// 			Properties: &armcontainerservice.JWTAuthenticatorProperties{
		// 				ClaimMappings: &armcontainerservice.JWTAuthenticatorClaimMappings{
		// 					Extra: []*armcontainerservice.JWTAuthenticatorExtraClaimMappingExpression{
		// 						{
		// 							Key: to.Ptr("example.com/extrakey"),
		// 							ValueExpression: to.Ptr("claims.customfield"),
		// 					}},
		// 					Groups: &armcontainerservice.JWTAuthenticatorClaimMappingExpression{
		// 						Expression: to.Ptr("claims.groups.split(',').map(group, 'aks:jwt:' + group)"),
		// 					},
		// 					Username: &armcontainerservice.JWTAuthenticatorClaimMappingExpression{
		// 						Expression: to.Ptr("'aks:jwt:' + claims.sub"),
		// 					},
		// 				},
		// 				ClaimValidationRules: []*armcontainerservice.JWTAuthenticatorValidationRule{
		// 					{
		// 						Expression: to.Ptr("has(claims.sub)"),
		// 						Message: to.Ptr("Sub is required"),
		// 					},
		// 					{
		// 						Expression: to.Ptr("claims.sub != ''"),
		// 						Message: to.Ptr("Sub cannot be empty"),
		// 				}},
		// 				Issuer: &armcontainerservice.JWTAuthenticatorIssuer{
		// 					Audiences: []*string{
		// 						to.Ptr("https://example.com/audience1"),
		// 						to.Ptr("https://example.com/audience2")},
		// 						URL: to.Ptr("https://example.com"),
		// 					},
		// 					ProvisioningState: to.Ptr(armcontainerservice.JWTAuthenticatorProvisioningStateSucceeded),
		// 					UserValidationRules: []*armcontainerservice.JWTAuthenticatorValidationRule{
		// 						{
		// 							Expression: to.Ptr("user.groups.all(group, group.startsWith('aks:jwt:admin:'))"),
		// 							Message: to.Ptr("Must be in admin user group"),
		// 					}},
		// 				},
		// 		}},
		// 	}
	}
}
