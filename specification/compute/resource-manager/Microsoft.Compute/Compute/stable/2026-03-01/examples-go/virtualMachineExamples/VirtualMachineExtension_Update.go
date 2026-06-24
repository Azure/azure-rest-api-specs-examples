package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineExamples/VirtualMachineExtension_Update.json
func ExampleVirtualMachineExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineExtensionsClient().BeginUpdate(ctx, "myResourceGroup", "myVM", "myVMExtension", armcompute.VirtualMachineExtensionUpdate{
		Properties: &armcompute.VirtualMachineExtensionUpdateProperties{
			AutoUpgradeMinorVersion: to.Ptr(true),
			Publisher:               to.Ptr("extPublisher"),
			Type:                    to.Ptr("extType"),
			TypeHandlerVersion:      to.Ptr("1.2"),
			SuppressFailures:        to.Ptr(true),
			Settings: map[string]any{
				"UserName": "xyz@microsoft.com",
			},
			ProtectedSettingsFromKeyVault: &armcompute.KeyVaultSecretReference{
				SourceVault: &armcompute.SubResource{
					ID: to.Ptr("/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName"),
				},
				SecretURL: to.Ptr("https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e"),
			},
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
	// res = armcompute.VirtualMachineExtensionsClientUpdateResponse{
	// 	VirtualMachineExtension: armcompute.VirtualMachineExtension{
	// 		Name: to.Ptr("myVMExtension"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/extensions/myVMExtension"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armcompute.VirtualMachineExtensionProperties{
	// 			AutoUpgradeMinorVersion: to.Ptr(true),
	// 			ProvisioningState: to.Ptr("Creating"),
	// 			Publisher: to.Ptr("extPublisher"),
	// 			Type: to.Ptr("extType"),
	// 			TypeHandlerVersion: to.Ptr("1.2"),
	// 			SuppressFailures: to.Ptr(true),
	// 			Settings: map[string]any{
	// 				"UserName": "xyz@microsoft.com",
	// 			},
	// 			ProtectedSettingsFromKeyVault: &armcompute.KeyVaultSecretReference{
	// 				SourceVault: &armcompute.SubResource{
	// 					ID: to.Ptr("/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName"),
	// 				},
	// 				SecretURL: to.Ptr("https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e"),
	// 			},
	// 		},
	// 	},
	// }
}
