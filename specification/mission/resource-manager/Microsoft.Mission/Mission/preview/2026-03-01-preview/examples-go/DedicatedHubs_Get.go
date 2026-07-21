package armenclave_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/enclave/armenclave"
)

// Generated from example definition: 2026-03-01-preview/DedicatedHubs_Get.json
func ExampleDedicatedHubClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armenclave.NewClientFactory("c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHubClient().Get(ctx, "TestResourceGroup", "TestCommunity", "TestDedicatedHub", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armenclave.DedicatedHubClientGetResponse{
	// 	DedicatedHubResource: armenclave.DedicatedHubResource{
	// 		ID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestResourceGroup/providers/Microsoft.Mission/communities/TestCommunity/dedicatedHubs/TestDedicatedHub"),
	// 		Name: to.Ptr("TestDedicatedHub"),
	// 		Type: to.Ptr("Microsoft.Mission/communities/dedicatedHubs"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armenclave.DedicatedHubProperties{
	// 			ProvisioningState: to.Ptr(armenclave.ProvisioningStateSucceeded),
	// 			VHubResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestResourceGroup/providers/Microsoft.Network/virtualHubs/TestVirtualHub"),
	// 			FirewallResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestResourceGroup/providers/Microsoft.Network/azureFirewalls/TestFirewall"),
	// 			FirewallPolicyResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestResourceGroup/providers/Microsoft.Network/firewallPolicies/TestFirewallPolicy"),
	// 			Designation: to.Ptr(armenclave.DesignationReserved),
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("test"),
	// 			"project": to.Ptr("mission"),
	// 		},
	// 		SystemData: &armenclave.SystemData{
	// 			CreatedBy: to.Ptr("testUser"),
	// 			CreatedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("testUser"),
	// 			LastModifiedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 		},
	// 	},
	// }
}
