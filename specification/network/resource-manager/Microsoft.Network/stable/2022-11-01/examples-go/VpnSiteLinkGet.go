package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/VpnSiteLinkGet.json
func ExampleVPNSiteLinksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVPNSiteLinksClient().Get(ctx, "rg1", "vpnSite1", "vpnSiteLink1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VPNSiteLink = armnetwork.VPNSiteLink{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/vpnSiteLink1"),
	// 	Name: to.Ptr("vpnSiteLink1"),
	// 	Type: to.Ptr("Microsoft.Network/vpnSites/vpnSiteLinks"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armnetwork.VPNSiteLinkProperties{
	// 		BgpProperties: &armnetwork.VPNLinkBgpSettings{
	// 			Asn: to.Ptr[int64](1234),
	// 			BgpPeeringAddress: to.Ptr("192.168.0.0"),
	// 		},
	// 		IPAddress: to.Ptr("50.50.50.56"),
	// 		LinkProperties: &armnetwork.VPNLinkProviderProperties{
	// 			LinkSpeedInMbps: to.Ptr[int32](0),
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
