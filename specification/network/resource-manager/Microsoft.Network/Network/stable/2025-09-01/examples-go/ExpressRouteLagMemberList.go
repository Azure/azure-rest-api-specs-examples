package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/ExpressRouteLagMemberList.json
func ExampleExpressRouteLagsClient_NewMembersListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRouteLagsClient().NewMembersListPager("rg1", "lagName", "linkName", nil)
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
		// page = armnetwork.ExpressRouteLagsClientMembersListResponse{
		// 	ExpressRouteLagMemberListResult: armnetwork.ExpressRouteLagMemberListResult{
		// 		Value: []*armnetwork.ExpressRouteLagMember{
		// 			{
		// 				Name: to.Ptr("member1"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/linkName/members/member1"),
		// 				Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 				Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					InterfaceName: to.Ptr("Ethernet0/0"),
		// 					PatchPanelID: to.Ptr("patchPanelId1"),
		// 					RackID: to.Ptr("rackId1"),
		// 					ColoLocation: to.Ptr(""),
		// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 				},
		// 				Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 			},
		// 			{
		// 				Name: to.Ptr("member2"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/linkName/members/member2"),
		// 				Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 				Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					InterfaceName: to.Ptr("Ethernet0/0"),
		// 					PatchPanelID: to.Ptr("patchPanelId2"),
		// 					RackID: to.Ptr("rackId2"),
		// 					ColoLocation: to.Ptr(""),
		// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 				},
		// 				Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 			},
		// 			{
		// 				Name: to.Ptr("member3"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/linkName/members/member3"),
		// 				Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 				Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					InterfaceName: to.Ptr("Ethernet0/0"),
		// 					PatchPanelID: to.Ptr("patchPanelId3"),
		// 					RackID: to.Ptr("rackId3"),
		// 					ColoLocation: to.Ptr(""),
		// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 				},
		// 				Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 			},
		// 			{
		// 				Name: to.Ptr("member4"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/linkName/members/member4"),
		// 				Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
		// 				Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					InterfaceName: to.Ptr("Ethernet0/0"),
		// 					PatchPanelID: to.Ptr("patchPanelId4"),
		// 					RackID: to.Ptr("rackId4"),
		// 					ColoLocation: to.Ptr(""),
		// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
		// 				},
		// 				Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
