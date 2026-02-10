package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ExpressRoutePortGet.json
func ExampleExpressRoutePortsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRoutePortsClient().Get(ctx, "rg1", "portName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRoutePort = armnetwork.ExpressRoutePort{
	// 	Name: to.Ptr("portName"),
	// 	Type: to.Ptr("Microsoft.Network/expressRoutePorts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.ExpressRoutePortPropertiesFormat{
	// 		AllocationDate: to.Ptr("Friday, July 1, 2018"),
	// 		BandwidthInGbps: to.Ptr[int32](100),
	// 		BillingType: to.Ptr(armnetwork.ExpressRoutePortsBillingTypeUnlimitedData),
	// 		Circuits: []*armnetwork.SubResource{
	// 		},
	// 		Encapsulation: to.Ptr(armnetwork.ExpressRoutePortsEncapsulationQinQ),
	// 		EtherType: to.Ptr("0x8100"),
	// 		Links: []*armnetwork.ExpressRouteLink{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link1"),
	// 				Name: to.Ptr("link1"),
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
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link2"),
	// 				Name: to.Ptr("link2"),
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
	// 		}},
	// 		Mtu: to.Ptr("1500"),
	// 		PeeringLocation: to.Ptr("peeringLocationName"),
	// 		ProvisionedBandwidthInGbps: to.Ptr[float32](0),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
