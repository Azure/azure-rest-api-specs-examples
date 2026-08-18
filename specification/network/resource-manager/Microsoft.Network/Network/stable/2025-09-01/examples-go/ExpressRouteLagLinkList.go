package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/ExpressRouteLagLinkList.json
func ExampleExpressRouteLagsClient_NewLinksListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRouteLagsClient().NewLinksListPager("rg1", "lagName", nil)
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
		// page = armnetwork.ExpressRouteLagsClientLinksListResponse{
		// 	ExpressRouteLagLinkListResult: armnetwork.ExpressRouteLagLinkListResult{
		// 		Value: []*armnetwork.ExpressRouteLagLink{
		// 			{
		// 				Name: to.Ptr("link1"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link1"),
		// 				Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 				Properties: &armnetwork.ExpressRouteLagLinkPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RouterName: to.Ptr("router1"),
		// 					InterfaceName: to.Ptr("ae"),
		// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
		// 					MacSecConfig: &armnetwork.ExpressRouteLinkMacSecConfig{
		// 						Cipher: to.Ptr(armnetwork.ExpressRouteLinkMacSecCipherGCMAES128),
		// 						SciState: to.Ptr(armnetwork.ExpressRouteLinkMacSecSciStateDisabled),
		// 					},
		// 					Members: []*armnetwork.ExpressRouteLagMember{
		// 						{
		// 							Name: to.Ptr("member1"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link1/members/member1"),
		// 							Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 							Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								InterfaceName: to.Ptr("Ethernet0/0"),
		// 								PatchPanelID: to.Ptr("patchPanelId1"),
		// 								RackID: to.Ptr("rackId1"),
		// 								ColoLocation: to.Ptr(""),
		// 								ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 								AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 							},
		// 							Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 						},
		// 						{
		// 							Name: to.Ptr("member2"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link1/members/member2"),
		// 							Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 							Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								InterfaceName: to.Ptr("Ethernet0/0"),
		// 								PatchPanelID: to.Ptr("patchPanelId2"),
		// 								RackID: to.Ptr("rackId2"),
		// 								ColoLocation: to.Ptr(""),
		// 								ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 								AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 							},
		// 							Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 						},
		// 						{
		// 							Name: to.Ptr("member3"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link1/members/member3"),
		// 							Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 							Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								InterfaceName: to.Ptr("Ethernet0/0"),
		// 								PatchPanelID: to.Ptr("patchPanelId3"),
		// 								RackID: to.Ptr("rackId3"),
		// 								ColoLocation: to.Ptr(""),
		// 								ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 								AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 							},
		// 							Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 						},
		// 						{
		// 							Name: to.Ptr("member4"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link1/members/member4"),
		// 							Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 							Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								InterfaceName: to.Ptr("Ethernet0/0"),
		// 								PatchPanelID: to.Ptr("patchPanelId4"),
		// 								RackID: to.Ptr("rackId4"),
		// 								ColoLocation: to.Ptr(""),
		// 								ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 								AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 							},
		// 							Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 						},
		// 					},
		// 				},
		// 				Type: to.Ptr("Microsoft.Network/expressRouteLags/links"),
		// 			},
		// 			{
		// 				Name: to.Ptr("link2"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link2"),
		// 				Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 				Properties: &armnetwork.ExpressRouteLagLinkPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RouterName: to.Ptr("router2"),
		// 					InterfaceName: to.Ptr("ae"),
		// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
		// 					MacSecConfig: &armnetwork.ExpressRouteLinkMacSecConfig{
		// 						Cipher: to.Ptr(armnetwork.ExpressRouteLinkMacSecCipherGCMAES128),
		// 						SciState: to.Ptr(armnetwork.ExpressRouteLinkMacSecSciStateDisabled),
		// 					},
		// 					Members: []*armnetwork.ExpressRouteLagMember{
		// 						{
		// 							Name: to.Ptr("member1"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link2/members/member1"),
		// 							Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 							Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								InterfaceName: to.Ptr("Ethernet0/0"),
		// 								PatchPanelID: to.Ptr("patchPanelId1"),
		// 								RackID: to.Ptr("rackId1"),
		// 								ColoLocation: to.Ptr(""),
		// 								ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 								AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 							},
		// 							Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 						},
		// 						{
		// 							Name: to.Ptr("member2"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link2/members/member2"),
		// 							Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 							Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								InterfaceName: to.Ptr("Ethernet0/0"),
		// 								PatchPanelID: to.Ptr("patchPanelId2"),
		// 								RackID: to.Ptr("rackId2"),
		// 								ColoLocation: to.Ptr(""),
		// 								ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 								AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 							},
		// 							Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 						},
		// 						{
		// 							Name: to.Ptr("member3"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link2/members/member3"),
		// 							Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 							Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								InterfaceName: to.Ptr("Ethernet0/0"),
		// 								PatchPanelID: to.Ptr("patchPanelId3"),
		// 								RackID: to.Ptr("rackId3"),
		// 								ColoLocation: to.Ptr(""),
		// 								ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 								AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 							},
		// 							Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 						},
		// 						{
		// 							Name: to.Ptr("member4"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/link2/members/member4"),
		// 							Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 							Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								InterfaceName: to.Ptr("Ethernet0/0"),
		// 								PatchPanelID: to.Ptr("patchPanelId4"),
		// 								RackID: to.Ptr("rackId4"),
		// 								ColoLocation: to.Ptr(""),
		// 								ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 								AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 							},
		// 							Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 						},
		// 					},
		// 				},
		// 				Type: to.Ptr("Microsoft.Network/expressRouteLags/links"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
