package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/GetInboundRoutes.json
func ExampleVirtualHubsClient_BeginGetInboundRoutes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubsClient().BeginGetInboundRoutes(ctx, "rg1", "virtualHub1", armnetwork.GetInboundRoutesParameters{
		ConnectionType: to.Ptr("ExpressRouteConnection"),
		ResourceURI:    to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGw1/expressRouteConnections/exrConn1"),
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
	// res.EffectiveRouteMapRouteList = armnetwork.EffectiveRouteMapRouteList{
	// 	Value: []*armnetwork.EffectiveRouteMapRoute{
	// 		{
	// 			AsPath: to.Ptr("65520-65520"),
	// 			BgpCommunities: to.Ptr("4293853166,12076,51004"),
	// 			Prefix: to.Ptr("192.168.50.0/24"),
	// 		},
	// 		{
	// 			AsPath: to.Ptr("65520-65520-12076-12076"),
	// 			BgpCommunities: to.Ptr("4293787629,12076,51027,4293734188"),
	// 			Prefix: to.Ptr("10.200.0.0/16"),
	// 	}},
	// }
}
