package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-03-01/VirtualNetworkApplianceCapabilities_CreateOrUpdate_PLGateway.json
func ExampleVirtualNetworkApplianceCapabilitiesClient_BeginCreateOrUpdate_createVirtualNetworkAppliancePlGatewayCapability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkApplianceCapabilitiesClient().BeginCreateOrUpdate(ctx, "rg1", "test-vna", "pl-gateway", &armnetwork.PLGatewayCapability{
		Kind: to.Ptr(armnetwork.VirtualNetworkApplianceCapabilityKindPLGateway),
		Properties: &armnetwork.VirtualNetworkApplianceCapabilityProperties{
			IPVersion: to.Ptr(armnetwork.VirtualNetworkApplianceCapabilityIPVersionIPv6),
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
	// res = armnetwork.VirtualNetworkApplianceCapabilitiesClientCreateOrUpdateResponse{
	// 	VirtualNetworkApplianceCapabilityClassification: &armnetwork.PLGatewayCapability{
	// 		Name: to.Ptr("pl-gateway"),
	// 		Type: to.Ptr("Microsoft.Network/virtualNetworkAppliances/capabilities"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkAppliances/test-vna/capabilities/pl-gateway"),
	// 		Kind: to.Ptr(armnetwork.VirtualNetworkApplianceCapabilityKindPLGateway),
	// 		Properties: &armnetwork.VirtualNetworkApplianceCapabilityProperties{
	// 			IPVersion: to.Ptr(armnetwork.VirtualNetworkApplianceCapabilityIPVersionIPv6),
	// 			LinkedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet"),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
