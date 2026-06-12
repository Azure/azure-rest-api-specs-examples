package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/LocalNetworkGatewayList.json
func ExampleLocalNetworkGatewaysClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocalNetworkGatewaysClient().NewListPager("rg1", nil)
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
		// page = armnetwork.LocalNetworkGatewaysClientListResponse{
		// 	LocalNetworkGatewayListResult: armnetwork.LocalNetworkGatewayListResult{
		// 		Value: []*armnetwork.LocalNetworkGateway{
		// 			{
		// 				Name: to.Ptr("localgw1"),
		// 				Type: to.Ptr("Microsoft.Network/localNetworkGateways"),
		// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw1"),
		// 				Location: to.Ptr("centralus"),
		// 				Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
		// 					GatewayIPAddress: to.Ptr("x.x.x.x"),
		// 					LocalNetworkAddressSpace: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("10.1.0.0/16"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("localgw2"),
		// 				Type: to.Ptr("Microsoft.Network/localNetworkGateways"),
		// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw2"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
		// 					GatewayIPAddress: to.Ptr("x.x.x.x"),
		// 					LocalNetworkAddressSpace: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("10.2.0.0/16"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
