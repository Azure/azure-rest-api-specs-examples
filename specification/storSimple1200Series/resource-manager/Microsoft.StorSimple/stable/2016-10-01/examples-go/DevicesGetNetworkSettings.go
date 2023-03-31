package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/DevicesGetNetworkSettings.json
func ExampleDevicesClient_GetNetworkSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().GetNetworkSettings(ctx, "HSDK-T4ZA3EAJFR", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSettings = armstorsimple1200series.NetworkSettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/networkSettings"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-T4ZA3EAJFR/networkSettings/default"),
	// 	Properties: &armstorsimple1200series.NetworkSettingsProperties{
	// 		NodeNetworks: []*armstorsimple1200series.NodeNetwork{
	// 			{
	// 				NetworkAdapters: []*armstorsimple1200series.NetworkAdapter{
	// 					{
	// 						DhcpStatus: to.Ptr(armstorsimple1200series.DhcpStatusEnabled),
	// 						IPv4Info: &armstorsimple1200series.IPConfig{
	// 							Gateway: to.Ptr("10.60.145.1"),
	// 							IPAddress: to.Ptr("10.60.145.5"),
	// 							PrefixLength: to.Ptr[int32](26),
	// 						},
	// 						LinkSpeed: to.Ptr[int64](868993459200),
	// 						NetworkAdapterName: to.Ptr("Ethernet"),
	// 				}},
	// 		}},
	// 		PrimaryDNSServer: to.Ptr("10.50.50.50"),
	// 		SecondaryDNSServer: to.Ptr("10.50.10.50"),
	// 	},
	// }
}
