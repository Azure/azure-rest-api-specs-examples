package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/AFDProfiles_Upgrade.json
func ExampleAFDProfilesClient_BeginUpgrade() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAFDProfilesClient().BeginUpgrade(ctx, "RG", "profile1", armcdn.ProfileUpgradeParameters{
		WafMappingList: []*armcdn.ProfileChangeSKUWafMapping{
			{
				ChangeToWafPolicy: &armcdn.ResourceReference{
					ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoorwebapplicationfirewallpolicies/waf2"),
				},
				SecurityPolicyName: to.Ptr("securityPolicy1"),
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
	// res = armcdn.AFDProfilesClientUpgradeResponse{
	// 	Profile: armcdn.Profile{
	// 		Name: to.Ptr("profile1"),
	// 		Type: to.Ptr("Microsoft.Cdn/profiles"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1"),
	// 		Kind: to.Ptr("frontdoor"),
	// 		Location: to.Ptr("Global"),
	// 		Properties: &armcdn.ProfileProperties{
	// 			ExtendedProperties: map[string]*string{
	// 			},
	// 			FrontDoorID: to.Ptr("id"),
	// 			OriginResponseTimeoutSeconds: to.Ptr[int32](60),
	// 			ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
	// 			ResourceState: to.Ptr(armcdn.ProfileResourceState("Enabled")),
	// 		},
	// 		SKU: &armcdn.SKU{
	// 			Name: to.Ptr(armcdn.SKUNameStandardAzureFrontDoor),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
