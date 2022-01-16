Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkeyvault%2Farmkeyvault%2Fv0.3.0/sdk/resourcemanager/keyvault/armkeyvault/README.md) on how to add the SDK to your project and authenticate.

```go
package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/updateVault.json
func ExampleVaultsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkeyvault.NewVaultsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<vault-name>",
		armkeyvault.VaultPatchParameters{
			Properties: &armkeyvault.VaultPatchProperties{
				AccessPolicies: []*armkeyvault.AccessPolicyEntry{
					{
						ObjectID: to.StringPtr("<object-id>"),
						Permissions: &armkeyvault.Permissions{
							Certificates: []*armkeyvault.CertificatePermissions{
								armkeyvault.CertificatePermissions("get").ToPtr(),
								armkeyvault.CertificatePermissions("list").ToPtr(),
								armkeyvault.CertificatePermissions("delete").ToPtr(),
								armkeyvault.CertificatePermissions("create").ToPtr(),
								armkeyvault.CertificatePermissions("import").ToPtr(),
								armkeyvault.CertificatePermissions("update").ToPtr(),
								armkeyvault.CertificatePermissions("managecontacts").ToPtr(),
								armkeyvault.CertificatePermissions("getissuers").ToPtr(),
								armkeyvault.CertificatePermissions("listissuers").ToPtr(),
								armkeyvault.CertificatePermissions("setissuers").ToPtr(),
								armkeyvault.CertificatePermissions("deleteissuers").ToPtr(),
								armkeyvault.CertificatePermissions("manageissuers").ToPtr(),
								armkeyvault.CertificatePermissions("recover").ToPtr(),
								armkeyvault.CertificatePermissions("purge").ToPtr()},
							Keys: []*armkeyvault.KeyPermissions{
								armkeyvault.KeyPermissions("encrypt").ToPtr(),
								armkeyvault.KeyPermissions("decrypt").ToPtr(),
								armkeyvault.KeyPermissions("wrapKey").ToPtr(),
								armkeyvault.KeyPermissions("unwrapKey").ToPtr(),
								armkeyvault.KeyPermissions("sign").ToPtr(),
								armkeyvault.KeyPermissions("verify").ToPtr(),
								armkeyvault.KeyPermissions("get").ToPtr(),
								armkeyvault.KeyPermissions("list").ToPtr(),
								armkeyvault.KeyPermissions("create").ToPtr(),
								armkeyvault.KeyPermissions("update").ToPtr(),
								armkeyvault.KeyPermissions("import").ToPtr(),
								armkeyvault.KeyPermissions("delete").ToPtr(),
								armkeyvault.KeyPermissions("backup").ToPtr(),
								armkeyvault.KeyPermissions("restore").ToPtr(),
								armkeyvault.KeyPermissions("recover").ToPtr(),
								armkeyvault.KeyPermissions("purge").ToPtr()},
							Secrets: []*armkeyvault.SecretPermissions{
								armkeyvault.SecretPermissions("get").ToPtr(),
								armkeyvault.SecretPermissions("list").ToPtr(),
								armkeyvault.SecretPermissions("set").ToPtr(),
								armkeyvault.SecretPermissions("delete").ToPtr(),
								armkeyvault.SecretPermissions("backup").ToPtr(),
								armkeyvault.SecretPermissions("restore").ToPtr(),
								armkeyvault.SecretPermissions("recover").ToPtr(),
								armkeyvault.SecretPermissions("purge").ToPtr()},
						},
						TenantID: to.StringPtr("<tenant-id>"),
					}},
				EnabledForDeployment:         to.BoolPtr(true),
				EnabledForDiskEncryption:     to.BoolPtr(true),
				EnabledForTemplateDeployment: to.BoolPtr(true),
				PublicNetworkAccess:          to.StringPtr("<public-network-access>"),
				SKU: &armkeyvault.SKU{
					Name:   armkeyvault.SKUNameStandard.ToPtr(),
					Family: armkeyvault.SKUFamily("A").ToPtr(),
				},
				TenantID: to.StringPtr("<tenant-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VaultsClientUpdateResult)
}
```
