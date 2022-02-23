Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkeyvault%2Farmkeyvault%2Fv0.3.1/sdk/resourcemanager/keyvault/armkeyvault/README.md) on how to add the SDK to your project and authenticate.

```go
package armkeyvault_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/ManagedHsm_CreateOrUpdate.json
func ExampleManagedHsmsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkeyvault.NewManagedHsmsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<name>",
		armkeyvault.ManagedHsm{
			Location: to.StringPtr("<location>"),
			SKU: &armkeyvault.ManagedHsmSKU{
				Name:   armkeyvault.ManagedHsmSKUNameStandardB1.ToPtr(),
				Family: armkeyvault.ManagedHsmSKUFamily("B").ToPtr(),
			},
			Tags: map[string]*string{
				"Dept":        to.StringPtr("hsm"),
				"Environment": to.StringPtr("dogfood"),
			},
			Properties: &armkeyvault.ManagedHsmProperties{
				EnablePurgeProtection: to.BoolPtr(true),
				EnableSoftDelete:      to.BoolPtr(true),
				InitialAdminObjectIDs: []*string{
					to.StringPtr("00000000-0000-0000-0000-000000000000")},
				SoftDeleteRetentionInDays: to.Int32Ptr(90),
				TenantID:                  to.StringPtr("<tenant-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ManagedHsmsClientCreateOrUpdateResult)
}
```
