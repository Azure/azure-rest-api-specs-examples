package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkGatewaySupportedVpnDevice.json
func ExampleVirtualNetworkGatewaysClient_SupportedVPNDevices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkGatewaysClient().SupportedVPNDevices(ctx, "rg1", "vpngw", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Value = "<?xml version=\"1.0\" encoding=\"utf-8\"?><RpVpnDeviceList version=\"1.0\"><Vendor name=\"Cisco\"><DeviceFamily name=\"ISR\"><FirmwareVersion name=\"IOS 15.1 (Preview)\" /></DeviceFamily><DeviceFamily name=\"ASA (Adaptive Security Appliance)\"><FirmwareVersion name=\"ASA_9.8.x_ActivePassive_NoBGP_NoCustomPolicies\"/><FirmwareVersion name=\"ASA_9.8(x)_CustomPolicies\" /></DeviceFamily></Vendor><Vendor name=\"Juniper\"><DeviceFamily name=\"Juniper_SRX_GA\"><FirmwareVersion name=\"Juniper_SRX_12.x_GA\" /></DeviceFamily><DeviceFamily name=\"Juniper_SSG_GA\"><FirmwareVersion name=\"Juniper_SSG_ScreenOS-6.2.x_GA\"/></DeviceFamily><DeviceFamily name=\"Juniper_JSeries_GA\"><FirmwareVersion name=\"Juniper_JSeries_JunOS12.x._GA\" /></DeviceFamily></Vendor><Vendor name=\"Ubiquiti\"><DeviceFamily name=\"EdgeRouter\"><FirmwareVersion name=\"Ubiquiti_EdgeOS_1.10.x-RouteBased_VTI\" /><FirmwareVersionname=\"Ubiquiti_EdgeOS_1.10.x-RouteBased_BGP\" /></DeviceFamily></Vendor></RpVpnDeviceList>"
}
