package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-04-01/diagnosticRunCommandExamples/VirtualMachineScaleSetVMDiagnosticRunCommand_CreateOrUpdate.json
func ExampleVirtualMachineScaleSetVMDiagnosticRunCommandsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMDiagnosticRunCommandsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myvmScaleSet", "0", "myRunCommand", armcompute.VirtualMachineDiagnosticRunCommand{
		Location: to.Ptr("West US"),
		Properties: &armcompute.VirtualMachineRunCommandProperties{
			Source: &armcompute.VirtualMachineRunCommandScriptSource{
				CommandID: to.Ptr("FleetDiagnosticsWindows"),
			},
			Parameters: []*armcompute.RunCommandInputParameter{
				{
					Name:  to.Ptr("param1"),
					Value: to.Ptr("value1"),
				},
				{
					Name:  to.Ptr("param2"),
					Value: to.Ptr("value2"),
				},
			},
			AsyncExecution:                  to.Ptr(false),
			TreatFailureAsDeploymentFailure: to.Ptr(true),
			RunAsPassword:                   to.Ptr("<runAsPassword>"),
			TimeoutInSeconds:                to.Ptr[int32](3600),
		},
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
	// res = armcompute.VirtualMachineScaleSetVMDiagnosticRunCommandsClientCreateOrUpdateResponse{
	// }
}
