package armworkloadmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armworkloadmanager"
)

// Generated from example definition: 2026-11-01-preview/RuntimeBindings_Update.json
func ExampleRuntimeBindingsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadmanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRuntimeBindingsClient().BeginUpdate(ctx, "rg-workload", "managed-agents-prod", "serverless-default", armworkloadmanager.RuntimeBindingUpdate{
		Properties: &armworkloadmanager.RuntimeBindingUpdateProperties{
			NetworkProfile: &armworkloadmanager.RuntimeNetworkProfile{
				SubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-network/providers/Microsoft.Network/virtualNetworks/workload-vnet/subnets/execution"),
				EgressMode:       to.Ptr(armworkloadmanager.EgressModeCustomerManaged),
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
	// res = armworkloadmanager.RuntimeBindingsClientUpdateResponse{
	// 	RuntimeBinding: armworkloadmanager.RuntimeBinding{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-workload/providers/Microsoft.Compute/workloadSpaces/managed-agents-prod/runtimeBindings/serverless-default"),
	// 		Name: to.Ptr("serverless-default"),
	// 		Type: to.Ptr("Microsoft.Compute/workloadSpaces/runtimeBindings"),
	// 		Location: to.Ptr("eastus2"),
	// 		Kind: to.Ptr(armworkloadmanager.RuntimeBindingKindServerlessContainers),
	// 		Properties: &armworkloadmanager.ManagedRuntimeBindingProperties{
	// 			ProvisioningMode: to.Ptr(armworkloadmanager.RuntimeBindingProvisioningModeManaged),
	// 			ManagedProfile: &armworkloadmanager.ManagedRuntimeProfile{
	// 				Provider: to.Ptr("ACI"),
	// 			},
	// 			NetworkProfile: &armworkloadmanager.RuntimeNetworkProfile{
	// 				SubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-network/providers/Microsoft.Network/virtualNetworks/workload-vnet/subnets/execution"),
	// 				EgressMode: to.Ptr(armworkloadmanager.EgressModeCustomerManaged),
	// 			},
	// 			ProvisioningState: to.Ptr(armworkloadmanager.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
