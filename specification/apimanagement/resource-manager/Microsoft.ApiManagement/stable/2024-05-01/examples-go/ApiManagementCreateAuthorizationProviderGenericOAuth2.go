package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateAuthorizationProviderGenericOAuth2.json
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
						"clientId":         to.Ptr("clientid"),
						"clientSecret":     to.Ptr("clientsecretvalue"),
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
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/eventbrite"),
	// 	Properties: &armapimanagement.AuthorizationProviderContractProperties{
	// 		DisplayName: to.Ptr("eventbrite"),
	// 		IdentityProvider: to.Ptr("oauth2"),
	// 		Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
	// 			GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
	// 				AuthorizationCode: map[string]*string{
	// 					"authorizationUrl": to.Ptr("https://www.eventbrite.com/oauth/authorize"),
	// 					"clientId": to.Ptr("clientid"),
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
