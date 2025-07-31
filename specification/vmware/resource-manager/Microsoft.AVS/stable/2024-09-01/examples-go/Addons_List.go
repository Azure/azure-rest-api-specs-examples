package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2024-09-01/Addons_List.json
func ExampleAddonsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAddonsClient().NewListPager("group1", "cloud1", nil)
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
		// page = armavs.AddonsClientListResponse{
		// 	AddonList: armavs.AddonList{
		// 		Value: []*armavs.Addon{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/addons/srm"),
		// 				Name: to.Ptr("srm"),
		// 				Properties: &armavs.AddonSrmProperties{
		// 					AddonType: to.Ptr(armavs.AddonTypeSRM),
		// 					LicenseKey: to.Ptr("41915178-A8FF-4A4D-B683-6D735AF5E3F5"),
		// 					ProvisioningState: to.Ptr(armavs.AddonProvisioningStateSucceeded),
		// 				},
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/addons"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
