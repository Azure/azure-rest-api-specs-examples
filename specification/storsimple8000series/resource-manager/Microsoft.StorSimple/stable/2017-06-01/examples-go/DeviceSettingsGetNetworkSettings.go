package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsGetNetworkSettings.json
func ExampleDeviceSettingsClient_GetNetworkSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeviceSettingsClient().GetNetworkSettings(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSettings = armstorsimple8000series.NetworkSettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/networkSettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/networkSettings/default"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.NetworkSettingsProperties{
	// 		DNSSettings: &armstorsimple8000series.DNSSettings{
	// 			PrimaryDNSServer: to.Ptr("10.171.65.60"),
	// 			SecondaryDNSServers: []*string{
	// 				to.Ptr("8.8.8.8")},
	// 				SecondaryIPv6DNSServers: []*string{
	// 				},
	// 			},
	// 			NetworkAdapters: &armstorsimple8000series.NetworkAdapterList{
	// 				Value: []*armstorsimple8000series.NetworkAdapters{
	// 					{
	// 						InterfaceID: to.Ptr(armstorsimple8000series.NetInterfaceIDData0),
	// 						IsDefault: to.Ptr(true),
	// 						IscsiAndCloudStatus: to.Ptr(armstorsimple8000series.ISCSIAndCloudStatusIscsiAndCloudEnabled),
	// 						Mode: to.Ptr(armstorsimple8000series.NetworkModeIPV4),
	// 						NetInterfaceStatus: to.Ptr(armstorsimple8000series.NetInterfaceStatusEnabled),
	// 						NicIPv4Settings: &armstorsimple8000series.NicIPv4{
	// 							Controller0IPv4Address: to.Ptr("10.168.241.143"),
	// 							Controller1IPv4Address: to.Ptr("10.168.241.121"),
	// 							IPv4Address: to.Ptr("10.168.241.187"),
	// 							IPv4Gateway: to.Ptr("10.168.241.1"),
	// 							IPv4Netmask: to.Ptr("255.255.252.0"),
	// 						},
	// 						NicIPv6Settings: &armstorsimple8000series.NicIPv6{
	// 							Controller0IPv6Address: to.Ptr(""),
	// 							Controller1IPv6Address: to.Ptr(""),
	// 							IPv6Address: to.Ptr(""),
	// 							IPv6Gateway: to.Ptr(""),
	// 							IPv6Prefix: to.Ptr(""),
	// 						},
	// 						Speed: to.Ptr[int64](1000000000),
	// 				}},
	// 			},
	// 			WebproxySettings: &armstorsimple8000series.WebproxySettings{
	// 				Authentication: to.Ptr(armstorsimple8000series.AuthenticationTypeNone),
	// 				ConnectionURI: to.Ptr(""),
	// 				Username: to.Ptr(""),
	// 			},
	// 		},
	// 	}
}
