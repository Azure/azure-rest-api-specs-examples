package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteCircuitARPTableList.json
func ExampleExpressRouteCircuitsClient_BeginListArpTable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCircuitsClient().BeginListArpTable(ctx, "rg1", "circuitName", "peeringName", "devicePath", nil)
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
	// res = armnetwork.ExpressRouteCircuitsClientListArpTableResponse{
	// 	ExpressRouteCircuitsArpTableListResult: armnetwork.ExpressRouteCircuitsArpTableListResult{
	// 		Value: []*armnetwork.ExpressRouteCircuitArpTable{
	// 			{
	// 				Age: to.Ptr[int32](0),
	// 				Interface: to.Ptr("Microsoft"),
	// 				IPAddress: to.Ptr("IPAddress"),
	// 				MacAddress: to.Ptr("macAddress"),
	// 			},
	// 		},
	// 	},
	// }
}
