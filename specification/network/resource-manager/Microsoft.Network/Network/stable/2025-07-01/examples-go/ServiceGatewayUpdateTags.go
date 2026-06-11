package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ServiceGatewayUpdateTags.json
func ExampleServiceGatewaysClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceGatewaysClient().UpdateTags(ctx, "rg1", "sg", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ServiceGatewaysClientUpdateTagsResponse{
	// 	ServiceGateway: armnetwork.ServiceGateway{
	// 		Name: to.Ptr("sg"),
	// 		Type: to.Ptr("Microsoft.Network/serviceGateways"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/serviceGateways/sg"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetwork.ServiceGatewayPropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			RouteTargetAddress: &armnetwork.RouteTargetAddressPropertiesFormat{
	// 				PrivateIPAddress: to.Ptr("10.0.1.4"),
	// 				PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
	// 				Subnet: &armnetwork.Subnet{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet"),
	// 				},
	// 			},
	// 			VirtualNetwork: &armnetwork.VirtualNetwork{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
