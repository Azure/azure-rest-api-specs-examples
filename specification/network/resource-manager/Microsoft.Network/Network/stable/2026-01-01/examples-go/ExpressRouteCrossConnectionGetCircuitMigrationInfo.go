package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-01-01/ExpressRouteCrossConnectionGetCircuitMigrationInfo.json
func ExampleExpressRouteCrossConnectionsClient_BeginGetCircuitMigrationInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCrossConnectionsClient().BeginGetCircuitMigrationInfo(ctx, "CrossConnection-SiliconValley", "<circuitServiceKey>", armnetwork.MigrateExpressRouteCircuitValidateAndHealthCheckRequest{
		TargetPeeringLocation: to.Ptr("SiliconValley"),
		TargetPortMapping: []*armnetwork.PortMapping{
			{
				SourcePortID: to.Ptr("sourcePort1"),
				TargetPortID: to.Ptr("targetPort1"),
			},
			{
				SourcePortID: to.Ptr("sourcePort2"),
				TargetPortID: to.Ptr("targetPort2"),
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
	// res = armnetwork.ExpressRouteCrossConnectionsClientGetCircuitMigrationInfoResponse{
	// 	MigrateExpressRouteCircuitHealthCheckResponse: armnetwork.MigrateExpressRouteCircuitHealthCheckResponse{
	// 		Status: to.Ptr("Succeeded"),
	// 		Phase: to.Ptr("Prepared"),
	// 		FailureReason: to.Ptr(""),
	// 		NewSTag: to.Ptr("12345"),
	// 		PreparedAt: to.Ptr(time.Date(2026, time.August, 5, 11, 30, 0, 0, time.UTC)),
	// 		PrepareExpiryTime: to.Ptr(time.Date(2026, time.August, 5, 12, 30, 0, 0, time.UTC)),
	// 		NewCrossConnectionURL: to.Ptr("https://management.azure.com/subscriptions/subid/resourceGroups/CrossConnection-SiliconValley/providers/Microsoft.Network/expressRouteCrossConnections/newCircuitServiceKey?api-version=2026-01-01"),
	// 		ShouldRollback: to.Ptr(false),
	// 		Details: &armnetwork.MigrateExpressRouteCircuitHealthCheckDetails{
	// 			PortMigrationInfos: []*armnetwork.PortMigrationInfo{
	// 				{
	// 					PortID: to.Ptr("targetPort1"),
	// 					Status: to.Ptr("Ready"),
	// 					Phase: to.Ptr("Prepared"),
	// 					FailureReason: to.Ptr(""),
	// 					Peerings: []*armnetwork.PeeringHealth{
	// 						{
	// 							Type: to.Ptr("Private"),
	// 							StatsCurrent: &armnetwork.PeeringStats{
	// 								Timestamp: to.Ptr(time.Date(2026, time.August, 5, 11, 31, 0, 0, time.UTC)),
	// 								Metrics: []*armnetwork.Metric{
	// 									{
	// 										Name: to.Ptr("PacketsIn"),
	// 										Value: to.Ptr[float64](100),
	// 										Unit: to.Ptr("count"),
	// 									},
	// 								},
	// 							},
	// 							StatsAtPrepare: &armnetwork.PeeringStats{
	// 								Timestamp: to.Ptr(time.Date(2026, time.August, 5, 11, 30, 0, 0, time.UTC)),
	// 								Metrics: []*armnetwork.Metric{
	// 									{
	// 										Name: to.Ptr("PacketsIn"),
	// 										Value: to.Ptr[float64](98),
	// 										Unit: to.Ptr("count"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 					SourcePortID: to.Ptr("sourcePort1"),
	// 					SourcePortStats: &armnetwork.SourcePortStats{
	// 						Peerings: []*armnetwork.PeeringHealth{
	// 							{
	// 								Type: to.Ptr("Private"),
	// 								StatsCurrent: &armnetwork.PeeringStats{
	// 									Timestamp: to.Ptr(time.Date(2026, time.August, 5, 11, 31, 0, 0, time.UTC)),
	// 									Metrics: []*armnetwork.Metric{
	// 										{
	// 											Name: to.Ptr("PacketsIn"),
	// 											Value: to.Ptr[float64](100),
	// 											Unit: to.Ptr("count"),
	// 										},
	// 									},
	// 								},
	// 								StatsAtPrepare: &armnetwork.PeeringStats{
	// 									Timestamp: to.Ptr(time.Date(2026, time.August, 5, 11, 30, 0, 0, time.UTC)),
	// 									Metrics: []*armnetwork.Metric{
	// 										{
	// 											Name: to.Ptr("PacketsIn"),
	// 											Value: to.Ptr[float64](98),
	// 											Unit: to.Ptr("count"),
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				{
	// 					PortID: to.Ptr("targetPort2"),
	// 					Status: to.Ptr("Ready"),
	// 					Phase: to.Ptr("Prepared"),
	// 					FailureReason: to.Ptr(""),
	// 					Peerings: []*armnetwork.PeeringHealth{
	// 						{
	// 							Type: to.Ptr("Private"),
	// 							StatsCurrent: &armnetwork.PeeringStats{
	// 								Timestamp: to.Ptr(time.Date(2026, time.August, 5, 11, 31, 0, 0, time.UTC)),
	// 								Metrics: []*armnetwork.Metric{
	// 									{
	// 										Name: to.Ptr("PacketsIn"),
	// 										Value: to.Ptr[float64](97),
	// 										Unit: to.Ptr("count"),
	// 									},
	// 								},
	// 							},
	// 							StatsAtPrepare: &armnetwork.PeeringStats{
	// 								Timestamp: to.Ptr(time.Date(2026, time.August, 5, 11, 30, 0, 0, time.UTC)),
	// 								Metrics: []*armnetwork.Metric{
	// 									{
	// 										Name: to.Ptr("PacketsIn"),
	// 										Value: to.Ptr[float64](96),
	// 										Unit: to.Ptr("count"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 					SourcePortID: to.Ptr("sourcePort2"),
	// 					SourcePortStats: &armnetwork.SourcePortStats{
	// 						Peerings: []*armnetwork.PeeringHealth{
	// 							{
	// 								Type: to.Ptr("Private"),
	// 								StatsCurrent: &armnetwork.PeeringStats{
	// 									Timestamp: to.Ptr(time.Date(2026, time.August, 5, 11, 31, 0, 0, time.UTC)),
	// 									Metrics: []*armnetwork.Metric{
	// 										{
	// 											Name: to.Ptr("PacketsIn"),
	// 											Value: to.Ptr[float64](97),
	// 											Unit: to.Ptr("count"),
	// 										},
	// 									},
	// 								},
	// 								StatsAtPrepare: &armnetwork.PeeringStats{
	// 									Timestamp: to.Ptr(time.Date(2026, time.August, 5, 11, 30, 0, 0, time.UTC)),
	// 									Metrics: []*armnetwork.Metric{
	// 										{
	// 											Name: to.Ptr("PacketsIn"),
	// 											Value: to.Ptr[float64](96),
	// 											Unit: to.Ptr("count"),
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
