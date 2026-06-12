package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/PublicIpAddressGetDdosProtectionStatus.json
func ExamplePublicIPAddressesClient_BeginDdosProtectionStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPublicIPAddressesClient().BeginDdosProtectionStatus(ctx, "rg1", "test-pip", nil)
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
	// res = armnetwork.PublicIPAddressesClientDdosProtectionStatusResponse{
	// 	PublicIPDdosProtectionStatusResult: armnetwork.PublicIPDdosProtectionStatusResult{
	// 		DdosProtectionPlanID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/ddosProtectionPlans/test-plan"),
	// 		IsWorkloadProtected: to.Ptr(armnetwork.IsWorkloadProtectedTrue),
	// 		PublicIPAddress: to.Ptr("10.0.1.5"),
	// 		PublicIPAddressID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-pip"),
	// 	},
	// }
}
