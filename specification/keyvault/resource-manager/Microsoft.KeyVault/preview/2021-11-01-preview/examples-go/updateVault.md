Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkeyvault%2Farmkeyvault%2Fv0.5.0/sdk/resourcemanager/keyvault/armkeyvault/README.md) on how to add the SDK to your project and authenticate.

```go
package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/updateVault.json
func ExampleVaultsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armkeyvault.NewVaultsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<vault-name>",
		armkeyvault.VaultPatchParameters{
			Properties: &armkeyvault.VaultPatchProperties{
				AccessPolicies: []*armkeyvault.AccessPolicyEntry{
					{
						ObjectID: to.Ptr("<object-id>"),
						Permissions: &armkeyvault.Permissions{
							Certificates: []*armkeyvault.CertificatePermissions{
								to.Ptr(armkeyvault.CertificatePermissionsGet),
								to.Ptr(armkeyvault.CertificatePermissionsList),
								to.Ptr(armkeyvault.CertificatePermissionsDelete),
								to.Ptr(armkeyvault.CertificatePermissionsCreate),
								to.Ptr(armkeyvault.CertificatePermissionsImport),
								to.Ptr(armkeyvault.CertificatePermissionsUpdate),
								to.Ptr(armkeyvault.CertificatePermissionsManagecontacts),
								to.Ptr(armkeyvault.CertificatePermissionsGetissuers),
								to.Ptr(armkeyvault.CertificatePermissionsListissuers),
								to.Ptr(armkeyvault.CertificatePermissionsSetissuers),
								to.Ptr(armkeyvault.CertificatePermissionsDeleteissuers),
								to.Ptr(armkeyvault.CertificatePermissionsManageissuers),
								to.Ptr(armkeyvault.CertificatePermissionsRecover),
								to.Ptr(armkeyvault.CertificatePermissionsPurge)},
							Keys: []*armkeyvault.KeyPermissions{
								to.Ptr(armkeyvault.KeyPermissionsEncrypt),
								to.Ptr(armkeyvault.KeyPermissionsDecrypt),
								to.Ptr(armkeyvault.KeyPermissionsWrapKey),
								to.Ptr(armkeyvault.KeyPermissionsUnwrapKey),
								to.Ptr(armkeyvault.KeyPermissionsSign),
								to.Ptr(armkeyvault.KeyPermissionsVerify),
								to.Ptr(armkeyvault.KeyPermissionsGet),
								to.Ptr(armkeyvault.KeyPermissionsList),
								to.Ptr(armkeyvault.KeyPermissionsCreate),
								to.Ptr(armkeyvault.KeyPermissionsUpdate),
								to.Ptr(armkeyvault.KeyPermissionsImport),
								to.Ptr(armkeyvault.KeyPermissionsDelete),
								to.Ptr(armkeyvault.KeyPermissionsBackup),
								to.Ptr(armkeyvault.KeyPermissionsRestore),
								to.Ptr(armkeyvault.KeyPermissionsRecover),
								to.Ptr(armkeyvault.KeyPermissionsPurge)},
							Secrets: []*armkeyvault.SecretPermissions{
								to.Ptr(armkeyvault.SecretPermissionsGet),
								to.Ptr(armkeyvault.SecretPermissionsList),
								to.Ptr(armkeyvault.SecretPermissionsSet),
								to.Ptr(armkeyvault.SecretPermissionsDelete),
								to.Ptr(armkeyvault.SecretPermissionsBackup),
								to.Ptr(armkeyvault.SecretPermissionsRestore),
								to.Ptr(armkeyvault.SecretPermissionsRecover),
								to.Ptr(armkeyvault.SecretPermissionsPurge)},
						},
						TenantID: to.Ptr("<tenant-id>"),
					}},
				EnabledForDeployment:         to.Ptr(true),
				EnabledForDiskEncryption:     to.Ptr(true),
				EnabledForTemplateDeployment: to.Ptr(true),
				PublicNetworkAccess:          to.Ptr("<public-network-access>"),
				SKU: &armkeyvault.SKU{
					Name:   to.Ptr(armkeyvault.SKUNameStandard),
					Family: to.Ptr(armkeyvault.SKUFamilyA),
				},
				TenantID: to.Ptr("<tenant-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
