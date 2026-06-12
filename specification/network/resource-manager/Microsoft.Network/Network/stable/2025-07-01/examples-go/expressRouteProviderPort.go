package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/expressRouteProviderPort.json
func ExampleManagementClient_ExpressRouteProviderPort() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().ExpressRouteProviderPort(ctx, "abc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ManagementClientExpressRouteProviderPortResponse{
	// 	ExpressRouteProviderPort: armnetwork.ExpressRouteProviderPort{
	// 		Type: to.Ptr("Microsoft.Network/expressRouteProviderPort"),
	// 		Etag: to.Ptr("W/\"c0e6477e-8150-4d4f-9bf6-bb10e6acb63a\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/ExpressRoutePortsLocations/SiliconValley/bvtazureixpportpair1"),
	// 		Location: to.Ptr("uswest"),
	// 		Properties: &armnetwork.ExpressRouteProviderPortProperties{
	// 			OverprovisionFactor: to.Ptr[int32](4),
	// 			PeeringLocation: to.Ptr("SiliconValley"),
	// 			PortBandwidthInMbps: to.Ptr[int32](4000),
	// 			PortPairDescriptor: to.Ptr("bvtazureixpportpair1"),
	// 			PrimaryAzurePort: to.Ptr("bvtazureixp01"),
	// 			RemainingBandwidthInMbps: to.Ptr[int32](1500),
	// 			SecondaryAzurePort: to.Ptr("bvtazureixp01"),
	// 			UsedBandwidthInMbps: to.Ptr[int32](2500),
	// 		},
	// 	},
	// }
}
