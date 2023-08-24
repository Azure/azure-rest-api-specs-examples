package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationProviderGenericOAuth2.json
func ExampleAuthorizationProviderClient_CreateOrUpdate_apiManagementCreateAuthorizationProviderGenericOAuth2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationProviderClient().CreateOrUpdate(ctx, "rg1", "apimService1", "eventbrite", armapimanagement.AuthorizationProviderContract{
		Properties: &armapimanagement.AuthorizationProviderContractProperties{
			DisplayName:      to.Ptr("eventbrite"),
			IdentityProvider: to.Ptr("oauth2"),
			Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
				GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: map[string]*string{
						"authorizationUrl": to.Ptr("https://www.eventbrite.com/oauth/authorize"),
						"clientId":         to.Ptr("ZYIJTBTABHOUQQDLZY"),
						"clientSecret":     to.Ptr("Q3iPSaKQ~fZFcJk5vKmqzUAfJagcJ8"),
						"refreshUrl":       to.Ptr("https://www.eventbrite.com/oauth/token"),
						"scopes":           nil,
						"tokenUrl":         to.Ptr("https://www.eventbrite.com/oauth/token"),
					},
				},
				RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
			},
		},
	}, &armapimanagement.AuthorizationProviderClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationProviderContract = armapimanagement.AuthorizationProviderContract{
	// 	Name: to.Ptr("eventbrite"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/eventbrite"),
	// 	Properties: &armapimanagement.AuthorizationProviderContractProperties{
	// 		DisplayName: to.Ptr("eventbrite"),
	// 		IdentityProvider: to.Ptr("oauth2"),
	// 		Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
	// 			GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
	// 				AuthorizationCode: map[string]*string{
	// 					"authorizationUrl": to.Ptr("https://www.eventbrite.com/oauth/authorize"),
	// 					"clientId": to.Ptr("ZYIJTBTABHOUQQDLZY"),
	// 					"refreshUrl": to.Ptr("https://www.eventbrite.com/oauth/token"),
	// 					"scopes": nil,
	// 					"tokenUrl": to.Ptr("https://www.eventbrite.com/oauth/token"),
	// 				},
	// 			},
	// 			RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
	// 		},
	// 	},
	// }
}
