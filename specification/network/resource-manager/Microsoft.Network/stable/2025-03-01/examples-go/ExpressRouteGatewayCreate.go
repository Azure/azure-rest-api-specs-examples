package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRouteGatewayCreate.json
func ExampleExpressRouteGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteGatewaysClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "gateway-2", armnetwork.ExpressRouteGateway{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.ExpressRouteGatewayProperties{
			AllowNonVirtualWanTraffic: to.Ptr(false),
			AutoScaleConfiguration: &armnetwork.ExpressRouteGatewayPropertiesAutoScaleConfiguration{
				Bounds: &armnetwork.ExpressRouteGatewayPropertiesAutoScaleConfigurationBounds{
					Min: to.Ptr[int32](3),
				},
			},
			VirtualHub: &armnetwork.VirtualHubID{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupId/providers/Microsoft.Network/virtualHubs/virtualHubName"),
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
	// res.ExpressRouteGateway = armnetwork.ExpressRouteGateway{
	// 	Name: to.Ptr("gateway-2"),
	// 	Type: to.Ptr("Microsoft.Network/expressRouteGateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.ExpressRouteGatewayProperties{
	// 		AllowNonVirtualWanTraffic: to.Ptr(false),
	// 		AutoScaleConfiguration: &armnetwork.ExpressRouteGatewayPropertiesAutoScaleConfiguration{
	// 			Bounds: &armnetwork.ExpressRouteGatewayPropertiesAutoScaleConfigurationBounds{
	// 				Min: to.Ptr[int32](3),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		VirtualHub: &armnetwork.VirtualHubID{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/virtualHubName"),
	// 		},
	// 	},
	// }
}
