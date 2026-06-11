package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteGatewayCreate.json
func ExampleExpressRouteGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupId/providers/Microsoft.Network/virtualHubs/virtualHubName"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ExpressRouteGatewaysClientCreateOrUpdateResponse{
	// 	ExpressRouteGateway: armnetwork.ExpressRouteGateway{
	// 		Name: to.Ptr("gateway-2"),
	// 		Type: to.Ptr("Microsoft.Network/expressRouteGateways"),
	// 		Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetwork.ExpressRouteGatewayProperties{
	// 			AllowNonVirtualWanTraffic: to.Ptr(false),
	// 			AutoScaleConfiguration: &armnetwork.ExpressRouteGatewayPropertiesAutoScaleConfiguration{
	// 				Bounds: &armnetwork.ExpressRouteGatewayPropertiesAutoScaleConfigurationBounds{
	// 					Min: to.Ptr[int32](3),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			VirtualHub: &armnetwork.VirtualHubID{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/virtualHubName"),
	// 			},
	// 		},
	// 	},
	// }
}
