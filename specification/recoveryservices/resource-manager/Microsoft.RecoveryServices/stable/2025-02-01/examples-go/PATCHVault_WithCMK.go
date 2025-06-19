package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/PATCHVault_WithCMK.json
func ExampleVaultsClient_BeginUpdate_updateResourceWithCustomerManagedKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVaultsClient().BeginUpdate(ctx, "HelloWorld", "swaggerExample", armrecoveryservices.PatchVault{
		Tags: map[string]*string{
			"PatchKey": to.Ptr("PatchKeyUpdated"),
		},
		Identity: &armrecoveryservices.IdentityData{
			Type: to.Ptr(armrecoveryservices.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armrecoveryservices.UserIdentity{
				"/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi": {},
			},
		},
		Properties: &armrecoveryservices.VaultProperties{
			Encryption: &armrecoveryservices.VaultPropertiesEncryption{
				InfrastructureEncryption: to.Ptr(armrecoveryservices.InfrastructureEncryptionStateEnabled),
				KekIdentity: &armrecoveryservices.CmkKekIdentity{
					UserAssignedIdentity: to.Ptr("/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi"),
				},
				KeyVaultProperties: &armrecoveryservices.CmkKeyVaultProperties{
					KeyURI: to.Ptr("https://cmk2xkv.vault.azure.net/keys/Key1/0767b348bb1a4c07baa6c4ec0055d2b3"),
				},
			},
		},
	}, &armrecoveryservices.VaultsClientBeginUpdateOptions{XMSAuthorizationAuxiliary: nil})
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
	// res.Vault = armrecoveryservices.Vault{
	// 	Name: to.Ptr("swaggerExample"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults"),
	// 	Etag: to.Ptr("W/\"datetime'2017-12-15T12%3A36%3A51.68Z'\""),
	// 	ID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/resourceGroups/HelloWorld/providers/Microsoft.RecoveryServices/vaults/swaggerExample"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"PatchKey": to.Ptr("PatchKeyUpdated"),
	// 	},
	// 	Identity: &armrecoveryservices.IdentityData{
	// 		Type: to.Ptr(armrecoveryservices.ResourceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armrecoveryservices.UserIdentity{
	// 			"/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi": &armrecoveryservices.UserIdentity{
	// 				ClientID: to.Ptr("fbe75b66-01c5-4f87-a220-233af3270436"),
	// 				PrincipalID: to.Ptr("075a0ca6-43f6-4434-9abf-c9b1b79f9219"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armrecoveryservices.VaultProperties{
	// 		Encryption: &armrecoveryservices.VaultPropertiesEncryption{
	// 			InfrastructureEncryption: to.Ptr(armrecoveryservices.InfrastructureEncryptionStateEnabled),
	// 			KekIdentity: &armrecoveryservices.CmkKekIdentity{
	// 				UseSystemAssignedIdentity: to.Ptr(false),
	// 				UserAssignedIdentity: to.Ptr("/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi"),
	// 			},
	// 			KeyVaultProperties: &armrecoveryservices.CmkKeyVaultProperties{
	// 				KeyURI: to.Ptr("https://cmk2xkv.vault.azure.net/keys/Key1/0767b348bb1a4c07baa6c4ec0055d2b3"),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicNetworkAccess: to.Ptr(armrecoveryservices.PublicNetworkAccessEnabled),
	// 	},
	// 	SKU: &armrecoveryservices.SKU{
	// 		Name: to.Ptr(armrecoveryservices.SKUNameStandard),
	// 	},
	// }
}
