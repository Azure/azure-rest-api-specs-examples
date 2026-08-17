package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/ExpressRouteLagCreate.json
func ExampleExpressRouteLagsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteLagsClient().BeginCreateOrUpdate(ctx, "rg1", "lagName", armnetwork.ExpressRouteLag{
		Location: to.Ptr("eastus2euap"),
		Properties: &armnetwork.ExpressRouteLagPropertiesFormat{
			PeeringLocation:            to.Ptr("peeringLocationName"),
			BandwidthInGbps:            to.Ptr[int32](100),
			Encapsulation:              to.Ptr(armnetwork.ExpressRouteLagEncapsulationDot1Q),
			BillingType:                to.Ptr(armnetwork.ExpressRouteLagBillingTypeMeteredData),
			NumberOfPorts:              to.Ptr[int32](4),
			MinimumActivePortsRequired: to.Ptr[int32](4),
			LacpTimer:                  to.Ptr(armnetwork.ExpressRouteLagLacpTimerFast),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ExpressRouteLagsClientCreateOrUpdateResponse{
	// 	ExpressRouteLag: armnetwork.ExpressRouteLag{
	// 		Name: to.Ptr("lagName"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName"),
	// 		Etag: to.Ptr("W/\"c78dc2d8-a570-4b71-8c00-352d8baba55c\""),
	// 		Type: to.Ptr("Microsoft.Network/expressRouteLags"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Properties: &armnetwork.ExpressRouteLagPropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 			ResourceGUID: to.Ptr("35f2e40c-e3a8-4637-a39d-38065f614fbf"),
	// 			Links: []*armnetwork.ExpressRouteLagLink{
	// 				{
	// 					Name: to.Ptr("link1"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link1"),
	// 					Etag: to.Ptr("W/\"c78dc2d8-a570-4b71-8c00-352d8baba55c\""),
	// 					Properties: &armnetwork.ExpressRouteLagLinkPropertiesFormat{
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 						RouterName: to.Ptr("router1"),
	// 						InterfaceName: to.Ptr("ae"),
	// 						AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 						MacSecConfig: &armnetwork.ExpressRouteLinkMacSecConfig{
	// 							Cipher: to.Ptr(armnetwork.ExpressRouteLinkMacSecCipherGCMAES128),
	// 							SciState: to.Ptr(armnetwork.ExpressRouteLinkMacSecSciStateDisabled),
	// 						},
	// 						Members: []*armnetwork.ExpressRouteLagMember{
	// 							{
	// 								Name: to.Ptr("member1"),
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link1/members/member1"),
	// 								Etag: to.Ptr("W/\"c78dc2d8-a570-4b71-8c00-352d8baba55c\""),
	// 								Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 									InterfaceName: to.Ptr("Ethernet0/0"),
	// 									PatchPanelID: to.Ptr("patchPanelId1"),
	// 									RackID: to.Ptr("rackId1"),
	// 									ColoLocation: to.Ptr(""),
	// 									ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 									AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 								},
	// 								Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
	// 							},
	// 							{
	// 								Name: to.Ptr("member2"),
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link1/members/member2"),
	// 								Etag: to.Ptr("W/\"c78dc2d8-a570-4b71-8c00-352d8baba55c\""),
	// 								Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 									InterfaceName: to.Ptr("Ethernet0/0"),
	// 									PatchPanelID: to.Ptr("patchPanelId2"),
	// 									RackID: to.Ptr("rackId2"),
	// 									ColoLocation: to.Ptr(""),
	// 									ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 									AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 								},
	// 								Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
	// 							},
	// 							{
	// 								Name: to.Ptr("member3"),
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link1/members/member3"),
	// 								Etag: to.Ptr("W/\"c78dc2d8-a570-4b71-8c00-352d8baba55c\""),
	// 								Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 									InterfaceName: to.Ptr("Ethernet0/0"),
	// 									PatchPanelID: to.Ptr("patchPanelId3"),
	// 									RackID: to.Ptr("rackId3"),
	// 									ColoLocation: to.Ptr(""),
	// 									ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 									AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 								},
	// 								Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
	// 							},
	// 							{
	// 								Name: to.Ptr("member4"),
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link1/members/member4"),
	// 								Etag: to.Ptr("W/\"c78dc2d8-a570-4b71-8c00-352d8baba55c\""),
	// 								Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 									InterfaceName: to.Ptr("Ethernet0/0"),
	// 									PatchPanelID: to.Ptr("patchPanelId4"),
	// 									RackID: to.Ptr("rackId4"),
	// 									ColoLocation: to.Ptr(""),
	// 									ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 									AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 								},
	// 								Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
	// 							},
	// 						},
	// 					},
	// 					Type: to.Ptr("Microsoft.Network/expressRouteLags/links"),
	// 				},
	// 				{
	// 					Name: to.Ptr("link2"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link2"),
	// 					Etag: to.Ptr("W/\"c78dc2d8-a570-4b71-8c00-352d8baba55c\""),
	// 					Properties: &armnetwork.ExpressRouteLagLinkPropertiesFormat{
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 						RouterName: to.Ptr("router2"),
	// 						InterfaceName: to.Ptr("ae"),
	// 						AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 						MacSecConfig: &armnetwork.ExpressRouteLinkMacSecConfig{
	// 							Cipher: to.Ptr(armnetwork.ExpressRouteLinkMacSecCipherGCMAES128),
	// 							SciState: to.Ptr(armnetwork.ExpressRouteLinkMacSecSciStateDisabled),
	// 						},
	// 						Members: []*armnetwork.ExpressRouteLagMember{
	// 							{
	// 								Name: to.Ptr("member1"),
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link2/members/member1"),
	// 								Etag: to.Ptr("W/\"c78dc2d8-a570-4b71-8c00-352d8baba55c\""),
	// 								Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 									InterfaceName: to.Ptr("Ethernet0/0"),
	// 									PatchPanelID: to.Ptr("patchPanelId1"),
	// 									RackID: to.Ptr("rackId1"),
	// 									ColoLocation: to.Ptr(""),
	// 									ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 									AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 								},
	// 								Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
	// 							},
	// 							{
	// 								Name: to.Ptr("member2"),
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link2/members/member2"),
	// 								Etag: to.Ptr("W/\"c78dc2d8-a570-4b71-8c00-352d8baba55c\""),
	// 								Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 									InterfaceName: to.Ptr("Ethernet0/0"),
	// 									PatchPanelID: to.Ptr("patchPanelId2"),
	// 									RackID: to.Ptr("rackId2"),
	// 									ColoLocation: to.Ptr(""),
	// 									ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 									AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 								},
	// 								Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
	// 							},
	// 							{
	// 								Name: to.Ptr("member3"),
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link2/members/member3"),
	// 								Etag: to.Ptr("W/\"c78dc2d8-a570-4b71-8c00-352d8baba55c\""),
	// 								Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 									InterfaceName: to.Ptr("Ethernet0/0"),
	// 									PatchPanelID: to.Ptr("patchPanelId3"),
	// 									RackID: to.Ptr("rackId3"),
	// 									ColoLocation: to.Ptr(""),
	// 									ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 									AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 								},
	// 								Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
	// 							},
	// 							{
	// 								Name: to.Ptr("member4"),
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link2/members/member4"),
	// 								Etag: to.Ptr("W/\"c78dc2d8-a570-4b71-8c00-352d8baba55c\""),
	// 								Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 									InterfaceName: to.Ptr("Ethernet0/0"),
	// 									PatchPanelID: to.Ptr("patchPanelId4"),
	// 									RackID: to.Ptr("rackId4"),
	// 									ColoLocation: to.Ptr(""),
	// 									ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 									AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 								},
	// 								Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
	// 							},
	// 						},
	// 					},
	// 					Type: to.Ptr("Microsoft.Network/expressRouteLags/links"),
	// 				},
	// 			},
	// 			PeeringLocation: to.Ptr("peeringLocationName"),
	// 			BandwidthInGbps: to.Ptr[int32](100),
	// 			Encapsulation: to.Ptr(armnetwork.ExpressRouteLagEncapsulationDot1Q),
	// 			ProvisionedBandwidthInGbps: to.Ptr[float64](300),
	// 			Mtu: to.Ptr("4088"),
	// 			AllocationDate: to.Ptr("4/13/2026 3:29:53 PM"),
	// 			EtherType: to.Ptr("0x8100"),
	// 			BillingType: to.Ptr(armnetwork.ExpressRouteLagBillingTypeMeteredData),
	// 			NumberOfPorts: to.Ptr[int32](4),
	// 			MinimumActivePortsRequired: to.Ptr[int32](4),
	// 			LacpTimer: to.Ptr(armnetwork.ExpressRouteLagLacpTimerFast),
	// 		},
	// 	},
	// }
}
