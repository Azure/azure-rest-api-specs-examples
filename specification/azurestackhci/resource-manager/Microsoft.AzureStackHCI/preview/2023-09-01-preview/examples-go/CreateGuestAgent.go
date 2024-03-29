package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/CreateGuestAgent.json
func ExampleGuestAgentClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGuestAgentClient().BeginCreate(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM", &armazurestackhci.GuestAgentClientBeginCreateOptions{Body: &armazurestackhci.GuestAgent{
		Properties: &armazurestackhci.GuestAgentProperties{
			Credentials: &armazurestackhci.GuestCredential{
				Password: to.Ptr("<password>"),
				Username: to.Ptr("tempuser"),
			},
			ProvisioningAction: to.Ptr(armazurestackhci.ProvisioningActionInstall),
		},
	},
	})
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
	// res.GuestAgent = armazurestackhci.GuestAgent{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/virtualMachineInstances/guestAgents"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.AzureStackHCI/virtualMachineInstances/default/guestAgents/default"),
	// 	Properties: &armazurestackhci.GuestAgentProperties{
	// 		ProvisioningAction: to.Ptr(armazurestackhci.ProvisioningActionInstall),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Status: to.Ptr("connected"),
	// 	},
	// }
}
