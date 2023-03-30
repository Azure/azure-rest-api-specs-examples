package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GuestAgent_ListByVm.json
func ExampleGuestAgentsClient_NewListByVMPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGuestAgentsClient().NewListByVMPager("testrg", "ContosoVm", nil)
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
		// page.GuestAgentList = armconnectedvmware.GuestAgentList{
		// 	Value: []*armconnectedvmware.GuestAgent{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VitualMachines/guestAgents"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VitualMachines/ContosoVm/guestAgents/default"),
		// 			Properties: &armconnectedvmware.GuestAgentProperties{
		// 				ProvisioningAction: to.Ptr(armconnectedvmware.ProvisioningActionInstall),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Status: to.Ptr("connected"),
		// 			},
		// 	}},
		// }
	}
}
