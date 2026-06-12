package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualNetworkTapGet.json
func ExampleVirtualNetworkTapsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkTapsClient().Get(ctx, "rg1", "testvtap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualNetworkTapsClientGetResponse{
	// 	VirtualNetworkTap: armnetwork.VirtualNetworkTap{
	// 		Name: to.Ptr("testvtap"),
	// 		Type: to.Ptr("Microsoft.Network/virtualNetworkTaps"),
	// 		Etag: to.Ptr("etag"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkTaps/testvtap"),
	// 		Location: to.Ptr("centraluseuap"),
	// 		Properties: &armnetwork.VirtualNetworkTapPropertiesFormat{
	// 			DestinationNetworkInterfaceIPConfiguration: &armnetwork.InterfaceIPConfiguration{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/testIPConfig1"),
	// 			},
	// 			DestinationPort: to.Ptr[int32](4789),
	// 			NetworkInterfaceTapConfigurations: []*armnetwork.InterfaceTapConfiguration{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface2/tapConfigurations/testtapConfiguration"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningState("Succeded")),
	// 			ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
	// 		},
	// 	},
	// }
}
