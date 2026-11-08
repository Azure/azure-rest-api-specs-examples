package armnetworksecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networksecurity/armnetworksecurity"
)

// Generated from example definition: 2026-10-01-preview/NeuroShieldProtectedResources_CreateOrUpdate_NoneIdentity.json
func ExampleNeuroShieldProtectedResourcesClient_BeginCreateOrUpdate_createOrUpdateProtectedResourceWithExplicitNoneIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworksecurity.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNeuroShieldProtectedResourcesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myNeuroShield", "vm-01", armnetworksecurity.NeuroShieldProtectedResource{
		Properties: &armnetworksecurity.NeuroShieldProtectedResourceProperties{
			TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myComputeRg/providers/Microsoft.Compute/virtualMachines/vm-01"),
			TargetIdentity: &armnetworksecurity.NeuroShieldProtectedResourceIdentity{
				Type: to.Ptr(armnetworksecurity.ManagedServiceIdentityTypeNone),
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
	// res = armnetworksecurity.NeuroShieldProtectedResourcesClientCreateOrUpdateResponse{
	// 	NeuroShieldProtectedResource: armnetworksecurity.NeuroShieldProtectedResource{
	// 		Name: to.Ptr("vm-01"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/microsoft.networksecurity/neuroShields/myNeuroShield/protectedResources/vm-01"),
	// 		Type: to.Ptr("microsoft.networksecurity/neuroShields/protectedResources"),
	// 		Properties: &armnetworksecurity.NeuroShieldProtectedResourceProperties{
	// 			ProvisioningState: to.Ptr(armnetworksecurity.ProvisioningStateNeuroShieldSucceeded),
	// 			TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myComputeRg/providers/Microsoft.Compute/virtualMachines/vm-01"),
	// 			TargetIdentity: &armnetworksecurity.NeuroShieldProtectedResourceIdentity{
	// 				Type: to.Ptr(armnetworksecurity.ManagedServiceIdentityTypeNone),
	// 			},
	// 		},
	// 	},
	// }
}
