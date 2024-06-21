package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/VirtualWANListByResourceGroup.json
func ExampleVirtualWansClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualWansClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.ListVirtualWANsResult = armnetwork.ListVirtualWANsResult{
		// 	Value: []*armnetwork.VirtualWAN{
		// 		{
		// 			Name: to.Ptr("wan1"),
		// 			Type: to.Ptr("Microsoft.Network/virtualWANs"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWANs/wan1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.VirtualWanProperties{
		// 				Type: to.Ptr("Basic"),
		// 				DisableVPNEncryption: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				VirtualHubs: []*armnetwork.SubResource{
		// 					{
		// 					},
		// 					{
		// 				}},
		// 				VPNSites: []*armnetwork.SubResource{
		// 					{
		// 					},
		// 					{
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("wan2"),
		// 			Type: to.Ptr("Microsoft.Network/virtualWANs"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWANs/wan2"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.VirtualWanProperties{
		// 				Type: to.Ptr("Basic"),
		// 				DisableVPNEncryption: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				VirtualHubs: []*armnetwork.SubResource{
		// 					{
		// 					},
		// 					{
		// 				}},
		// 				VPNSites: []*armnetwork.SubResource{
		// 					{
		// 					},
		// 					{
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
