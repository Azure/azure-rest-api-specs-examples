package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7678455846b1000fd31db27596e4ca3d299a872/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/GlobalReachConnections_CreateOrUpdate.json
func ExampleGlobalReachConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGlobalReachConnectionsClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", "connection1", armavs.GlobalReachConnection{
		Properties: &armavs.GlobalReachConnectionProperties{
			AuthorizationKey:        to.Ptr("01010101-0101-0101-0101-010101010101"),
			PeerExpressRouteCircuit: to.Ptr("/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/mygroup/providers/Microsoft.Network/expressRouteCircuits/mypeer"),
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
	// res.GlobalReachConnection = armavs.GlobalReachConnection{
	// 	Name: to.Ptr("connection1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/globalReachConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/globalReachConnections/connection1"),
	// 	Properties: &armavs.GlobalReachConnectionProperties{
	// 		AddressPrefix: to.Ptr("10.2.3.16/29"),
	// 		AuthorizationKey: to.Ptr("01010101-0101-0101-0101-010101010101"),
	// 		CircuitConnectionStatus: to.Ptr(armavs.GlobalReachConnectionStatusConnected),
	// 		ExpressRouteID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2/providers/Microsoft.Network/expressroutecircuits/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2-xconnect"),
	// 		PeerExpressRouteCircuit: to.Ptr("/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/mygroup/providers/Microsoft.Network/expressRouteCircuits/mypeer"),
	// 		ProvisioningState: to.Ptr(armavs.GlobalReachConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
