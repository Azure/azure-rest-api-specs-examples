package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71a0c7adf2a6e169ab9a33c7cf36bb93db083e86/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-10-03-preview/examples/runCommand/RunCommands_Get.json
func ExampleMachineRunCommandsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMachineRunCommandsClient().Get(ctx, "myResourceGroup", "myMachine", "myRunCommand", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MachineRunCommand = armhybridcompute.MachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/machines/runcommands"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/runcommands/myRunCommand"),
	// 	Location: to.Ptr("eastus2"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armhybridcompute.MachineRunCommandProperties{
	// 		AsyncExecution: to.Ptr(false),
	// 		InstanceView: &armhybridcompute.MachineRunCommandInstanceView{
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
	// 			Error: to.Ptr(""),
	// 			ExecutionMessage: to.Ptr(""),
	// 			ExecutionState: to.Ptr(armhybridcompute.ExecutionStateSucceeded),
	// 			ExitCode: to.Ptr[int32](0),
	// 			Output: to.Ptr("Hello World"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
	// 		},
	// 		Parameters: []*armhybridcompute.RunCommandInputParameter{
	// 			{
	// 				Name: to.Ptr("param1"),
	// 				Value: to.Ptr("value1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("param2"),
	// 				Value: to.Ptr("value2"),
	// 		}},
	// 		ProtectedParameters: []*armhybridcompute.RunCommandInputParameter{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RunAsUser: to.Ptr("user1"),
	// 		Source: &armhybridcompute.MachineRunCommandScriptSource{
	// 			Script: to.Ptr("Write-Host Hello World!"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 	},
	// }
}
