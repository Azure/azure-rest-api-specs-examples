package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-05-01/NetworkVirtualApplianceVhubSourceIPProtocolPut.json
func ExampleVirtualAppliancesClient_BeginCreateOrUpdate_createOrUpdateAManagedNvaInAVirtualHubWithSourceIPProtocolLoadDistribution() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualAppliancesClient().BeginCreateOrUpdate(ctx, "rg1", "nva-vhub", armnetwork.VirtualAppliance{
		Location: to.Ptr("West US"),
		Properties: &armnetwork.VirtualAppliancePropertiesFormat{
			NvaSKU: &armnetwork.VirtualApplianceSKUProperties{
				BundledScaleUnit:   to.Ptr("1"),
				MarketPlaceVersion: to.Ptr("12.1"),
				Vendor:             to.Ptr("Cisco SDWAN"),
			},
			VirtualHub: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
			},
			AddressFamily: []*armnetwork.IPVersion{
				to.Ptr(armnetwork.IPVersionIPv4),
			},
			VirtualApplianceAsn: to.Ptr[int64](10000),
			LoadBalancerConfiguration: &armnetwork.VirtualApplianceLoadBalancerConfiguration{
				LoadDistribution:     to.Ptr(armnetwork.NetworkVirtualApplianceLoadDistributionSourceIPProtocol),
				IdleTimeoutInMinutes: to.Ptr[int32](15),
			},
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
	// res = armnetwork.VirtualAppliancesClientCreateOrUpdateResponse{
	// 	VirtualAppliance: armnetwork.VirtualAppliance{
	// 		Name: to.Ptr("nva-vhub"),
	// 		Type: to.Ptr("Microsoft.Network/networkVirtualAppliances"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva-vhub"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armnetwork.VirtualAppliancePropertiesFormat{
	// 			NvaSKU: &armnetwork.VirtualApplianceSKUProperties{
	// 				BundledScaleUnit: to.Ptr("1"),
	// 				MarketPlaceVersion: to.Ptr("12.1"),
	// 				Vendor: to.Ptr("Cisco SDWAN"),
	// 			},
	// 			VirtualHub: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
	// 			},
	// 			AddressFamily: []*armnetwork.IPVersion{
	// 				to.Ptr(armnetwork.IPVersionIPv4),
	// 			},
	// 			VirtualApplianceAsn: to.Ptr[int64](10000),
	// 			LoadBalancerConfiguration: &armnetwork.VirtualApplianceLoadBalancerConfiguration{
	// 				LoadDistribution: to.Ptr(armnetwork.NetworkVirtualApplianceLoadDistributionSourceIPProtocol),
	// 				IdleTimeoutInMinutes: to.Ptr[int32](15),
	// 			},
	// 			PrivateIPAddress: to.Ptr("192.168.12.11"),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
