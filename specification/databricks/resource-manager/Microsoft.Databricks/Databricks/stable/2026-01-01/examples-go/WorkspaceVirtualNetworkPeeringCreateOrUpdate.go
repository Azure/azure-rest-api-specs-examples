package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks/v2"
)

// Generated from example definition: 2026-01-01/WorkspaceVirtualNetworkPeeringCreateOrUpdate.json
func ExampleVNetPeeringClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("0140911e-1040-48da-8bc9-b99fb3dd88a6/", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVNetPeeringClient().BeginCreateOrUpdate(ctx, "subramantest", "adbworkspace", "vNetPeeringTest", armdatabricks.VirtualNetworkPeering{
		Properties: &armdatabricks.VirtualNetworkPeeringPropertiesFormat{
			AllowForwardedTraffic:     to.Ptr(false),
			AllowGatewayTransit:       to.Ptr(false),
			AllowVirtualNetworkAccess: to.Ptr(true),
			RemoteVirtualNetwork: &armdatabricks.VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork{
				ID: to.Ptr("/subscriptions/0140911e-1040-48da-8bc9-b99fb3dd88a6/resourceGroups/subramantest/providers/Microsoft.Network/virtualNetworks/subramanvnet"),
			},
			UseRemoteGateways: to.Ptr(false),
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
	// res = armdatabricks.VNetPeeringClientCreateOrUpdateResponse{
	// 	VirtualNetworkPeering: armdatabricks.VirtualNetworkPeering{
	// 		Name: to.Ptr("vNetPeeringTest"),
	// 		ID: to.Ptr("/subscriptions/0140911e-1040-48da-8bc9-b99fb3dd88a6/resourceGroups/subramantest/providers/Microsoft.Databricks/workspaces/adbworkspace/virtualNetworkPeerings/vNetPeeringTest"),
	// 		Properties: &armdatabricks.VirtualNetworkPeeringPropertiesFormat{
	// 			AllowForwardedTraffic: to.Ptr(false),
	// 			AllowGatewayTransit: to.Ptr(false),
	// 			AllowVirtualNetworkAccess: to.Ptr(true),
	// 			DatabricksAddressSpace: &armdatabricks.AddressSpace{
	// 				AddressPrefixes: []*string{
	// 					to.Ptr("10.139.0.0/16"),
	// 				},
	// 			},
	// 			DatabricksVirtualNetwork: &armdatabricks.VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork{
	// 				ID: to.Ptr("/subscriptions/0140911e-1040-48da-8bc9-b99fb3dd88a6/resourceGroups/databricks-rg-adbworkspace-2jsxhmzoyooxm/providers/Microsoft.Network/virtualNetworks/workers-vnet"),
	// 			},
	// 			PeeringState: to.Ptr(armdatabricks.PeeringStateInitiated),
	// 			ProvisioningState: to.Ptr(armdatabricks.PeeringProvisioningStateUpdating),
	// 			RemoteAddressSpace: &armdatabricks.AddressSpace{
	// 				AddressPrefixes: []*string{
	// 					to.Ptr("10.203.0.0/16"),
	// 				},
	// 			},
	// 			RemoteVirtualNetwork: &armdatabricks.VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork{
	// 				ID: to.Ptr("/subscriptions/0140911e-1040-48da-8bc9-b99fb3dd88a6/resourceGroups/subramantest/providers/Microsoft.Network/virtualNetworks/subramanvnet"),
	// 			},
	// 			UseRemoteGateways: to.Ptr(false),
	// 		},
	// 	},
	// }
}
