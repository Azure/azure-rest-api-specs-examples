package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GETExtension.json
func ExampleMachineExtensionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMachineExtensionsClient().Get(ctx, "myResourceGroup", "myMachine", "CustomScriptExtension", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MachineExtension = armconnectedvmware.MachineExtension{
	// 	Name: to.Ptr("CustomScriptExtension"),
	// 	Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VirtualMachines/extensions"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/Extensions/CustomScriptExtension"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armconnectedvmware.MachineExtensionProperties{
	// 		Type: to.Ptr("string"),
	// 		AutoUpgradeMinorVersion: to.Ptr(false),
	// 		InstanceView: &armconnectedvmware.MachineExtensionPropertiesInstanceView{
	// 			Name: to.Ptr("CustomScriptExtension"),
	// 			Type: to.Ptr("CustomScriptExtension"),
	// 			Status: &armconnectedvmware.MachineExtensionInstanceViewStatus{
	// 				Code: to.Ptr("success"),
	// 				DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 				Level: to.Ptr(armconnectedvmware.StatusLevelTypes("Information")),
	// 				Message: to.Ptr("Finished executing command, StdOut: , StdErr:"),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-08T20:42:10.999Z"); return t}()),
	// 			},
	// 			TypeHandlerVersion: to.Ptr("1.10.3"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Publisher: to.Ptr("Microsoft.Compute"),
	// 		Settings: "@{commandToExecute=powershell.exe -c \"Get-Process | Where-Object { $_.CPU -gt 10000 }\"}",
	// 		TypeHandlerVersion: to.Ptr("1.10.3"),
	// 	},
	// }
}
