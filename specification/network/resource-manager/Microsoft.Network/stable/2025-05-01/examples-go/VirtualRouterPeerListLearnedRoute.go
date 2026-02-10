package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualRouterPeerListLearnedRoute.json
func ExampleVirtualHubBgpConnectionsClient_BeginListLearnedRoutes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubBgpConnectionsClient().BeginListLearnedRoutes(ctx, "rg1", "virtualRouter1", "peer1", nil)
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
	// res.Value = map[string][]*armnetwork.PeerRoute{
	// 	"RouteServiceRole_IN_0": []*armnetwork.PeerRoute{
	// 		{
	// 			AsPath: to.Ptr("65002-65001"),
	// 			LocalAddress: to.Ptr("10.85.3.4"),
	// 			Network: to.Ptr("10.101.0.0/16"),
	// 			NextHop: to.Ptr("10.85.4.4"),
	// 			Origin: to.Ptr("EBgp"),
	// 			SourcePeer: to.Ptr("10.85.4.4"),
	// 			Weight: to.Ptr[int32](32768),
	// 		},
	// 		{
	// 			AsPath: to.Ptr("65002-65001"),
	// 			LocalAddress: to.Ptr("10.85.3.5"),
	// 			Network: to.Ptr("10.101.0.0/16"),
	// 			NextHop: to.Ptr("10.85.4.4"),
	// 			Origin: to.Ptr("EBgp"),
	// 			SourcePeer: to.Ptr("10.85.4.4"),
	// 			Weight: to.Ptr[int32](32768),
	// 	}},
	// 	"RouteServiceRole_IN_1": []*armnetwork.PeerRoute{
	// 		{
	// 			AsPath: to.Ptr("65002-65001"),
	// 			LocalAddress: to.Ptr("10.85.3.4"),
	// 			Network: to.Ptr("10.101.0.0/16"),
	// 			NextHop: to.Ptr("10.85.4.4"),
	// 			Origin: to.Ptr("EBgp"),
	// 			SourcePeer: to.Ptr("10.85.4.4"),
	// 			Weight: to.Ptr[int32](32768),
	// 		},
	// 		{
	// 			AsPath: to.Ptr("65002-65001"),
	// 			LocalAddress: to.Ptr("10.85.3.5"),
	// 			Network: to.Ptr("10.101.0.0/16"),
	// 			NextHop: to.Ptr("10.85.4.4"),
	// 			Origin: to.Ptr("EBgp"),
	// 			SourcePeer: to.Ptr("10.85.4.4"),
	// 			Weight: to.Ptr[int32](32768),
	// 	}},
	// }
}
