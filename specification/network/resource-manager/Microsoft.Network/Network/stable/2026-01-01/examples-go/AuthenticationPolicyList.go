package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-01-01/AuthenticationPolicyList.json
func ExampleAuthenticationPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthenticationPoliciesClient().NewListPager("rg1", nil)
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
		// page = armnetwork.AuthenticationPoliciesClientListResponse{
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
		// 		},
		// 	},
		// }
	}
}
