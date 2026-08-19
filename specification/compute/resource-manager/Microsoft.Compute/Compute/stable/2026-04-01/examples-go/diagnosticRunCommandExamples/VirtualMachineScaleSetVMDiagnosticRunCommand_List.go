package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-04-01/diagnosticRunCommandExamples/VirtualMachineScaleSetVMDiagnosticRunCommand_List.json
func ExampleVirtualMachineScaleSetVMDiagnosticRunCommandsClient_NewDiagnosticListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineScaleSetVMDiagnosticRunCommandsClient().NewDiagnosticListPager("myResourceGroup", "myvmScaleSet", "0", nil)
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
		// page = armcompute.VirtualMachineScaleSetVMDiagnosticRunCommandsClientDiagnosticListResponse{
		// 	VirtualMachineDiagnosticRunCommandsListResult: armcompute.VirtualMachineDiagnosticRunCommandsListResult{
		// 		Value: []*armcompute.VirtualMachineDiagnosticRunCommand{
		// 			{
		// 				Name: to.Ptr("myRunCommand"),
		// 				Location: to.Ptr("westus"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/runCommands"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/runCommands/myRunCommand"),
		// 				Properties: &armcompute.VirtualMachineRunCommandProperties{
		// 					Source: &armcompute.VirtualMachineRunCommandScriptSource{
		// 						CommandID: to.Ptr("FleetDiagnosticsWindows"),
		// 					},
		// 					Parameters: []*armcompute.RunCommandInputParameter{
		// 						{
		// 							Name: to.Ptr("param1"),
		// 							Value: to.Ptr("value1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("param2"),
		// 							Value: to.Ptr("value2"),
		// 						},
		// 					},
		// 					AsyncExecution: to.Ptr(false),
		// 					TreatFailureAsDeploymentFailure: to.Ptr(false),
		// 					TimeoutInSeconds: to.Ptr[int32](0),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
