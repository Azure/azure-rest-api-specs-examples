package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteGatewayCreate.json
func ExampleExpressRouteGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRouteGatewaysClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "resourceGroupName", "gateway-2", armnetwork.ExpressRouteGateway{
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
	// TODO: use response item
	_ = res
}
