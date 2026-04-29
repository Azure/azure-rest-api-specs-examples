package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/runCommandExamples/VirtualMachineRunCommand_List.json
func ExampleVirtualMachineRunCommandsClient_NewListByVirtualMachinePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineRunCommandsClient().NewListByVirtualMachinePager("myResourceGroup", "myVM", nil)
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
		// page = armcompute.VirtualMachineRunCommandsClientListByVirtualMachineResponse{
		// 	VirtualMachineRunCommandsListResult: armcompute.VirtualMachineRunCommandsListResult{
		// 		Value: []*armcompute.VirtualMachineRunCommand{
		// 			{
		// 				Name: to.Ptr("myRunCommand"),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/runCommands/myRunCommand"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachines/runCommands"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("value1"),
		// 					"tag2": to.Ptr("value2"),
		// 				},
		// 				Properties: &armcompute.VirtualMachineRunCommandProperties{
		// 					Source: &armcompute.VirtualMachineRunCommandScriptSource{
		// 						Script: to.Ptr("Write-Host Hello World!"),
		// 					},
		// 					Parameters: []*armcompute.RunCommandInputParameter{
		// 						{
		// 							Name: to.Ptr("param1"),
		// 							Value: to.Ptr("value1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("param2"),
		// 							Value: to.Ptr("value2"),
		// 						},
		// 					},
		// 					AsyncExecution: to.Ptr(false),
		// 					TreatFailureAsDeploymentFailure: to.Ptr(false),
		// 					RunAsUser: to.Ptr("user1"),
		// 					TimeoutInSeconds: to.Ptr[int32](0),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
		// 					ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
