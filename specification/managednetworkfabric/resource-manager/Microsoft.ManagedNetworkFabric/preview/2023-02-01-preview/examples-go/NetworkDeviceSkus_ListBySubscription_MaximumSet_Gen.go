package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkDeviceSkus_ListBySubscription_MaximumSet_Gen.json
func ExampleNetworkDeviceSKUsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkDeviceSKUsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.NetworkDeviceSKUsListResult = armmanagednetworkfabric.NetworkDeviceSKUsListResult{
		// 	Value: []*armmanagednetworkfabric.NetworkDeviceSKU{
		// 		{
		// 			Name: to.Ptr("A-DCS-7280DR3A-36-F"),
		// 			Type: to.Ptr("NetworkDeviceSkus"),
		// 			ID: to.Ptr("A-DCS-7280DR3A-36-F"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-04T13:05:13.867Z"); return t}()),
		// 				CreatedBy: to.Ptr("email@address.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-04T13:05:13.867Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("email@address.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 			},
		// 			Properties: &armmanagednetworkfabric.NetworkDeviceSKUProperties{
		// 				Interfaces: []*armmanagednetworkfabric.DeviceInterfaceProperties{
		// 					{
		// 						Identifier: to.Ptr("1"),
		// 						InterfaceType: to.Ptr("Ethernet"),
		// 						SupportedConnectorTypes: []*armmanagednetworkfabric.SupportedConnectorProperties{
		// 							{
		// 								ConnectorType: to.Ptr("10GBASE-LRL"),
		// 								MaxSpeedInMbps: to.Ptr[int32](10240),
		// 						}},
		// 					},
		// 					{
		// 						Identifier: to.Ptr("2"),
		// 						InterfaceType: to.Ptr("Ethernet"),
		// 						SupportedConnectorTypes: []*armmanagednetworkfabric.SupportedConnectorProperties{
		// 							{
		// 								ConnectorType: to.Ptr("10GBASE-LRL"),
		// 								MaxSpeedInMbps: to.Ptr[int32](10240),
		// 						}},
		// 				}},
		// 				Limits: &armmanagednetworkfabric.DeviceLimits{
		// 					MaxBidirectionalForwardingDetectionPeers: to.Ptr[int32](0),
		// 					MaxBorderGatewayProtocolPeers: to.Ptr[int32](0),
		// 					MaxSubInterfaces: to.Ptr[int32](0),
		// 					MaxTunnelInterfaces: to.Ptr[int32](0),
		// 					MaxVirtualRouterFunctions: to.Ptr[int32](0),
		// 					PhysicalInterfaceCount: to.Ptr[int32](2),
		// 				},
		// 				Manufacturer: to.Ptr("Arista"),
		// 				Model: to.Ptr("DCS-7280DR3A-36-F"),
		// 				SupportedRoleTypes: []*armmanagednetworkfabric.NetworkDeviceRoleName{
		// 					to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleNameCE),
		// 					to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleNameToR)},
		// 					SupportedVersions: []*armmanagednetworkfabric.SupportedVersionProperties{
		// 						{
		// 							IsCurrent: to.Ptr(armmanagednetworkfabric.IsCurrentVersionTrue),
		// 							IsTest: to.Ptr(armmanagednetworkfabric.IsTestVersionFalse),
		// 							VendorFirmwareVersion: to.Ptr("1.1.2"),
		// 							VendorOsVersion: to.Ptr("4.22.2"),
		// 							Version: to.Ptr("4.22.2,1.1.2"),
		// 					}},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("A-DCS-7280DR3K-24-F"),
		// 				Type: to.Ptr("NetworkDeviceSkus"),
		// 				ID: to.Ptr("A-DCS-7280DR3K-24-F"),
		// 				SystemData: &armmanagednetworkfabric.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-04T13:05:13.867Z"); return t}()),
		// 					CreatedBy: to.Ptr("email@address.com"),
		// 					CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-04T13:05:13.867Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("email@address.com"),
		// 					LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 				},
		// 				Properties: &armmanagednetworkfabric.NetworkDeviceSKUProperties{
		// 					Interfaces: []*armmanagednetworkfabric.DeviceInterfaceProperties{
		// 						{
		// 							Identifier: to.Ptr(""),
		// 							InterfaceType: to.Ptr(""),
		// 							SupportedConnectorTypes: []*armmanagednetworkfabric.SupportedConnectorProperties{
		// 								{
		// 									ConnectorType: to.Ptr(""),
		// 									MaxSpeedInMbps: to.Ptr[int32](0),
		// 							}},
		// 					}},
		// 					Limits: &armmanagednetworkfabric.DeviceLimits{
		// 						MaxBidirectionalForwardingDetectionPeers: to.Ptr[int32](0),
		// 						MaxBorderGatewayProtocolPeers: to.Ptr[int32](0),
		// 						MaxSubInterfaces: to.Ptr[int32](0),
		// 						MaxTunnelInterfaces: to.Ptr[int32](0),
		// 						MaxVirtualRouterFunctions: to.Ptr[int32](0),
		// 						PhysicalInterfaceCount: to.Ptr[int32](0),
		// 					},
		// 					Manufacturer: to.Ptr("Arista"),
		// 					Model: to.Ptr("DCS-7280DR3K-24-F"),
		// 					SupportedRoleTypes: []*armmanagednetworkfabric.NetworkDeviceRoleName{
		// 						to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleNameNPB)},
		// 						SupportedVersions: []*armmanagednetworkfabric.SupportedVersionProperties{
		// 							{
		// 								IsCurrent: to.Ptr(armmanagednetworkfabric.IsCurrentVersionTrue),
		// 								IsTest: to.Ptr(armmanagednetworkfabric.IsTestVersionFalse),
		// 								VendorFirmwareVersion: to.Ptr(""),
		// 								VendorOsVersion: to.Ptr(""),
		// 								Version: to.Ptr(""),
		// 						}},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("O-OM2248-10G-US"),
		// 					Type: to.Ptr("NetworkDeviceSkus"),
		// 					ID: to.Ptr("O-OM2248-10G-US"),
		// 					SystemData: &armmanagednetworkfabric.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-04T13:05:13.867Z"); return t}()),
		// 						CreatedBy: to.Ptr("email@address.com"),
		// 						CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-04T13:05:13.867Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("email@address.com"),
		// 						LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 					},
		// 					Properties: &armmanagednetworkfabric.NetworkDeviceSKUProperties{
		// 						Interfaces: []*armmanagednetworkfabric.DeviceInterfaceProperties{
		// 							{
		// 								Identifier: to.Ptr(""),
		// 								InterfaceType: to.Ptr(""),
		// 								SupportedConnectorTypes: []*armmanagednetworkfabric.SupportedConnectorProperties{
		// 									{
		// 										ConnectorType: to.Ptr(""),
		// 										MaxSpeedInMbps: to.Ptr[int32](0),
		// 								}},
		// 						}},
		// 						Limits: &armmanagednetworkfabric.DeviceLimits{
		// 							MaxBidirectionalForwardingDetectionPeers: to.Ptr[int32](0),
		// 							MaxBorderGatewayProtocolPeers: to.Ptr[int32](0),
		// 							MaxSubInterfaces: to.Ptr[int32](0),
		// 							MaxTunnelInterfaces: to.Ptr[int32](0),
		// 							MaxVirtualRouterFunctions: to.Ptr[int32](0),
		// 							PhysicalInterfaceCount: to.Ptr[int32](0),
		// 						},
		// 						Manufacturer: to.Ptr("OpenGear"),
		// 						Model: to.Ptr("OM2248-10G-US"),
		// 						SupportedRoleTypes: []*armmanagednetworkfabric.NetworkDeviceRoleName{
		// 							to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleNameTS)},
		// 							SupportedVersions: []*armmanagednetworkfabric.SupportedVersionProperties{
		// 								{
		// 									IsCurrent: to.Ptr(armmanagednetworkfabric.IsCurrentVersionTrue),
		// 									IsTest: to.Ptr(armmanagednetworkfabric.IsTestVersionFalse),
		// 									VendorFirmwareVersion: to.Ptr(""),
		// 									VendorOsVersion: to.Ptr(""),
		// 									Version: to.Ptr(""),
		// 							}},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("A-DCS-7010TX-48-F"),
		// 						Type: to.Ptr("NetworkDeviceSkus"),
		// 						ID: to.Ptr("A-DCS-7010TX-48-F"),
		// 						SystemData: &armmanagednetworkfabric.SystemData{
		// 							CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-04T13:05:13.867Z"); return t}()),
		// 							CreatedBy: to.Ptr("email@address.com"),
		// 							CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 							LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-04T13:05:13.867Z"); return t}()),
		// 							LastModifiedBy: to.Ptr("email@address.com"),
		// 							LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 						},
		// 						Properties: &armmanagednetworkfabric.NetworkDeviceSKUProperties{
		// 							Interfaces: []*armmanagednetworkfabric.DeviceInterfaceProperties{
		// 								{
		// 									Identifier: to.Ptr(""),
		// 									InterfaceType: to.Ptr(""),
		// 									SupportedConnectorTypes: []*armmanagednetworkfabric.SupportedConnectorProperties{
		// 										{
		// 											ConnectorType: to.Ptr(""),
		// 											MaxSpeedInMbps: to.Ptr[int32](0),
		// 									}},
		// 							}},
		// 							Limits: &armmanagednetworkfabric.DeviceLimits{
		// 								MaxBidirectionalForwardingDetectionPeers: to.Ptr[int32](0),
		// 								MaxBorderGatewayProtocolPeers: to.Ptr[int32](0),
		// 								MaxSubInterfaces: to.Ptr[int32](0),
		// 								MaxTunnelInterfaces: to.Ptr[int32](0),
		// 								MaxVirtualRouterFunctions: to.Ptr[int32](0),
		// 								PhysicalInterfaceCount: to.Ptr[int32](0),
		// 							},
		// 							Manufacturer: to.Ptr("Arista"),
		// 							Model: to.Ptr("DCS-7010TX-48-F"),
		// 							SupportedRoleTypes: []*armmanagednetworkfabric.NetworkDeviceRoleName{
		// 								to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleNameManagement)},
		// 								SupportedVersions: []*armmanagednetworkfabric.SupportedVersionProperties{
		// 									{
		// 										IsCurrent: to.Ptr(armmanagednetworkfabric.IsCurrentVersionTrue),
		// 										IsTest: to.Ptr(armmanagednetworkfabric.IsTestVersionFalse),
		// 										VendorFirmwareVersion: to.Ptr(""),
		// 										VendorOsVersion: to.Ptr(""),
		// 										Version: to.Ptr(""),
		// 								}},
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("A-DCS-7280DR3-24-F"),
		// 							Type: to.Ptr("NetworkDeviceSkus"),
		// 							ID: to.Ptr("A-DCS-7280DR3-24-F"),
		// 							SystemData: &armmanagednetworkfabric.SystemData{
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-04T13:05:13.867Z"); return t}()),
		// 								CreatedBy: to.Ptr("email@address.com"),
		// 								CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-04T13:05:13.867Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("email@address.com"),
		// 								LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 							},
		// 							Properties: &armmanagednetworkfabric.NetworkDeviceSKUProperties{
		// 								Interfaces: []*armmanagednetworkfabric.DeviceInterfaceProperties{
		// 									{
		// 										Identifier: to.Ptr(""),
		// 										InterfaceType: to.Ptr(""),
		// 										SupportedConnectorTypes: []*armmanagednetworkfabric.SupportedConnectorProperties{
		// 											{
		// 												ConnectorType: to.Ptr(""),
		// 												MaxSpeedInMbps: to.Ptr[int32](0),
		// 										}},
		// 								}},
		// 								Limits: &armmanagednetworkfabric.DeviceLimits{
		// 									MaxBidirectionalForwardingDetectionPeers: to.Ptr[int32](0),
		// 									MaxBorderGatewayProtocolPeers: to.Ptr[int32](0),
		// 									MaxSubInterfaces: to.Ptr[int32](0),
		// 									MaxTunnelInterfaces: to.Ptr[int32](0),
		// 									MaxVirtualRouterFunctions: to.Ptr[int32](0),
		// 									PhysicalInterfaceCount: to.Ptr[int32](0),
		// 								},
		// 								Manufacturer: to.Ptr("Arista"),
		// 								Model: to.Ptr("DCS-7280DR3-24-F"),
		// 								SupportedRoleTypes: []*armmanagednetworkfabric.NetworkDeviceRoleName{
		// 									to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleNameCE),
		// 									to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleNameToR)},
		// 									SupportedVersions: []*armmanagednetworkfabric.SupportedVersionProperties{
		// 										{
		// 											IsCurrent: to.Ptr(armmanagednetworkfabric.IsCurrentVersionTrue),
		// 											IsTest: to.Ptr(armmanagednetworkfabric.IsTestVersionFalse),
		// 											VendorFirmwareVersion: to.Ptr(""),
		// 											VendorOsVersion: to.Ptr(""),
		// 											Version: to.Ptr(""),
		// 									}},
		// 								},
		// 						}},
		// 					}
	}
}
