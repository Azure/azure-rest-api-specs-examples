package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b749953e21e5c3f275d839862323920ef7bf716e/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VirtualNetworkLink_Put.json
func ExampleVirtualNetworkLinksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkLinksClient().BeginCreateOrUpdate(ctx, "sampleResourceGroup", "sampleDnsForwardingRuleset", "sampleVirtualNetworkLink", armdnsresolver.VirtualNetworkLink{
		Properties: &armdnsresolver.VirtualNetworkLinkProperties{
			Metadata: map[string]*string{
				"additionalProp1": to.Ptr("value1"),
			},
			VirtualNetwork: &armdnsresolver.SubResource{
				ID: to.Ptr("/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork"),
			},
		},
	}, &armdnsresolver.VirtualNetworkLinksClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkLink = armdnsresolver.VirtualNetworkLink{
	// 	Name: to.Ptr("sampleVirtualNetworkLink"),
	// 	Type: to.Ptr("Microsoft.Network/dnsForwardingRulesets/virtualNetworkLinks"),
	// 	ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsForwardingRuleset/sampleDnsForwardingRuleset/virtualNetworkLinks/sampleVirtualNetworkLink"),
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armdnsresolver.VirtualNetworkLinkProperties{
	// 		Metadata: map[string]*string{
	// 			"additionalProp1": to.Ptr("value1"),
	// 		},
	// 		ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
	// 		VirtualNetwork: &armdnsresolver.SubResource{
	// 			ID: to.Ptr("/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork"),
	// 		},
	// 	},
	// 	SystemData: &armdnsresolver.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.1075056Z"); return t}()),
	// 		CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.1974346Z"); return t}()),
	// 		LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 	},
	// }
}
