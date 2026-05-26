package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementRefreshAuthorizationProviderSecret.json
func ExampleAuthorizationProviderClient_RefreshSecret() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationProviderClient().RefreshSecret(ctx, "rg1", "apimService1", "aadwithkeyvault", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.AuthorizationProviderClientRefreshSecretResponse{
	// 	AuthorizationProviderContract: armapimanagement.AuthorizationProviderContract{
	// 		Name: to.Ptr("aadwithkeyvault"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithkeyvault"),
	// 		Properties: &armapimanagement.AuthorizationProviderContractProperties{
	// 			DisplayName: to.Ptr("Azure AD with Key Vault"),
	// 			IdentityProvider: to.Ptr("aad"),
	// 			Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
	// 				GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
	// 					AuthorizationCode: map[string]*string{
	// 						"clientId": to.Ptr("53790825-fdd3-4b80-bc7a-4c3aaf25801d"),
	// 						"scopes": to.Ptr("User.Read.All Group.Read.All"),
	// 						"tenantId": to.Ptr("common"),
	// 					},
	// 				},
	// 				KeyVault: &armapimanagement.AuthorizationProviderKeyVaultContract{
	// 					LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
	// 						Code: to.Ptr("Success"),
	// 						Message: to.Ptr("Successfully refreshed secret from Key Vault."),
	// 						TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-08T10:30:45.1234567Z"); return t}()),
	// 					},
	// 					SecretIdentifier: to.Ptr("https://my.vault.azure.net/secrets/clientSecret"),
	// 					Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-08T10:30:45.1234567Z"); return t}()),
	// 				},
	// 				RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
	// 			},
	// 		},
	// 	},
	// }
}
