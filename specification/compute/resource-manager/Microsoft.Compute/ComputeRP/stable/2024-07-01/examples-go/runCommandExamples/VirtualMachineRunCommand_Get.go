package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/runCommandExamples/VirtualMachineRunCommand_Get.json
func ExampleVirtualMachineRunCommandsClient_GetByVirtualMachine() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineRunCommandsClient().GetByVirtualMachine(ctx, "myResourceGroup", "myVM", "myRunCommand", &armcompute.VirtualMachineRunCommandsClientGetByVirtualMachineOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineRunCommand = armcompute.VirtualMachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachines/runCommands"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/runCommands/myRunCommand"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineRunCommandProperties{
	// 		AsyncExecution: to.Ptr(false),
	// 		ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
	// 		OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
	// 		Parameters: []*armcompute.RunCommandInputParameter{
	// 			{
	// 				Name: to.Ptr("param1"),
	// 				Value: to.Ptr("value1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("param2"),
	// 				Value: to.Ptr("value2"),
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RunAsUser: to.Ptr("user1"),
	// 		Source: &armcompute.VirtualMachineRunCommandScriptSource{
	// 			Script: to.Ptr("Write-Host Hello World! ; Remove-Item C:	est	estFile.txt"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 		TreatFailureAsDeploymentFailure: to.Ptr(false),
	// 	},
	// }
}
