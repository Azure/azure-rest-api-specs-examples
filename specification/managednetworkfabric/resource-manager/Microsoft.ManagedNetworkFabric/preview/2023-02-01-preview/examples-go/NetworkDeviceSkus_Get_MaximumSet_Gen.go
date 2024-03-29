package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkDeviceSkus_Get_MaximumSet_Gen.json
func ExampleNetworkDeviceSKUsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkDeviceSKUsClient().Get(ctx, "DefaultSku", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkDeviceSKU = armmanagednetworkfabric.NetworkDeviceSKU{
	// 	ID: to.Ptr("resourceId"),
	// 	Properties: &armmanagednetworkfabric.NetworkDeviceSKUProperties{
	// 		Interfaces: []*armmanagednetworkfabric.DeviceInterfaceProperties{
	// 			{
	// 				Identifier: to.Ptr("1"),
	// 				InterfaceType: to.Ptr("Ethernet"),
	// 				SupportedConnectorTypes: []*armmanagednetworkfabric.SupportedConnectorProperties{
	// 					{
	// 						ConnectorType: to.Ptr("10GBASE-LRL"),
	// 						MaxSpeedInMbps: to.Ptr[int32](10240),
	// 				}},
	// 			},
	// 			{
	// 				Identifier: to.Ptr("2"),
	// 				InterfaceType: to.Ptr("Ethernet"),
	// 				SupportedConnectorTypes: []*armmanagednetworkfabric.SupportedConnectorProperties{
	// 					{
	// 						ConnectorType: to.Ptr("10GBASE-LRL"),
	// 						MaxSpeedInMbps: to.Ptr[int32](10240),
	// 				}},
	// 		}},
	// 		Limits: &armmanagednetworkfabric.DeviceLimits{
	// 			MaxBidirectionalForwardingDetectionPeers: to.Ptr[int32](0),
	// 			MaxBorderGatewayProtocolPeers: to.Ptr[int32](0),
	// 			MaxSubInterfaces: to.Ptr[int32](0),
	// 			MaxTunnelInterfaces: to.Ptr[int32](0),
	// 			MaxVirtualRouterFunctions: to.Ptr[int32](0),
	// 			PhysicalInterfaceCount: to.Ptr[int32](2),
	// 		},
	// 		Manufacturer: to.Ptr("Arista"),
	// 		Model: to.Ptr("DCS-7280DR3A-36-F"),
	// 		SupportedRoleTypes: []*armmanagednetworkfabric.NetworkDeviceRoleName{
	// 			to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleNameCE),
	// 			to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleNameToR)},
	// 			SupportedVersions: []*armmanagednetworkfabric.SupportedVersionProperties{
	// 				{
	// 					IsCurrent: to.Ptr(armmanagednetworkfabric.IsCurrentVersionTrue),
	// 					IsTest: to.Ptr(armmanagednetworkfabric.IsTestVersionFalse),
	// 					VendorFirmwareVersion: to.Ptr("1.1.2"),
	// 					VendorOsVersion: to.Ptr("4.22.2"),
	// 					Version: to.Ptr("4.22.2,1.1.2"),
	// 			}},
	// 		},
	// 	}
}
