package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5d2adf9b7fda669b4a2538c65e937ee74fe3f966/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_CreateOrUpdate.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMRunCommandsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myvmScaleSet", "0", "myRunCommand", armcompute.VirtualMachineRunCommand{
		Location: to.Ptr("West US"),
		Properties: &armcompute.VirtualMachineRunCommandProperties{
			AsyncExecution:           to.Ptr(false),
			ErrorBlobManagedIdentity: &armcompute.RunCommandManagedIdentity{},
			ErrorBlobURI:             to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
			OutputBlobManagedIdentity: &armcompute.RunCommandManagedIdentity{
				ClientID: to.Ptr("22d35efb-0c99-4041-8c5b-6d24db33a69a"),
			},
			OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
			Parameters: []*armcompute.RunCommandInputParameter{
				{
					Name:  to.Ptr("param1"),
					Value: to.Ptr("value1"),
				},
				{
					Name:  to.Ptr("param2"),
					Value: to.Ptr("value2"),
				}},
			RunAsPassword: to.Ptr("<runAsPassword>"),
			RunAsUser:     to.Ptr("user1"),
			Source: &armcompute.VirtualMachineRunCommandScriptSource{
				ScriptURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1"),
				ScriptURIManagedIdentity: &armcompute.RunCommandManagedIdentity{
					ObjectID: to.Ptr("4231e4d2-33e4-4e23-96b2-17888afa6072"),
				},
			},
			TimeoutInSeconds:                to.Ptr[int32](3600),
			TreatFailureAsDeploymentFailure: to.Ptr(true),
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
	// res.VirtualMachineRunCommand = armcompute.VirtualMachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/runCommands"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/runCommands/myRunCommand"),
	// 	Location: to.Ptr("westus"),
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
	// 			ScriptURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 		TreatFailureAsDeploymentFailure: to.Ptr(true),
	// 	},
	// }
}
