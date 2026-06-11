package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteLinkList.json
func ExampleExpressRouteLinksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRouteLinksClient().NewListPager("rg1", "portName", nil)
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
		// page = armnetwork.ExpressRouteLinksClientListResponse{
		// 	ExpressRouteLinkListResult: armnetwork.ExpressRouteLinkListResult{
		// 		Value: []*armnetwork.ExpressRouteLink{
		// 			{
		// 				Name: to.Ptr("link1"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link1"),
		// 				Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
		// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
		// 					ColoLocation: to.Ptr("coloLocation1"),
		// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 					InterfaceName: to.Ptr("Ethernet 0/0"),
		// 					PatchPanelID: to.Ptr("patchPanelId1"),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RackID: to.Ptr("rackId1"),
		// 					RouterName: to.Ptr("router1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("link2"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link2"),
		// 				Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
		// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
		// 					ColoLocation: to.Ptr("coloLocation2"),
		// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 					InterfaceName: to.Ptr("Ethernet 0/0"),
		// 					PatchPanelID: to.Ptr("patchPanelId2"),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RackID: to.Ptr("rackId2"),
		// 					RouterName: to.Ptr("router2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
