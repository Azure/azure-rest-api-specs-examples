package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkGatewayGetVpnclientConnectionHealth.json
func ExampleVirtualNetworkGatewaysClient_BeginGetVpnclientConnectionHealth() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkGatewaysClient().BeginGetVpnclientConnectionHealth(ctx, "p2s-vnet-test", "vpnp2sgw", nil)
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
	// res.VPNClientConnectionHealthDetailListResult = armnetwork.VPNClientConnectionHealthDetailListResult{
	// 	Value: []*armnetwork.VPNClientConnectionHealthDetail{
	// 		{
	// 			EgressBytesTransferred: to.Ptr[int64](33420),
	// 			EgressPacketsTransferred: to.Ptr[int64](557),
	// 			IngressBytesTransferred: to.Ptr[int64](33420),
	// 			IngressPacketsTransferred: to.Ptr[int64](557),
	// 			MaxBandwidth: to.Ptr[int64](240000000),
	// 			MaxPacketsPerSecond: to.Ptr[int64](4),
	// 			PrivateIPAddress: to.Ptr("192.168.210.2"),
	// 			PublicIPAddress: to.Ptr("167.220.2.232:45522"),
	// 			VPNConnectionDuration: to.Ptr[int64](900),
	// 			VPNConnectionID: to.Ptr("IKEv2_1e1cfe59-5c7c-4315-a876-b11fbfdfeed4"),
	// 			VPNConnectionTime: to.Ptr("2019-05-02T22:26:22"),
	// 			VPNUserName: to.Ptr("gwp2schildcert"),
	// 		},
	// 		{
	// 			EgressBytesTransferred: to.Ptr[int64](23420),
	// 			EgressPacketsTransferred: to.Ptr[int64](357),
	// 			IngressBytesTransferred: to.Ptr[int64](23420),
	// 			IngressPacketsTransferred: to.Ptr[int64](357),
	// 			MaxBandwidth: to.Ptr[int64](220000000),
	// 			MaxPacketsPerSecond: to.Ptr[int64](4),
	// 			PrivateIPAddress: to.Ptr("192.168.210.1"),
	// 			PublicIPAddress: to.Ptr("167.220.2.232:45213"),
	// 			VPNConnectionDuration: to.Ptr[int64](800),
	// 			VPNConnectionID: to.Ptr("IKEv2_571cfe59-2c7d-1415-e813-c51fbfdfea16"),
	// 			VPNConnectionTime: to.Ptr("2019-05-01T21:06:12"),
	// 			VPNUserName: to.Ptr("gwp2schildcert"),
	// 	}},
	// }
}
