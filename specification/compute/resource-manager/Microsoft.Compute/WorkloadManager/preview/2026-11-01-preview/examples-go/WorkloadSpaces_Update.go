package armworkloadmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armworkloadmanager"
)

// Generated from example definition: 2026-11-01-preview/WorkloadSpaces_Update.json
func ExampleWorkloadSpacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadmanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadSpacesClient().BeginUpdate(ctx, "rg-workload", "managed-agents-prod", armworkloadmanager.WorkloadSpaceUpdate{
		Tags: map[string]*string{
			"environment": to.Ptr("Production"),
			"costCenter":  to.Ptr("AI-Platform"),
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
	// res = armworkloadmanager.WorkloadSpacesClientUpdateResponse{
	// 	WorkloadSpace: armworkloadmanager.WorkloadSpace{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-workload/providers/Microsoft.Compute/workloadSpaces/managed-agents-prod"),
	// 		Name: to.Ptr("managed-agents-prod"),
	// 		Type: to.Ptr("Microsoft.Compute/workloadSpaces"),
	// 		Location: to.Ptr("eastus2"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("Production"),
	// 			"costCenter": to.Ptr("AI-Platform"),
	// 		},
	// 		Properties: &armworkloadmanager.WorkloadSpaceProperties{
	// 			ProvisioningState: to.Ptr(armworkloadmanager.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
