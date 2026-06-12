package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteLinkGet.json
func ExampleExpressRouteLinksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteLinksClient().Get(ctx, "rg1", "portName", "linkName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ExpressRouteLinksClientGetResponse{
	// 	ExpressRouteLink: armnetwork.ExpressRouteLink{
	// 		Name: to.Ptr("linkName"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/linkName"),
	// 		Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
	// 			AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 			ColoLocation: to.Ptr("coloLocationName"),
	// 			ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 			InterfaceName: to.Ptr("Ethernet 0/0"),
	// 			PatchPanelID: to.Ptr("patchPanelId1"),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			RackID: to.Ptr("rackId1"),
	// 			RouterName: to.Ptr("router1"),
	// 		},
	// 	},
	// }
}
