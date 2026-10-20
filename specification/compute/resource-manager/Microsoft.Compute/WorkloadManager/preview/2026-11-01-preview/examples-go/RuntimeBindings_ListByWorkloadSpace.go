package armworkloadmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armworkloadmanager"
)

// Generated from example definition: 2026-11-01-preview/RuntimeBindings_ListByWorkloadSpace.json
func ExampleRuntimeBindingsClient_NewListByWorkloadSpacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadmanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRuntimeBindingsClient().NewListByWorkloadSpacePager("rg-workload", "managed-agents-prod", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armworkloadmanager.RuntimeBindingsClientListByWorkloadSpaceResponse{
		// 	RuntimeBindingListResult: armworkloadmanager.RuntimeBindingListResult{
		// 		Value: []*armworkloadmanager.RuntimeBinding{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-workload/providers/Microsoft.Compute/workloadSpaces/managed-agents-prod/runtimeBindings/kubernetes-default"),
		// 				Name: to.Ptr("kubernetes-default"),
		// 				Type: to.Ptr("Microsoft.Compute/workloadSpaces/runtimeBindings"),
		// 				Location: to.Ptr("eastus2"),
		// 				Kind: to.Ptr(armworkloadmanager.RuntimeBindingKindKubernetes),
		// 				Properties: &armworkloadmanager.ManagedRuntimeBindingProperties{
		// 					ProvisioningMode: to.Ptr(armworkloadmanager.RuntimeBindingProvisioningModeManaged),
		// 					ManagedProfile: &armworkloadmanager.ManagedRuntimeProfile{
		// 						Offering: to.Ptr("Automatic"),
		// 					},
		// 					ProvisioningState: to.Ptr(armworkloadmanager.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
