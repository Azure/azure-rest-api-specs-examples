package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/GuestAgents_ListByVirtualMachineInstance.json
func ExampleGuestAgentsClient_NewListByVirtualMachineInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGuestAgentsClient().NewListByVirtualMachineInstancePager("subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", nil)
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
		// page = armazurestackhcivm.GuestAgentsClientListByVirtualMachineInstanceResponse{
		// 	GuestAgentListResult: armazurestackhcivm.GuestAgentListResult{
		// 		Value: []*armazurestackhcivm.GuestAgent{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/virtualMachineInstances/guestAgents"),
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/virtualMachineInstances/default/guestAgents/default"),
		// 				Properties: &armazurestackhcivm.GuestAgentProperties{
		// 					ProvisioningAction: to.Ptr(armazurestackhcivm.ProvisioningActionInstall),
		// 					ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumSucceeded),
		// 					Status: to.Ptr("connected"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
