package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-01-01/AuthenticationPolicyListAll.json
func ExampleAuthenticationPoliciesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthenticationPoliciesClient().NewListAllPager(nil)
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
		// page = armnetwork.AuthenticationPoliciesClientListAllResponse{
		// 	AuthenticationPolicyListResult: armnetwork.AuthenticationPolicyListResult{
		// 		Value: []*armnetwork.AuthenticationPolicy{
		// 			{
		// 				Name: to.Ptr("authPolicy1"),
		// 				Type: to.Ptr("Microsoft.Network/authenticationPolicies"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/authenticationPolicies/authPolicy1"),
		// 				Location: to.Ptr("westus"),
		// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 				Properties: &armnetwork.AuthenticationPolicyPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 					UserTrustProviderType: to.Ptr(armnetwork.UserTrustProviderTypeEntra),
		// 					OnUnauthenticatedRequest: to.Ptr(armnetwork.OnUnauthenticatedRequestDeny),
		// 					AuthenticationProperties: &armnetwork.AuthenticationProviderProperties{
		// 						Issuer: to.Ptr("https://login.microsoftonline.com/00000000-0000-0000-0000-000000000000/"),
		// 						JwksURI: to.Ptr("https://login.microsoftonline.com/00000000-0000-0000-0000-000000000000/discovery/v2.0/keys"),
		// 						Audience: to.Ptr("api://11111111-1111-1111-1111-111111111111"),
		// 						ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("authPolicy2"),
		// 				Type: to.Ptr("Microsoft.Network/authenticationPolicies"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg2/providers/Microsoft.Network/authenticationPolicies/authPolicy2"),
		// 				Location: to.Ptr("eastus"),
		// 				Etag: to.Ptr("W/\"11111111-1111-1111-1111-111111111111\""),
		// 				Properties: &armnetwork.AuthenticationPolicyPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("33333333-3333-3333-3333-333333333333"),
		// 					UserTrustProviderType: to.Ptr(armnetwork.UserTrustProviderTypeEntra),
		// 					OnUnauthenticatedRequest: to.Ptr(armnetwork.OnUnauthenticatedRequestAuthenticate),
		// 					AuthenticationProperties: &armnetwork.AuthenticationProviderProperties{
		// 						Issuer: to.Ptr("https://login.microsoftonline.com/55555555-5555-5555-5555-555555555555/"),
		// 						ClientID: to.Ptr("44444444-4444-4444-4444-444444444444"),
		// 						ClientSecret: to.Ptr("https://myvault.vault.azure.net/secrets/authpolicy2-secret"),
		// 						Scope: []*string{
		// 							to.Ptr("openid"),
		// 							to.Ptr("profile"),
		// 						},
		// 						SessionTimeout: to.Ptr("86400"),
		// 						SessionCookieName: to.Ptr("ApplicationGatewayUserSession"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
