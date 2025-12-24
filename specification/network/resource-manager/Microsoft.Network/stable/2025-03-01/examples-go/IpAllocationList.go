package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/IpAllocationList.json
func ExampleIPAllocationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIPAllocationsClient().NewListPager(nil)
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
		// page.IPAllocationListResult = armnetwork.IPAllocationListResult{
		// 	Value: []*armnetwork.IPAllocation{
		// 		{
		// 			Name: to.Ptr("test-ipallocation1"),
		// 			Type: to.Ptr("Microsoft.Network/IpAllocations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/IpAllocations/test-ipallocation1"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armnetwork.IPAllocationPropertiesFormat{
		// 				Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
		// 				AllocationTags: map[string]*string{
		// 					"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1"),
		// 				},
		// 				IpamAllocationID: to.Ptr("916d3b28-663f-448b-9abc-1bea9d5fed8f"),
		// 				Prefix: to.Ptr("3.2.5.0/24"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-ipallocation2"),
		// 			Type: to.Ptr("Microsoft.Network/IpAllocations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/IpAllocations/test-ipallocation2"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armnetwork.IPAllocationPropertiesFormat{
		// 				Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
		// 				AllocationTags: map[string]*string{
		// 					"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet2"),
		// 				},
		// 				IpamAllocationID: to.Ptr("57dc7256-2ff7-43f2-b9c8-85a70b5c6408"),
		// 				Prefix: to.Ptr("3.2.6.0/24"),
		// 			},
		// 	}},
		// }
	}
}
