package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/UpdateVMExtension.json
func ExampleVirtualMachineExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineExtensionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-name>",
		"<vm-extension-name>",
		armcompute.VirtualMachineExtensionUpdate{
			Properties: &armcompute.VirtualMachineExtensionUpdateProperties{
				Type:                    to.Ptr("<type>"),
				AutoUpgradeMinorVersion: to.Ptr(true),
				ProtectedSettingsFromKeyVault: map[string]interface{}{
					"secretUrl": "https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e",
					"sourceVault": map[string]interface{}{
						"id": "/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName",
					},
				},
				Publisher: to.Ptr("<publisher>"),
				Settings: map[string]interface{}{
					"UserName": "xyz@microsoft.com",
				},
				SuppressFailures:   to.Ptr(true),
				TypeHandlerVersion: to.Ptr("<type-handler-version>"),
			},
		},
		&armcompute.VirtualMachineExtensionsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
