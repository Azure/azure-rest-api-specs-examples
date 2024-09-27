package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/VirtualNetworkTapCreate.json
func ExampleVirtualNetworkTapsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkTapsClient().BeginCreateOrUpdate(ctx, "rg1", "test-vtap", armnetwork.VirtualNetworkTap{
		Location: to.Ptr("centraluseuap"),
		Properties: &armnetwork.VirtualNetworkTapPropertiesFormat{
			DestinationNetworkInterfaceIPConfiguration: &armnetwork.InterfaceIPConfiguration{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/ipconfig1"),
			},
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
	// res.VirtualNetworkTap = armnetwork.VirtualNetworkTap{
	// 	Name: to.Ptr("testvtap"),
	// 	Type: to.Ptr("Microsoft.Network/virtualNetworkTaps"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkTaps/testvtap"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Etag: to.Ptr("etag"),
	// 	Properties: &armnetwork.VirtualNetworkTapPropertiesFormat{
	// 		DestinationNetworkInterfaceIPConfiguration: &armnetwork.InterfaceIPConfiguration{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/testIPConfig1"),
	// 		},
	// 		DestinationPort: to.Ptr[int32](4789),
	// 		NetworkInterfaceTapConfigurations: []*armnetwork.InterfaceTapConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface2/tapConfigurations/testtapConfiguration"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
	// 	},
	// }
}
