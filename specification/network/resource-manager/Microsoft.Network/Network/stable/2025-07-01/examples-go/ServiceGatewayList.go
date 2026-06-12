package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ServiceGatewayList.json
func ExampleServiceGatewaysClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceGatewaysClient().NewListPager("rg1", nil)
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
		// page = armnetwork.ServiceGatewaysClientListResponse{
		// 	ServiceGatewayListResult: armnetwork.ServiceGatewayListResult{
		// 		Value: []*armnetwork.ServiceGateway{
		// 			{
		// 				Name: to.Ptr("sg"),
		// 				Type: to.Ptr("Microsoft.Network/serviceGateways"),
		// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/serviceGateways/sg"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armnetwork.ServiceGatewayPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 					RouteTargetAddress: &armnetwork.RouteTargetAddressPropertiesFormat{
		// 						PrivateIPAddress: to.Ptr("10.0.1.4"),
		// 						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
		// 						Subnet: &armnetwork.Subnet{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet"),
		// 						},
		// 					},
		// 					VirtualNetwork: &armnetwork.VirtualNetwork{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sg2"),
		// 				Type: to.Ptr("Microsoft.Network/serviceGateways"),
		// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/serviceGateways/sg2"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armnetwork.ServiceGatewayPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 					RouteTargetAddress: &armnetwork.RouteTargetAddressPropertiesFormat{
		// 						PrivateIPAddress: to.Ptr("10.0.1.4"),
		// 						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
		// 						Subnet: &armnetwork.Subnet{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet"),
		// 						},
		// 					},
		// 					VirtualNetwork: &armnetwork.VirtualNetwork{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
