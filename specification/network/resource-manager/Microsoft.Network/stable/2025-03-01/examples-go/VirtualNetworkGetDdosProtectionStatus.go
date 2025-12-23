package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkGetDdosProtectionStatus.json
func ExampleVirtualNetworksClient_BeginListDdosProtectionStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworksClient().BeginListDdosProtectionStatus(ctx, "rg1", "test-vnet", &armnetwork.VirtualNetworksClientBeginListDdosProtectionStatusOptions{Top: to.Ptr[int32](75),
		SkipToken: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	for res.More() {
		page, err := res.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualNetworkDdosProtectionStatusResult = armnetwork.VirtualNetworkDdosProtectionStatusResult{
		// 	Value: []*armnetwork.PublicIPDdosProtectionStatusResult{
		// 		{
		// 			DdosProtectionPlanID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ddosProtectionPlans/test-plan"),
		// 			IsWorkloadProtected: to.Ptr(armnetwork.IsWorkloadProtectedTrue),
		// 			PublicIPAddress: to.Ptr("10.0.1.5"),
		// 			PublicIPAddressID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-pip"),
		// 		},
		// 		{
		// 			IsWorkloadProtected: to.Ptr(armnetwork.IsWorkloadProtectedFalse),
		// 			PublicIPAddress: to.Ptr("10.0.1.6"),
		// 			PublicIPAddressID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-pip2"),
		// 	}},
		// }
	}
}
