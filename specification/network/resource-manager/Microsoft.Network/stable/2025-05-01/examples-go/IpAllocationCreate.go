package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/IpAllocationCreate.json
func ExampleIPAllocationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPAllocationsClient().BeginCreateOrUpdate(ctx, "rg1", "test-ipallocation", armnetwork.IPAllocation{
		Location: to.Ptr("centraluseuap"),
		Properties: &armnetwork.IPAllocationPropertiesFormat{
			Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
			AllocationTags: map[string]*string{
				"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1"),
			},
			Prefix: to.Ptr("3.2.5.0/24"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPAllocation = armnetwork.IPAllocation{
	// 	Name: to.Ptr("test-ipallocation"),
	// 	Type: to.Ptr("Microsoft.Network/IpAllocations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/IpAllocations/test-ipallocation"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: &armnetwork.IPAllocationPropertiesFormat{
	// 		Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
	// 		AllocationTags: map[string]*string{
	// 			"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1"),
	// 		},
	// 		IpamAllocationID: to.Ptr("916d3b28-663f-448b-9abc-1bea9d5fed8f"),
	// 		Prefix: to.Ptr("3.2.5.0/24"),
	// 	},
	// }
}
