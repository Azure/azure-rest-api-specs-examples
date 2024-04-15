package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_List.json
func ExampleProfilesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProfilesClient().NewListPager(nil)
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
		// page.ProfileListResult = armcdn.ProfileListResult{
		// 	Value: []*armcdn.Profile{
		// 		{
		// 			Name: to.Ptr("profile1"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG1/providers/Microsoft.Cdn/profiles/profile1"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr("frontdoor"),
		// 			Properties: &armcdn.ProfileProperties{
		// 				FrontDoorID: to.Ptr("3b4682da-b3e2-47a1-96ca-08ab3cb7294e"),
		// 				LogScrubbing: &armcdn.ProfileLogScrubbing{
		// 					ScrubbingRules: []*armcdn.ProfileScrubbingRules{
		// 					},
		// 					State: to.Ptr(armcdn.ProfileScrubbingStateEnabled),
		// 				},
		// 				OriginResponseTimeoutSeconds: to.Ptr[int32](30),
		// 				ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
		// 				ResourceState: to.Ptr(armcdn.ProfileResourceStateActive),
		// 			},
		// 			SKU: &armcdn.SKU{
		// 				Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("profile2"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG1/providers/Microsoft.Cdn/profiles/profile2"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr("frontdoor"),
		// 			Properties: &armcdn.ProfileProperties{
		// 				FrontDoorID: to.Ptr("3b4682da-b3e2-47a1-96ca-08ab3cb7294e"),
		// 				LogScrubbing: &armcdn.ProfileLogScrubbing{
		// 					ScrubbingRules: []*armcdn.ProfileScrubbingRules{
		// 					},
		// 					State: to.Ptr(armcdn.ProfileScrubbingStateEnabled),
		// 				},
		// 				OriginResponseTimeoutSeconds: to.Ptr[int32](30),
		// 				ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
		// 				ResourceState: to.Ptr(armcdn.ProfileResourceStateActive),
		// 			},
		// 			SKU: &armcdn.SKU{
		// 				Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
		// 			},
		// 	}},
		// }
	}
}
