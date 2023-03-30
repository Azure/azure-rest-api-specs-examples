package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/LISTExtension.json
func ExampleMachineExtensionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMachineExtensionsClient().NewListPager("myResourceGroup", "myMachine", &armconnectedvmware.MachineExtensionsClientListOptions{Expand: nil})
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
		// page.MachineExtensionsListResult = armconnectedvmware.MachineExtensionsListResult{
		// 	Value: []*armconnectedvmware.MachineExtension{
		// 		{
		// 			Name: to.Ptr("CustomScriptExtension"),
		// 			Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VirtualMachines/extensions"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/Extensions/CustomScriptExtension"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armconnectedvmware.MachineExtensionProperties{
		// 				Type: to.Ptr("CustomScriptExtension"),
		// 				AutoUpgradeMinorVersion: to.Ptr(false),
		// 				InstanceView: &armconnectedvmware.MachineExtensionPropertiesInstanceView{
		// 					Name: to.Ptr("CustomScriptExtension"),
		// 					Type: to.Ptr("CustomScriptExtension"),
		// 					Status: &armconnectedvmware.MachineExtensionInstanceViewStatus{
		// 						Code: to.Ptr("success"),
		// 						DisplayStatus: to.Ptr("Provisioning succeeded"),
		// 						Level: to.Ptr(armconnectedvmware.StatusLevelTypes("Information")),
		// 						Message: to.Ptr("formattedMessage: Finished executing command, StdOut: , StdErr: "),
		// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-13T17:18:57.405Z"); return t}()),
		// 					},
		// 					TypeHandlerVersion: to.Ptr("1.10.3"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Publisher: to.Ptr("Microsoft.Compute"),
		// 				Settings: map[string]any{
		// 					"commandToExecute": "powershell.exe -c \"Get-Process | Where-Object { $_.CPU -gt 10000 }\"",
		// 				},
		// 				TypeHandlerVersion: to.Ptr("1.10.3"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("winosupdateextension"),
		// 			Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VirtualMachines/extensions"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/Extensions/winosupdateextension"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armconnectedvmware.MachineExtensionProperties{
		// 				Type: to.Ptr("windowsosupdateextension"),
		// 				AutoUpgradeMinorVersion: to.Ptr(false),
		// 				InstanceView: &armconnectedvmware.MachineExtensionPropertiesInstanceView{
		// 					Name: to.Ptr("winosupdateextension"),
		// 					Type: to.Ptr("windowsosupdateextension"),
		// 					Status: &armconnectedvmware.MachineExtensionInstanceViewStatus{
		// 					},
		// 					TypeHandlerVersion: to.Ptr("1.0.0.0"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Publisher: to.Ptr("microsoft.softwareupdatemanagement.test"),
		// 				Settings: map[string]any{
		// 				},
		// 				TypeHandlerVersion: to.Ptr("1.0.0.0"),
		// 			},
		// 	}},
		// }
	}
}
