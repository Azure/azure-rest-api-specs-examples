package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_List.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineScaleSetVMRunCommandsClient().NewListPager("myResourceGroup", "myvmScaleSet", "0", &armcompute.VirtualMachineScaleSetVMRunCommandsClientListOptions{Expand: nil})
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
		// page.VirtualMachineRunCommandsListResult = armcompute.VirtualMachineRunCommandsListResult{
		// 	Value: []*armcompute.VirtualMachineRunCommand{
		// 		{
		// 			Name: to.Ptr("myRunCommand"),
		// 			Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/runCommands"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/runCommands/myRunCommand"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcompute.VirtualMachineRunCommandProperties{
		// 				AsyncExecution: to.Ptr(false),
		// 				ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
		// 				OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
		// 				Parameters: []*armcompute.RunCommandInputParameter{
		// 					{
		// 						Name: to.Ptr("param1"),
		// 						Value: to.Ptr("value1"),
		// 					},
		// 					{
		// 						Name: to.Ptr("param2"),
		// 						Value: to.Ptr("value2"),
		// 				}},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RunAsUser: to.Ptr("user1"),
		// 				Source: &armcompute.VirtualMachineRunCommandScriptSource{
		// 					Script: to.Ptr("Write-Host Hello World! ; Remove-Item C:	est	estFile.txt"),
		// 				},
		// 				TimeoutInSeconds: to.Ptr[int32](0),
		// 				TreatFailureAsDeploymentFailure: to.Ptr(false),
		// 			},
		// 	}},
		// }
	}
}
