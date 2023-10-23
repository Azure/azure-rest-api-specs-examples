package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1a87e1a5deb3f986ea1474d233d6680f1e19fc1/specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspaceVirtualNetPeeringList.json
func ExampleVNetPeeringClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVNetPeeringClient().NewListByWorkspacePager("rg", "myWorkspace", nil)
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
		// page.VirtualNetworkPeeringList = armdatabricks.VirtualNetworkPeeringList{
		// 	Value: []*armdatabricks.VirtualNetworkPeering{
		// 		{
		// 			Name: to.Ptr("vNetPeeringTest"),
		// 			ID: to.Ptr("/subscriptions/0140911e-1040-48da-8bc9-b99fb3dd88a6/resourceGroups/subramantest/providers/Microsoft.Databricks/workspaces/adbworkspace/virtualNetworkPeerings/vNetPeeringTest"),
		// 			Properties: &armdatabricks.VirtualNetworkPeeringPropertiesFormat{
		// 				AllowForwardedTraffic: to.Ptr(false),
		// 				AllowGatewayTransit: to.Ptr(false),
		// 				AllowVirtualNetworkAccess: to.Ptr(true),
		// 				DatabricksAddressSpace: &armdatabricks.AddressSpace{
		// 					AddressPrefixes: []*string{
		// 						to.Ptr("10.139.0.0/16")},
		// 					},
		// 					DatabricksVirtualNetwork: &armdatabricks.VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork{
		// 						ID: to.Ptr("/subscriptions/0140911e-1040-48da-8bc9-b99fb3dd88a6/resourceGroups/databricks-rg-adbworkspace-2jsxhmzoyooxm/providers/Microsoft.Network/virtualNetworks/workers-vnet"),
		// 					},
		// 					PeeringState: to.Ptr(armdatabricks.PeeringStateInitiated),
		// 					ProvisioningState: to.Ptr(armdatabricks.PeeringProvisioningStateSucceeded),
		// 					RemoteAddressSpace: &armdatabricks.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("10.203.0.0/16")},
		// 						},
		// 						RemoteVirtualNetwork: &armdatabricks.VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork{
		// 							ID: to.Ptr("/subscriptions/0140911e-1040-48da-8bc9-b99fb3dd88a6/resourceGroups/subramantest/providers/Microsoft.Network/virtualNetworks/subramanvnet"),
		// 						},
		// 						UseRemoteGateways: to.Ptr(false),
		// 					},
		// 			}},
		// 		}
	}
}
