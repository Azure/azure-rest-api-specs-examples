package armworkloadmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armworkloadmanager"
)

// Generated from example definition: 2026-11-01-preview/WorkloadSpaces_Get.json
func ExampleWorkloadSpacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadmanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkloadSpacesClient().Get(ctx, "rg-workload", "managed-agents-prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armworkloadmanager.WorkloadSpacesClientGetResponse{
	// 	WorkloadSpace: armworkloadmanager.WorkloadSpace{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-workload/providers/Microsoft.Compute/workloadSpaces/managed-agents-prod"),
	// 		Name: to.Ptr("managed-agents-prod"),
	// 		Type: to.Ptr("Microsoft.Compute/workloadSpaces"),
	// 		Location: to.Ptr("eastus2"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("Production"),
	// 		},
	// 		Properties: &armworkloadmanager.WorkloadSpaceProperties{
	// 			ProvisioningState: to.Ptr(armworkloadmanager.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
