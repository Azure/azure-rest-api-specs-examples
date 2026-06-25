package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/runCommandExamples/VirtualMachineRunCommand.json
func ExampleVirtualMachinesClient_BeginRunCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("24fb23e3-6ba3-41f0-9b6e-e41131d5d61e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginRunCommand(ctx, "crptestar98131", "vm3036", armcompute.RunCommandInput{
		CommandID: to.Ptr("RunPowerShellScript"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachinesClientRunCommandResponse{
	// 	RunCommandResult: armcompute.RunCommandResult{
	// 		Value: []*armcompute.InstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("ComponentStatus/StdOut/succeeded"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 				Message: to.Ptr("This is a sample script with parameters value1 value2"),
	// 			},
	// 			{
	// 				Code: to.Ptr("ComponentStatus/StdErr/succeeded"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 				Message: to.Ptr(""),
	// 			},
	// 		},
	// 	},
	// }
}
