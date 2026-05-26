package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListAuthorizationProviders.json
func ExampleAuthorizationProviderClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthorizationProviderClient().NewListByServicePager("rg1", "apimService1", nil)
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
		// page = armapimanagement.AuthorizationProviderClientListByServiceResponse{
		// 	AuthorizationProviderCollection: armapimanagement.AuthorizationProviderCollection{
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.AuthorizationProviderContract{
		// 			{
		// 				Name: to.Ptr("aadwithauthcode"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode"),
		// 				Properties: &armapimanagement.AuthorizationProviderContractProperties{
		// 					DisplayName: to.Ptr("aadwithauthcode"),
		// 					IdentityProvider: to.Ptr("aad"),
		// 					Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
		// 						GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
		// 							AuthorizationCode: map[string]*string{
		// 								"clientId": to.Ptr("53790825-fdd3-4b80-bc7a-4c3aaf25801d"),
		// 								"loginUri": to.Ptr("https://login.windows.net"),
		// 								"resourceUri": to.Ptr("https://graph.microsoft.com"),
		// 								"scopes": to.Ptr("User.Read.All Group.Read.All"),
		// 								"tenantId": to.Ptr("common"),
		// 							},
		// 						},
		// 						RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("aadwithkeyvault"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithkeyvault"),
		// 				Properties: &armapimanagement.AuthorizationProviderContractProperties{
		// 					DisplayName: to.Ptr("Azure AD with Key Vault"),
		// 					IdentityProvider: to.Ptr("aad"),
		// 					Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
		// 						GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
		// 							AuthorizationCode: map[string]*string{
		// 								"clientId": to.Ptr("53790825-fdd3-4b80-bc7a-4c3aaf25801d"),
		// 								"resourceUri": to.Ptr("https://graph.microsoft.com"),
		// 								"scopes": to.Ptr("User.Read.All Group.Read.All"),
		// 							},
		// 						},
		// 						KeyVault: &armapimanagement.AuthorizationProviderKeyVaultContract{
		// 							LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
		// 								Code: to.Ptr("Success"),
		// 								Message: to.Ptr("Successfully refreshed secret from Key Vault."),
		// 								TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-08T10:30:45.1234567Z"); return t}()),
		// 							},
		// 							SecretIdentifier: to.Ptr("https://my.vault.azure.net/secrets/clientSecret"),
		// 							Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-08T10:30:45.1234567Z"); return t}()),
		// 						},
		// 						RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("aadwithclientcred"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithclientcred"),
		// 				Properties: &armapimanagement.AuthorizationProviderContractProperties{
		// 					DisplayName: to.Ptr("aadwithclientcred"),
		// 					IdentityProvider: to.Ptr("aad"),
		// 					Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
		// 						GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
		// 							ClientCredentials: map[string]*string{
		// 								"loginUri": to.Ptr("https://login.windows.net"),
		// 								"resourceUri": to.Ptr("https://graph.microsoft.com"),
		// 								"scopes": to.Ptr("User.Read.All Group.Read.All"),
		// 								"tenantId": to.Ptr("common"),
		// 							},
		// 						},
		// 						RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("google"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/google"),
		// 				Properties: &armapimanagement.AuthorizationProviderContractProperties{
		// 					DisplayName: to.Ptr("google"),
		// 					IdentityProvider: to.Ptr("google"),
		// 					Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
		// 						GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
		// 							AuthorizationCode: map[string]*string{
		// 								"clientId": to.Ptr("99999999-xxxxxxxxxxxxxxxxxxxxxxxx.apps.googleusercontent.com"),
		// 								"scopes": to.Ptr("openid https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email"),
		// 							},
		// 						},
		// 						RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("eventbrite"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/eventbrite"),
		// 				Properties: &armapimanagement.AuthorizationProviderContractProperties{
		// 					DisplayName: to.Ptr("eventbrite"),
		// 					IdentityProvider: to.Ptr("oauth2"),
		// 					Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
		// 						GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
		// 							AuthorizationCode: map[string]*string{
		// 								"authorizationUrl": to.Ptr("https://www.eventbrite.com/oauth/authorize"),
		// 								"clientId": to.Ptr("clientid"),
		// 								"refreshUrl": to.Ptr("https://www.eventbrite.com/oauth/token"),
		// 								"tokenUrl": to.Ptr("https://www.eventbrite.com/oauth/token"),
		// 							},
		// 						},
		// 						RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
