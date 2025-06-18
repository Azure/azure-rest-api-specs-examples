package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/extension/Extension_Update.json
func ExampleMachineExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMachineExtensionsClient().BeginUpdate(ctx, "myResourceGroup", "myMachine", "CustomScriptExtension", armhybridcompute.MachineExtensionUpdate{
		Properties: &armhybridcompute.MachineExtensionUpdateProperties{
			Type:                   to.Ptr("CustomScriptExtension"),
			EnableAutomaticUpgrade: to.Ptr(true),
			Publisher:              to.Ptr("Microsoft.Compute"),
			Settings: map[string]any{
				"commandToExecute": "powershell.exe -c \"Get-Process | Where-Object { $_.CPU -lt 100 }\"",
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
	// 		EnableAutomaticUpgrade: to.Ptr(true),
	// 		InstanceView: &armhybridcompute.MachineExtensionInstanceView{
	// 			Name: to.Ptr("CustomScriptExtension"),
	// 			Type: to.Ptr("CustomScriptExtension"),
	// 			Status: &armhybridcompute.MachineExtensionInstanceViewStatus{
	// 				Code: to.Ptr("success"),
	// 				Level: to.Ptr(armhybridcompute.StatusLevelTypes("Information")),
	// 				Message: to.Ptr("Finished executing command, StdOut: , StdErr:"),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-08T20:42:10.999Z"); return t}()),
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
