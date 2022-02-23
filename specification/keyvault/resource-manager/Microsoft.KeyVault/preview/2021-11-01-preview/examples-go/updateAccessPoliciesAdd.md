Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkeyvault%2Farmkeyvault%2Fv0.3.1/sdk/resourcemanager/keyvault/armkeyvault/README.md) on how to add the SDK to your project and authenticate.

```go
package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/updateAccessPoliciesAdd.json
func ExampleVaultsClient_UpdateAccessPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkeyvault.NewVaultsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateAccessPolicy(ctx,
		"<resource-group-name>",
		"<vault-name>",
		armkeyvault.AccessPolicyUpdateKindAdd,
		armkeyvault.VaultAccessPolicyParameters{
			Properties: &armkeyvault.VaultAccessPolicyProperties{
				AccessPolicies: []*armkeyvault.AccessPolicyEntry{
					{
						ObjectID: to.StringPtr("<object-id>"),
						Permissions: &armkeyvault.Permissions{
							Certificates: []*armkeyvault.CertificatePermissions{
								armkeyvault.CertificatePermissions("get").ToPtr()},
							Keys: []*armkeyvault.KeyPermissions{
								armkeyvault.KeyPermissions("encrypt").ToPtr()},
							Secrets: []*armkeyvault.SecretPermissions{
								armkeyvault.SecretPermissions("get").ToPtr()},
						},
						TenantID: to.StringPtr("<tenant-id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VaultsClientUpdateAccessPolicyResult)
}
```
