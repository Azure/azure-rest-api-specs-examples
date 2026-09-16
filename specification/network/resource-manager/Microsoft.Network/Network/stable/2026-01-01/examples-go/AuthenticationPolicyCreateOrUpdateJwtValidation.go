package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-01-01/AuthenticationPolicyCreateOrUpdateJwtValidation.json
func ExampleAuthenticationPoliciesClient_BeginCreateOrUpdate_createsOrUpdatesAJwtValidationAuthenticationPolicyWithinAResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAuthenticationPoliciesClient().BeginCreateOrUpdate(ctx, "rg1", "jwtValidationPolicy", armnetwork.AuthenticationPolicy{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.AuthenticationPolicyPropertiesFormat{
			UserTrustProviderType:    to.Ptr(armnetwork.UserTrustProviderTypeEntra),
			OnUnauthenticatedRequest: to.Ptr(armnetwork.OnUnauthenticatedRequestDeny),
			AuthenticationProperties: &armnetwork.AuthenticationProviderProperties{
				Issuer:   to.Ptr("https://login.microsoftonline.com/11111111-1111-1111-1111-111111111111/"),
				ClientID: to.Ptr("00000000-0000-0000-0000-000000000001"),
				JwksURI:  to.Ptr("https://login.microsoftonline.com/11111111-1111-1111-1111-111111111111/discovery/v2.0/keys"),
				Audience: to.Ptr("api://myapp"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.AuthenticationPoliciesClientCreateOrUpdateResponse{
	// 	AuthenticationPolicy: armnetwork.AuthenticationPolicy{
	// 		Name: to.Ptr("jwtValidationPolicy"),
	// 		Type: to.Ptr("Microsoft.Network/authenticationPolicies"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/authenticationPolicies/jwtValidationPolicy"),
	// 		Location: to.Ptr("eastus"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
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
