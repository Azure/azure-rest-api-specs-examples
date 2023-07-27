package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/VirtualNetworkGatewaySetVpnClientIpsecParameters.json
func ExampleVirtualNetworkGatewaysClient_BeginSetVpnclientIPSecParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkGatewaysClient().BeginSetVpnclientIPSecParameters(ctx, "rg1", "vpngw", armnetwork.VPNClientIPsecParameters{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VPNClientIPsecParameters = armnetwork.VPNClientIPsecParameters{
	// 	DhGroup: to.Ptr(armnetwork.DhGroupDHGroup2),
	// 	IkeEncryption: to.Ptr(armnetwork.IkeEncryptionAES256),
	// 	IkeIntegrity: to.Ptr(armnetwork.IkeIntegritySHA384),
	// 	IPSecEncryption: to.Ptr(armnetwork.IPSecEncryptionAES256),
	// 	IPSecIntegrity: to.Ptr(armnetwork.IPSecIntegritySHA256),
	// 	PfsGroup: to.Ptr(armnetwork.PfsGroupPFS2),
	// 	SaDataSizeKilobytes: to.Ptr[int32](429497),
	// 	SaLifeTimeSeconds: to.Ptr[int32](86473),
	// }
}
