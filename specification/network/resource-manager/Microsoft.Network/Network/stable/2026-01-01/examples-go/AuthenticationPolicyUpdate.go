package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-01-01/AuthenticationPolicyUpdate.json
func ExampleAuthenticationPoliciesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthenticationPoliciesClient().Update(ctx, "rg1", "authPolicy1", armnetwork.AuthenticationPolicyUpdateParameters{
		Tags: map[string]*string{
			"environment": to.Ptr("production"),
		},
		Identity: &armnetwork.ManagedServiceIdentity{
			Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armnetwork.ManagedServiceIdentityUserAssignedIdentities{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.AuthenticationPoliciesClientUpdateResponse{
	// 	AuthenticationPolicy: armnetwork.AuthenticationPolicy{
	// 		Name: to.Ptr("authPolicy1"),
	// 		Type: to.Ptr("Microsoft.Network/authenticationPolicies"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/authenticationPolicies/authPolicy1"),
	// 		Location: to.Ptr("eastus"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 		Identity: &armnetwork.ManagedServiceIdentity{
	// 			Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armnetwork.ManagedServiceIdentityUserAssignedIdentities{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armnetwork.ManagedServiceIdentityUserAssignedIdentities{
	// 					PrincipalID: to.Ptr("33333333-3333-3333-3333-333333333333"),
	// 					ClientID: to.Ptr("44444444-4444-4444-4444-444444444444"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 		},
	// 		Properties: &armnetwork.AuthenticationPolicyPropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 			UserTrustProviderType: to.Ptr(armnetwork.UserTrustProviderTypeEntra),
	// 			OnUnauthenticatedRequest: to.Ptr(armnetwork.OnUnauthenticatedRequestDeny),
	// 			AuthenticationProperties: &armnetwork.AuthenticationProviderProperties{
	// 				Issuer: to.Ptr("https://login.microsoftonline.com/11111111-1111-1111-1111-111111111111/"),
	// 				ClientID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 				JwksURI: to.Ptr("https://login.microsoftonline.com/11111111-1111-1111-1111-111111111111/discovery/v2.0/keys"),
	// 				Audience: to.Ptr("api://myapp"),
	// 			},
	// 		},
	// 	},
	// }
}
