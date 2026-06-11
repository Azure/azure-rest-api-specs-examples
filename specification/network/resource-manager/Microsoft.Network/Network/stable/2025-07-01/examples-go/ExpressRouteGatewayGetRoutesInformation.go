package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteGatewayGetRoutesInformation.json
func ExampleExpressRouteGatewaysClient_BeginGetRoutesInformation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteGatewaysClient().BeginGetRoutesInformation(ctx, "rg1", "ergw1", &armnetwork.ExpressRouteGatewaysClientBeginGetRoutesInformationOptions{
		AttemptRefresh: to.Ptr(false)})
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
	// res = armnetwork.ExpressRouteGatewaysClientGetRoutesInformationResponse{
	// 	GatewayRouteSetsInformation: armnetwork.GatewayRouteSetsInformation{
	// 		LastComputedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-01T00:00:00Z"); return t}()),
	// 		NextEligibleComputeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-01T01:00:00Z"); return t}()),
	// 		RouteSetVersion: to.Ptr("1"),
	// 		RouteSets: []*armnetwork.GatewayRouteSet{
	// 			{
	// 				Name: to.Ptr("Set-1"),
	// 				Locations: []*string{
	// 					to.Ptr("Washington DC"),
	// 					to.Ptr("London2"),
	// 				},
	// 				Details: map[string][]*armnetwork.RouteSourceDetails{
	// 					"10.11.0.0/24": []*armnetwork.RouteSourceDetails{
	// 						{
	// 							Circuit: to.Ptr("x0"),
	// 							Pri: to.Ptr("1"),
	// 							Sec: to.Ptr("1"),
	// 						},
	// 						{
	// 							Circuit: to.Ptr("x1"),
	// 							Pri: to.Ptr("2"),
	// 							Sec: to.Ptr("2"),
	// 						},
	// 					},
	// 					"10.57.0.0/16": []*armnetwork.RouteSourceDetails{
	// 						{
	// 							Circuit: to.Ptr("x0"),
	// 							Pri: to.Ptr("1"),
	// 							Sec: to.Ptr("1"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		CircuitsMetadataMap: map[string]*armnetwork.CircuitMetadataMap{
	// 			"x0": &armnetwork.CircuitMetadataMap{
	// 				Name: to.Ptr("circuit1"),
	// 				Link: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuit1"),
	// 				Location: to.Ptr("Hong Kong"),
	// 			},
	// 			"x1": &armnetwork.CircuitMetadataMap{
	// 				Name: to.Ptr("circuit2"),
	// 				Link: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuit2"),
	// 				Location: to.Ptr("Hong Kong2"),
	// 			},
	// 		},
	// 	},
	// }
}
