```go
package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/createVault.json
func ExampleVaultsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkeyvault.NewVaultsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"sample-resource-group",
		"sample-vault",
		armkeyvault.VaultCreateOrUpdateParameters{
			Location: to.Ptr("westus"),
			Properties: &armkeyvault.VaultProperties{
				AccessPolicies: []*armkeyvault.AccessPolicyEntry{
					{
						ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
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
						TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
					}},
				EnabledForDeployment:         to.Ptr(true),
				EnabledForDiskEncryption:     to.Ptr(true),
				EnabledForTemplateDeployment: to.Ptr(true),
				PublicNetworkAccess:          to.Ptr("Enabled"),
				SKU: &armkeyvault.SKU{
					Name:   to.Ptr(armkeyvault.SKUNameStandard),
					Family: to.Ptr(armkeyvault.SKUFamilyA),
				},
				TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
			},
		},
		nil)
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkeyvault%2Farmkeyvault%2Fv1.0.0/sdk/resourcemanager/keyvault/armkeyvault/README.md) on how to add the SDK to your project and authenticate.
