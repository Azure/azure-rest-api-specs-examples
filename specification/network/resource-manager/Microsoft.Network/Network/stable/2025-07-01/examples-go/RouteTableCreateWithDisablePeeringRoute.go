package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/RouteTableCreateWithDisablePeeringRoute.json
func ExampleRouteTablesClient_BeginCreateOrUpdate_createRouteTableWithDisablePeeringRoute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRouteTablesClient().BeginCreateOrUpdate(ctx, "rg1", "testrt", armnetwork.RouteTable{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.RouteTablePropertiesFormat{
			DisableBgpRoutePropagation: to.Ptr(true),
			DisablePeeringRoute:        to.Ptr(armnetwork.DisablePeeringRouteAll),
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
	// res = armnetwork.RouteTablesClientCreateOrUpdateResponse{
	// 	RouteTable: armnetwork.RouteTable{
	// 		Name: to.Ptr("testrt"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/routeTables/testrt"),
	// 		Type: to.Ptr("Microsoft.Network/routeTables"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetwork.RouteTablePropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			DisableBgpRoutePropagation: to.Ptr(true),
	// 			DisablePeeringRoute: to.Ptr(armnetwork.DisablePeeringRouteAll),
	// 			Routes: []*armnetwork.Route{
	// 			},
	// 		},
	// 	},
	// }
}
