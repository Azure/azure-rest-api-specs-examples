package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/ExpressRouteGatewayGet.json
func ExampleExpressRouteGatewaysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteGatewaysClient().Get(ctx, "resourceGroupName", "expressRouteGatewayName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRouteGateway = armnetwork.ExpressRouteGateway{
	// 	Name: to.Ptr("expressRouteGatewayName"),
	// 	Type: to.Ptr("Microsoft.Network/expressRouteGateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.ExpressRouteGatewayProperties{
	// 		AllowNonVirtualWanTraffic: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		VirtualHub: &armnetwork.VirtualHubID{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/virtualHubName"),
	// 		},
	// 	},
	// }
}
