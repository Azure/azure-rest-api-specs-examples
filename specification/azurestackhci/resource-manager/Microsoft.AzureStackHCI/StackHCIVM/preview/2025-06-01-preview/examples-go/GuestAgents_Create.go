package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/GuestAgents_Create.json
func ExampleGuestAgentsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGuestAgentsClient().BeginCreate(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", armazurestackhcivm.GuestAgent{
		Properties: &armazurestackhcivm.GuestAgentProperties{
			Credentials: &armazurestackhcivm.GuestCredential{
				Password: to.Ptr("<password>"),
				Username: to.Ptr("tempuser"),
			},
			ProvisioningAction: to.Ptr(armazurestackhcivm.ProvisioningActionInstall),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhcivm.GuestAgentsClientCreateResponse{
	// 	GuestAgent: &armazurestackhcivm.GuestAgent{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/virtualMachineInstances/guestAgents"),
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.AzureStackHCI/virtualMachineInstances/default/guestAgents/default"),
	// 		Properties: &armazurestackhcivm.GuestAgentProperties{
	// 			ProvisioningAction: to.Ptr(armazurestackhcivm.ProvisioningActionInstall),
	// 			ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumSucceeded),
	// 			Status: to.Ptr("connected"),
	// 		},
	// 	},
	// }
}
