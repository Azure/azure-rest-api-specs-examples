package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_Get.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetVMRunCommandsClient().Get(ctx, "myResourceGroup", "myvmScaleSet", "0", "myRunCommand", &armcompute.VirtualMachineScaleSetVMRunCommandsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineRunCommand = armcompute.VirtualMachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/runCommands"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/runCommands/myRunCommand"),
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
	// 			ScriptURI: to.Ptr("https://gist.githubusercontent.com/myusername/75fd3634w7511116063c60bcc50bee0/raw/04a4c68ac9e1d36asfasdc64bd1d889b104c7abdb8/HelloWorld.ps1"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 		TreatFailureAsDeploymentFailure: to.Ptr(false),
	// 	},
	// }
}
