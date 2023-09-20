package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/runCommandExamples/VirtualMachineRunCommand.json
func ExampleVirtualMachinesClient_BeginRunCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RunCommandResult = armcompute.RunCommandResult{
	// 	Value: []*armcompute.InstanceViewStatus{
	// 		{
	// 			Code: to.Ptr("ComponentStatus/StdOut/succeeded"),
	// 			DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 			Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			Message: to.Ptr("This is a sample script with parameters value1 value2"),
	// 		},
	// 		{
	// 			Code: to.Ptr("ComponentStatus/StdErr/succeeded"),
	// 			DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 			Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			Message: to.Ptr(""),
	// 	}},
	// }
}
