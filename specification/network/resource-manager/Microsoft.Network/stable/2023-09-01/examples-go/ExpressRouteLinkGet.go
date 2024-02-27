package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/ExpressRouteLinkGet.json
func ExampleExpressRouteLinksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.ExpressRouteLink = armnetwork.ExpressRouteLink{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/linkName"),
	// 	Name: to.Ptr("linkName"),
	// 	Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
	// 		AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 		ColoLocation: to.Ptr("coloLocationName"),
	// 		ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 		InterfaceName: to.Ptr("Ethernet 0/0"),
	// 		PatchPanelID: to.Ptr("patchPanelId1"),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RackID: to.Ptr("rackId1"),
	// 		RouterName: to.Ptr("router1"),
	// 	},
	// }
}
