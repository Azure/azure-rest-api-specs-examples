package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationProviderAADClientCred.json
func ExampleAuthorizationProviderClient_CreateOrUpdate_apiManagementCreateAuthorizationProviderAadClientCred() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationProviderClient().CreateOrUpdate(ctx, "rg1", "apimService1", "aadwithclientcred", armapimanagement.AuthorizationProviderContract{
		Properties: &armapimanagement.AuthorizationProviderContractProperties{
			DisplayName:      to.Ptr("aadwithclientcred"),
			IdentityProvider: to.Ptr("aad"),
			Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
				GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: map[string]*string{
						"resourceUri": to.Ptr("https://graph.microsoft.com"),
						"scopes":      to.Ptr("User.Read.All Group.Read.All"),
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
	// 	Name: to.Ptr("aadwithclientcred"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithclientcred"),
	// 	Properties: &armapimanagement.AuthorizationProviderContractProperties{
	// 		DisplayName: to.Ptr("aadwithclientcred"),
	// 		IdentityProvider: to.Ptr("aad"),
	// 		Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
	// 			GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
	// 				ClientCredentials: map[string]*string{
	// 					"loginUri": to.Ptr("https://login.windows.net"),
	// 					"resourceUri": to.Ptr("https://graph.microsoft.com"),
	// 					"scopes": to.Ptr("User.Read.All Group.Read.All"),
	// 					"tenantId": to.Ptr("common"),
	// 				},
	// 			},
	// 			RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
	// 		},
	// 	},
	// }
}
