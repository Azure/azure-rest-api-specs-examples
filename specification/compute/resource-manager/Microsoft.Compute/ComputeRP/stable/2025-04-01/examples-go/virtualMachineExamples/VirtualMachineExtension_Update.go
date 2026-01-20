package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineExamples/VirtualMachineExtension_Update.json
func ExampleVirtualMachineExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineExtensionsClient().BeginUpdate(ctx, "myResourceGroup", "myVM", "myVMExtension", armcompute.VirtualMachineExtensionUpdate{
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
			Settings: map[string]any{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineExtension = armcompute.VirtualMachineExtension{
	// 	Name: to.Ptr("myVMExtension"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/extensions/myVMExtension"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcompute.VirtualMachineExtensionProperties{
	// 		Type: to.Ptr("extType"),
	// 		AutoUpgradeMinorVersion: to.Ptr(true),
	// 		ProtectedSettingsFromKeyVault: &armcompute.KeyVaultSecretReference{
	// 			SecretURL: to.Ptr("https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e"),
	// 			SourceVault: &armcompute.SubResource{
	// 				ID: to.Ptr("/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName"),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Publisher: to.Ptr("extPublisher"),
	// 		Settings: map[string]any{
	// 			"UserName": "xyz@microsoft.com",
	// 		},
	// 		SuppressFailures: to.Ptr(true),
	// 		TypeHandlerVersion: to.Ptr("1.2"),
	// 	},
	// }
}
