package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkListUsage.json
func ExampleVirtualNetworksClient_NewListUsagePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworksClient().NewListUsagePager("rg1", "vnetName", nil)
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
		// page.VirtualNetworkListUsageResult = armnetwork.VirtualNetworkListUsageResult{
		// 	Value: []*armnetwork.VirtualNetworkUsage{
		// 		{
		// 			Name: &armnetwork.VirtualNetworkUsageName{
		// 				LocalizedValue: to.Ptr("Subnet size and usage"),
		// 				Value: to.Ptr("SubnetSpace"),
		// 			},
		// 			CurrentValue: to.Ptr[float64](-1),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetName/subnets/GatewaySubnet"),
		// 			Limit: to.Ptr[float64](-1),
		// 			Unit: to.Ptr("Count"),
		// 		},
		// 		{
		// 			Name: &armnetwork.VirtualNetworkUsageName{
		// 				LocalizedValue: to.Ptr("Subnet size and usage"),
		// 				Value: to.Ptr("SubnetSpace"),
		// 			},
		// 			CurrentValue: to.Ptr[float64](2),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetName/subnets/newSubnet"),
		// 			Limit: to.Ptr[float64](3),
		// 			Unit: to.Ptr("Count"),
		// 	}},
		// }
	}
}
