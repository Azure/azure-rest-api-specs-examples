package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkVirtualApplianceSaaSPut.json
func ExampleVirtualAppliancesClient_BeginCreateOrUpdate_createSaaSNetworkVirtualAppliance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualAppliancesClient().BeginCreateOrUpdate(ctx, "rg1", "nva", armnetwork.VirtualAppliance{
		Location: to.Ptr("West US"),
		Properties: &armnetwork.VirtualAppliancePropertiesFormat{
			Delegation: &armnetwork.DelegationProperties{
				ServiceName: to.Ptr("PaloAltoNetworks.Cloudngfw/firewalls"),
			},
			VirtualHub: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
			},
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
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
	// 		Name: to.Ptr("nva"),
	// 		Type: to.Ptr("Microsoft.Network/networkVirtualAppliances"),
	// 		Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armnetwork.VirtualAppliancePropertiesFormat{
	// 			Delegation: &armnetwork.DelegationProperties{
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				ServiceName: to.Ptr("PaloAltoNetworks.Cloudngfw/firewalls"),
	// 			},
	// 			DeploymentType: to.Ptr("PartnerManaged"),
	// 			PartnerManagedResource: &armnetwork.PartnerManagedResourceProperties{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Lifter/firewalls/paloAltoFirewall"),
	// 				InternalLoadBalancerID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/trustILB"),
	// 				StandardLoadBalancerID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/trustSLB"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			VirtualHub: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
