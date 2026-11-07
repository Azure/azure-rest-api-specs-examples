package armnetworksecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networksecurity/armnetworksecurity"
)

// Generated from example definition: 2026-10-01-preview/NeuroShieldProtectedResources_Get.json
func ExampleNeuroShieldProtectedResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworksecurity.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNeuroShieldProtectedResourcesClient().Get(ctx, "myResourceGroup", "myNeuroShield", "vm-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetworksecurity.NeuroShieldProtectedResourcesClientGetResponse{
	// 	NeuroShieldProtectedResource: armnetworksecurity.NeuroShieldProtectedResource{
	// 		Name: to.Ptr("vm-01"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/microsoft.networksecurity/neuroShields/myNeuroShield/protectedResources/vm-01"),
	// 		Type: to.Ptr("microsoft.networksecurity/neuroShields/protectedResources"),
	// 		Properties: &armnetworksecurity.NeuroShieldProtectedResourceProperties{
	// 			ProvisioningState: to.Ptr(armnetworksecurity.ProvisioningStateNeuroShieldSucceeded),
	// 			TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myComputeRg/providers/Microsoft.Compute/virtualMachines/vm-01"),
	// 			AgentHealth: &armnetworksecurity.NeuroShieldAgentHealth{
	// 				LastSeenTime: to.Ptr(time.Date(2026, time.May, 28, 10, 30, 0, 0, time.UTC)),
	// 			},
	// 			InventoryUploadRequest: &armnetworksecurity.NeuroShieldInventoryUploadRequestMetadata{
	// 				RequestID: to.Ptr("33333333-3333-3333-3333-333333333333"),
	// 				RequestedAt: to.Ptr(time.Date(2026, time.May, 28, 10, 20, 0, 0, time.UTC)),
	// 				ExpiresAt: to.Ptr(time.Date(2026, time.May, 28, 10, 35, 0, 0, time.UTC)),
	// 			},
	// 			TargetIdentity: &armnetworksecurity.NeuroShieldProtectedResourceIdentity{
	// 				Type: to.Ptr(armnetworksecurity.ManagedServiceIdentityTypeSystemAssigned),
	// 				PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 			},
	// 		},
	// 	},
	// }
}
