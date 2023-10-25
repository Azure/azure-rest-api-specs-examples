package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/listVaultBySubscription.json
func ExampleVaultsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVaultsClient().NewListBySubscriptionPager(&armkeyvault.VaultsClientListBySubscriptionOptions{Top: to.Ptr[int32](1)})
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
		// page.VaultListResult = armkeyvault.VaultListResult{
		// 	Value: []*armkeyvault.Vault{
		// 		{
		// 			Name: to.Ptr("sample-vault"),
		// 			Type: to.Ptr("Microsoft.KeyVault/vaults"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkeyvault.VaultProperties{
		// 				AccessPolicies: []*armkeyvault.AccessPolicyEntry{
		// 					{
		// 						ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						Permissions: &armkeyvault.Permissions{
		// 							Certificates: []*armkeyvault.CertificatePermissions{
		// 								to.Ptr(armkeyvault.CertificatePermissionsGet),
		// 								to.Ptr(armkeyvault.CertificatePermissionsList),
		// 								to.Ptr(armkeyvault.CertificatePermissionsDelete),
		// 								to.Ptr(armkeyvault.CertificatePermissionsCreate),
		// 								to.Ptr(armkeyvault.CertificatePermissionsImport),
		// 								to.Ptr(armkeyvault.CertificatePermissionsUpdate),
		// 								to.Ptr(armkeyvault.CertificatePermissionsManagecontacts),
		// 								to.Ptr(armkeyvault.CertificatePermissionsGetissuers),
		// 								to.Ptr(armkeyvault.CertificatePermissionsListissuers),
		// 								to.Ptr(armkeyvault.CertificatePermissionsSetissuers),
		// 								to.Ptr(armkeyvault.CertificatePermissionsDeleteissuers),
		// 								to.Ptr(armkeyvault.CertificatePermissionsManageissuers),
		// 								to.Ptr(armkeyvault.CertificatePermissionsRecover),
		// 								to.Ptr(armkeyvault.CertificatePermissionsPurge)},
		// 								Keys: []*armkeyvault.KeyPermissions{
		// 									to.Ptr(armkeyvault.KeyPermissionsEncrypt),
		// 									to.Ptr(armkeyvault.KeyPermissionsDecrypt),
		// 									to.Ptr(armkeyvault.KeyPermissionsWrapKey),
		// 									to.Ptr(armkeyvault.KeyPermissionsUnwrapKey),
		// 									to.Ptr(armkeyvault.KeyPermissionsSign),
		// 									to.Ptr(armkeyvault.KeyPermissionsVerify),
		// 									to.Ptr(armkeyvault.KeyPermissionsGet),
		// 									to.Ptr(armkeyvault.KeyPermissionsList),
		// 									to.Ptr(armkeyvault.KeyPermissionsCreate),
		// 									to.Ptr(armkeyvault.KeyPermissionsUpdate),
		// 									to.Ptr(armkeyvault.KeyPermissionsImport),
		// 									to.Ptr(armkeyvault.KeyPermissionsDelete),
		// 									to.Ptr(armkeyvault.KeyPermissionsBackup),
		// 									to.Ptr(armkeyvault.KeyPermissionsRestore),
		// 									to.Ptr(armkeyvault.KeyPermissionsRecover),
		// 									to.Ptr(armkeyvault.KeyPermissionsPurge)},
		// 									Secrets: []*armkeyvault.SecretPermissions{
		// 										to.Ptr(armkeyvault.SecretPermissionsGet),
		// 										to.Ptr(armkeyvault.SecretPermissionsList),
		// 										to.Ptr(armkeyvault.SecretPermissionsSet),
		// 										to.Ptr(armkeyvault.SecretPermissionsDelete),
		// 										to.Ptr(armkeyvault.SecretPermissionsBackup),
		// 										to.Ptr(armkeyvault.SecretPermissionsRestore),
		// 										to.Ptr(armkeyvault.SecretPermissionsRecover),
		// 										to.Ptr(armkeyvault.SecretPermissionsPurge)},
		// 									},
		// 									TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							}},
		// 							EnableSoftDelete: to.Ptr(true),
		// 							EnabledForDeployment: to.Ptr(true),
		// 							EnabledForDiskEncryption: to.Ptr(true),
		// 							EnabledForTemplateDeployment: to.Ptr(true),
		// 							HsmPoolResourceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							ProvisioningState: to.Ptr(armkeyvault.VaultProvisioningStateSucceeded),
		// 							SKU: &armkeyvault.SKU{
		// 								Name: to.Ptr(armkeyvault.SKUNamePremium),
		// 								Family: to.Ptr(armkeyvault.SKUFamilyA),
		// 							},
		// 							TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							VaultURI: to.Ptr("https://sample-vault.vault.azure.net/"),
		// 						},
		// 						SystemData: &armkeyvault.SystemData{
		// 							CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.0000000Z"); return t}()),
		// 							CreatedBy: to.Ptr("keyVaultUser1"),
		// 							CreatedByType: to.Ptr(armkeyvault.IdentityTypeUser),
		// 							LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.0000000Z"); return t}()),
		// 							LastModifiedBy: to.Ptr("keyVaultUser2"),
		// 							LastModifiedByType: to.Ptr(armkeyvault.IdentityTypeUser),
		// 						},
		// 						Tags: map[string]*string{
		// 						},
		// 				}},
		// 			}
	}
}
