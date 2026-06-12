package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/RouteTableCreateWithEcmpRoute.json
func ExampleRouteTablesClient_BeginCreateOrUpdate_createRouteTableWithEcmpRoute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRouteTablesClient().BeginCreateOrUpdate(ctx, "rg1", "testrt-ecmp", armnetwork.RouteTable{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.RouteTablePropertiesFormat{
			DisableBgpRoutePropagation: to.Ptr(false),
			Routes: []*armnetwork.Route{
				{
					Name: to.Ptr("ecmp-route"),
					Properties: &armnetwork.RoutePropertiesFormat{
						AddressPrefix: to.Ptr("10.1.0.0/16"),
						NextHopType:   to.Ptr(armnetwork.RouteNextHopTypeVirtualApplianceEcmp),
						NextHop: &armnetwork.RouteNextHopEcmp{
							NextHopIPAddresses: []*string{
								to.Ptr("10.0.0.4"),
								to.Ptr("10.0.0.5"),
								to.Ptr("10.0.0.6"),
							},
						},
					},
				},
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
	// res = armnetwork.RouteTablesClientCreateOrUpdateResponse{
	// 	RouteTable: armnetwork.RouteTable{
	// 		Name: to.Ptr("testrt-ecmp"),
	// 		Type: to.Ptr("Microsoft.Network/routeTables"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/routeTables/testrt-ecmp"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetwork.RouteTablePropertiesFormat{
	// 			DisableBgpRoutePropagation: to.Ptr(false),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Routes: []*armnetwork.Route{
	// 				{
	// 					Name: to.Ptr("ecmp-route"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/routeTables/testrt-ecmp/routes/ecmp-route"),
	// 					Properties: &armnetwork.RoutePropertiesFormat{
	// 						AddressPrefix: to.Ptr("10.1.0.0/16"),
	// 						NextHopType: to.Ptr(armnetwork.RouteNextHopTypeVirtualApplianceEcmp),
	// 						NextHop: &armnetwork.RouteNextHopEcmp{
	// 							NextHopIPAddresses: []*string{
	// 								to.Ptr("10.0.0.4"),
	// 								to.Ptr("10.0.0.5"),
	// 								to.Ptr("10.0.0.6"),
	// 							},
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
