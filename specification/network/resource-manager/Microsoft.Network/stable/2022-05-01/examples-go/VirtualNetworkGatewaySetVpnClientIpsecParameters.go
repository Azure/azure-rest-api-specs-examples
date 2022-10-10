package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkGatewaySetVpnClientIpsecParameters.json
func ExampleVirtualNetworkGatewaysClient_BeginSetVpnclientIPSecParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualNetworkGatewaysClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginSetVpnclientIPSecParameters(ctx, "rg1", "vpngw", armnetwork.VPNClientIPsecParameters{
		DhGroup:             to.Ptr(armnetwork.DhGroupDHGroup2),
		IkeEncryption:       to.Ptr(armnetwork.IkeEncryptionAES256),
		IkeIntegrity:        to.Ptr(armnetwork.IkeIntegritySHA384),
		IPSecEncryption:     to.Ptr(armnetwork.IPSecEncryptionAES256),
		IPSecIntegrity:      to.Ptr(armnetwork.IPSecIntegritySHA256),
		PfsGroup:            to.Ptr(armnetwork.PfsGroupPFS2),
		SaDataSizeKilobytes: to.Ptr[int32](429497),
		SaLifeTimeSeconds:   to.Ptr[int32](86473),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
