package armworkloadmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armworkloadmanager"
)

// Generated from example definition: 2026-11-01-preview/RuntimeBindings_CreateOrUpdate_ReferencedKubernetes.json
func ExampleRuntimeBindingsClient_BeginCreateOrUpdate_attachAReferencedKubernetesRuntimeBinding() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadmanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRuntimeBindingsClient().BeginCreateOrUpdate(ctx, "rg-workload", "managed-agents-prod", "customer-aks", armworkloadmanager.RuntimeBinding{
		Location: to.Ptr("eastus2"),
		Kind:     to.Ptr(armworkloadmanager.RuntimeBindingKindKubernetes),
		Properties: &armworkloadmanager.ReferencedRuntimeBindingProperties{
			ProvisioningMode: to.Ptr(armworkloadmanager.RuntimeBindingProvisioningModeReferenced),
			ResourceID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aks-rg/providers/Microsoft.ContainerService/managedClusters/aks-prod"),
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
	// res = armworkloadmanager.RuntimeBindingsClientCreateOrUpdateResponse{
	// 	RuntimeBinding: armworkloadmanager.RuntimeBinding{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-workload/providers/Microsoft.Compute/workloadSpaces/managed-agents-prod/runtimeBindings/customer-aks"),
	// 		Name: to.Ptr("customer-aks"),
	// 		Type: to.Ptr("Microsoft.Compute/workloadSpaces/runtimeBindings"),
	// 		Location: to.Ptr("eastus2"),
	// 		Kind: to.Ptr(armworkloadmanager.RuntimeBindingKindKubernetes),
	// 		Properties: &armworkloadmanager.ReferencedRuntimeBindingProperties{
	// 			ProvisioningMode: to.Ptr(armworkloadmanager.RuntimeBindingProvisioningModeReferenced),
	// 			ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aks-rg/providers/Microsoft.ContainerService/managedClusters/aks-prod"),
	// 			ProvisioningState: to.Ptr(armworkloadmanager.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
