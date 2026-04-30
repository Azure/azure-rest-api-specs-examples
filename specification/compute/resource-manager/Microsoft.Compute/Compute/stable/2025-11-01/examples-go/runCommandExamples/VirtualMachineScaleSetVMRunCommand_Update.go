package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/runCommandExamples/VirtualMachineScaleSetVMRunCommand_Update.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMRunCommandsClient().BeginUpdate(ctx, "myResourceGroup", "myvmScaleSet", "0", "myRunCommand", armcompute.VirtualMachineRunCommandUpdate{
		Properties: &armcompute.VirtualMachineRunCommandProperties{
			Source: &armcompute.VirtualMachineRunCommandScriptSource{
				ScriptURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1"),
				ScriptURIManagedIdentity: &armcompute.RunCommandManagedIdentity{
					ObjectID: to.Ptr("4231e4d2-33e4-4e23-96b2-17888afa6072"),
				},
			},
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
	// res = armcompute.VirtualMachineScaleSetVMRunCommandsClientUpdateResponse{
	// 	VirtualMachineRunCommand: &armcompute.VirtualMachineRunCommand{
	// 		Name: to.Ptr("myRunCommand"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/runCommands"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/runCommands/myRunCommand"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 		Properties: &armcompute.VirtualMachineRunCommandProperties{
	// 			Source: &armcompute.VirtualMachineRunCommandScriptSource{
	// 				ScriptURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1"),
	// 			},
	// 			Parameters: []*armcompute.RunCommandInputParameter{
	// 				{
	// 					Name: to.Ptr("param1"),
	// 					Value: to.Ptr("value1"),
	// 				},
	// 				{
	// 					Name: to.Ptr("param2"),
	// 					Value: to.Ptr("value2"),
	// 				},
	// 			},
	// 			AsyncExecution: to.Ptr(false),
	// 			TreatFailureAsDeploymentFailure: to.Ptr(false),
	// 			RunAsUser: to.Ptr("user1"),
	// 			TimeoutInSeconds: to.Ptr[int32](3600),
	// 			ProvisioningState: to.Ptr("Updating"),
	// 			OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
	// 			ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
	// 		},
	// 	},
	// }
}
