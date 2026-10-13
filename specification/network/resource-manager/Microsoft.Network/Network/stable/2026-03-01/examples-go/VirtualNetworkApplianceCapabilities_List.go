package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-03-01/VirtualNetworkApplianceCapabilities_List.json
func ExampleVirtualNetworkApplianceCapabilitiesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkApplianceCapabilitiesClient().NewListPager("rg1", "test-vna", nil)
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
		// page = armnetwork.VirtualNetworkApplianceCapabilitiesClientListResponse{
		// 	VirtualNetworkApplianceCapabilityListResult: armnetwork.VirtualNetworkApplianceCapabilityListResult{
		// 		Value: []armnetwork.VirtualNetworkApplianceCapabilityClassification{
		// 			&armnetwork.PLGatewayFastpathCapability{
		// 				Name: to.Ptr("pl-fastpath"),
		// 				Type: to.Ptr("Microsoft.Network/virtualNetworkAppliances/capabilities"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkAppliances/test-vna/capabilities/pl-fastpath"),
		// 				Kind: to.Ptr(armnetwork.VirtualNetworkApplianceCapabilityKindPLGatewayFastpath),
		// 				Properties: &armnetwork.VirtualNetworkApplianceCapabilityProperties{
		// 					IPVersion: to.Ptr(armnetwork.VirtualNetworkApplianceCapabilityIPVersionDualStack),
		// 					LinkedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet"),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			&armnetwork.PLIPForwardersCapability{
		// 				Name: to.Ptr("pl-ipforwarders"),
		// 				Type: to.Ptr("Microsoft.Network/virtualNetworkAppliances/capabilities"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkAppliances/test-vna/capabilities/pl-ipforwarders"),
		// 				Kind: to.Ptr(armnetwork.VirtualNetworkApplianceCapabilityKindPLIPForwarders),
		// 				Properties: &armnetwork.VirtualNetworkApplianceCapabilityProperties{
		// 					IPVersion: to.Ptr(armnetwork.VirtualNetworkApplianceCapabilityIPVersionIPv6),
		// 					LinkedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet"),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
