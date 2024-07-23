package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/extension/Extension_CreateOrUpdate.json
func ExampleMachineExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMachineExtensionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myMachine", "CustomScriptExtension", armhybridcompute.MachineExtension{
		Location: to.Ptr("eastus2euap"),
		Properties: &armhybridcompute.MachineExtensionProperties{
			Type:      to.Ptr("CustomScriptExtension"),
			Publisher: to.Ptr("Microsoft.Compute"),
			Settings: map[string]any{
				"commandToExecute": "powershell.exe -c \"Get-Process | Where-Object { $_.CPU -gt 10000 }\"",
			},
			TypeHandlerVersion: to.Ptr("1.10"),
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
	// res.MachineExtension = armhybridcompute.MachineExtension{
	// 	Name: to.Ptr("CustomScriptExtension"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/machines/extensions"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/Extensions/CustomScriptExtension"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armhybridcompute.MachineExtensionProperties{
	// 		Type: to.Ptr("string"),
	// 		AutoUpgradeMinorVersion: to.Ptr(false),
	// 		InstanceView: &armhybridcompute.MachineExtensionInstanceView{
	// 			Name: to.Ptr("CustomScriptExtension"),
	// 			Type: to.Ptr("CustomScriptExtension"),
	// 			Status: &armhybridcompute.MachineExtensionInstanceViewStatus{
	// 				Code: to.Ptr("success"),
	// 				Level: to.Ptr(armhybridcompute.StatusLevelTypes("Information")),
	// 				Message: to.Ptr("Finished executing command, StdOut: , StdErr:"),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-08T20:42:10.999Z"); return t}()),
	// 			},
	// 			TypeHandlerVersion: to.Ptr("1.10.3"),
	// 		},
	// 		ProtectedSettings: map[string]any{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Publisher: to.Ptr("Microsoft.Compute"),
	// 		Settings: map[string]any{
	// 			"commandToExecute": "powershell.exe -c \"Get-Process | Where-Object { $_.CPU -gt 10000 }\"",
	// 		},
	// 		TypeHandlerVersion: to.Ptr("1.10.3"),
	// 	},
	// }
}
