package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachineExtensions_Update.json
func ExampleVirtualMachineExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineExtensionsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "myResourceGroup", "myVM", "myVMExtension", armcompute.VirtualMachineExtensionUpdate{
		Properties: &armcompute.VirtualMachineExtensionUpdateProperties{
			Type:                    to.Ptr("extType"),
			AutoUpgradeMinorVersion: to.Ptr(true),
			ProtectedSettingsFromKeyVault: &armcompute.KeyVaultSecretReference{
				SecretURL: to.Ptr("https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e"),
				SourceVault: &armcompute.SubResource{
					ID: to.Ptr("/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName"),
				},
			},
			Publisher: to.Ptr("extPublisher"),
			Settings: map[string]interface{}{
				"UserName": "xyz@microsoft.com",
			},
			SuppressFailures:   to.Ptr(true),
			TypeHandlerVersion: to.Ptr("1.2"),
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
