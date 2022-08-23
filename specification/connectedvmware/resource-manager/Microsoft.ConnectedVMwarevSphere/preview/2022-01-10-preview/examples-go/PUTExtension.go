package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/PUTExtension.json
func ExampleMachineExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconnectedvmware.NewMachineExtensionsClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myMachine", "CustomScriptExtension", armconnectedvmware.MachineExtension{
		Location: to.Ptr("eastus2euap"),
		Properties: &armconnectedvmware.MachineExtensionProperties{
			Type:      to.Ptr("CustomScriptExtension"),
			Publisher: to.Ptr("Microsoft.Compute"),
			Settings: map[string]interface{}{
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
	// TODO: use response item
	_ = res
}
