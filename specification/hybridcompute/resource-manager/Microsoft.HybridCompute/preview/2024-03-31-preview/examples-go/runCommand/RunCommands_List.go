package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/runCommand/RunCommands_List.json
func ExampleMachineRunCommandsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMachineRunCommandsClient().NewListPager("myResourceGroup", "myMachine", &armhybridcompute.MachineRunCommandsClientListOptions{Expand: nil})
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
		// page.MachineRunCommandsListResult = armhybridcompute.MachineRunCommandsListResult{
		// 	Value: []*armhybridcompute.MachineRunCommand{
		// 		{
		// 			Name: to.Ptr("myRunCommand_1"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/machines/runcommands"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/runcommands/myRunCommand_1"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armhybridcompute.MachineRunCommandProperties{
		// 				AsyncExecution: to.Ptr(false),
		// 				InstanceView: &armhybridcompute.MachineRunCommandInstanceView{
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
		// 					Error: to.Ptr(""),
		// 					ExecutionMessage: to.Ptr(""),
		// 					ExecutionState: to.Ptr(armhybridcompute.ExecutionStateSucceeded),
		// 					ExitCode: to.Ptr[int32](0),
		// 					Output: to.Ptr("Hello World"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
		// 				},
		// 				Parameters: []*armhybridcompute.RunCommandInputParameter{
		// 					{
		// 						Name: to.Ptr("param1"),
		// 						Value: to.Ptr("value1"),
		// 					},
		// 					{
		// 						Name: to.Ptr("param2"),
		// 						Value: to.Ptr("value2"),
		// 				}},
		// 				ProtectedParameters: []*armhybridcompute.RunCommandInputParameter{
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RunAsUser: to.Ptr("user1"),
		// 				Source: &armhybridcompute.MachineRunCommandScriptSource{
		// 					Script: to.Ptr("Write-Host Hello World!"),
		// 				},
		// 				TimeoutInSeconds: to.Ptr[int32](3600),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myRunCommand_2"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/machines/runcommands"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/runcommands/myRunCommand_2"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armhybridcompute.MachineRunCommandProperties{
		// 				AsyncExecution: to.Ptr(false),
		// 				InstanceView: &armhybridcompute.MachineRunCommandInstanceView{
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
		// 					Error: to.Ptr(""),
		// 					ExecutionMessage: to.Ptr(""),
		// 					ExecutionState: to.Ptr(armhybridcompute.ExecutionStateSucceeded),
		// 					ExitCode: to.Ptr[int32](0),
		// 					Output: to.Ptr("<some output>"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
		// 				},
		// 				Parameters: []*armhybridcompute.RunCommandInputParameter{
		// 				},
		// 				ProtectedParameters: []*armhybridcompute.RunCommandInputParameter{
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RunAsUser: to.Ptr("userA"),
		// 				Source: &armhybridcompute.MachineRunCommandScriptSource{
		// 					Script: to.Ptr("Get-Process | Where-Object { $_.CPU -gt 10000 }"),
		// 				},
		// 				TimeoutInSeconds: to.Ptr[int32](100),
		// 			},
		// 	}},
		// }
	}
}
