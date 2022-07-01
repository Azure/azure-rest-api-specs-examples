package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/runCommandExamples/VirtualMachineRunCommand_CreateOrUpdate.json
func ExampleVirtualMachineRunCommandsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineRunCommandsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		"myRunCommand",
		armcompute.VirtualMachineRunCommand{
			Location: to.Ptr("West US"),
			Properties: &armcompute.VirtualMachineRunCommandProperties{
				AsyncExecution: to.Ptr(false),
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
					Script: to.Ptr("Write-Host Hello World!"),
				},
				TimeoutInSeconds: to.Ptr[int32](3600),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
