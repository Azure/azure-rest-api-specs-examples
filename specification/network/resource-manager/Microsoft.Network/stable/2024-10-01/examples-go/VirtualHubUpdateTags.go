package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a9dbb28e788355a47dc5bad3ea5f8da212b4bf6/specification/network/resource-manager/Microsoft.Network/stable/2024-10-01/examples/VirtualHubUpdateTags.json
func ExampleVirtualHubsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualHubsClient().UpdateTags(ctx, "rg1", "virtualHub2", armnetwork.TagsObject{
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualHub = armnetwork.VirtualHub{
	// 	Name: to.Ptr("virtualHub2"),
	// 	Type: to.Ptr("Microsoft.Network/virtualHubs"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VirtualHubProperties{
	// 		AddressPrefix: to.Ptr("10.168.0.0/24"),
	// 		AllowBranchToBranchTraffic: to.Ptr(false),
	// 		HubRoutingPreference: to.Ptr(armnetwork.HubRoutingPreferenceExpressRoute),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		SKU: to.Ptr("Basic"),
	// 		VirtualHubRouteTableV2S: []*armnetwork.VirtualHubRouteTableV2{
	// 		},
	// 		VirtualRouterAsn: to.Ptr[int64](65515),
	// 		VirtualRouterAutoScaleConfiguration: &armnetwork.VirtualRouterAutoScaleConfiguration{
	// 			MinCapacity: to.Ptr[int32](2),
	// 		},
	// 		VirtualRouterIPs: []*string{
	// 			to.Ptr("10.10.1.12"),
	// 			to.Ptr("10.10.1.13")},
	// 			VirtualWan: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
	// 			},
	// 		},
	// 	}
}
