Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.3.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

```go
package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountSetBlobInventoryPolicy.json
func ExampleBlobInventoryPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobInventoryPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.BlobInventoryPolicyName("default"),
		armstorage.BlobInventoryPolicy{
			Properties: &armstorage.BlobInventoryPolicyProperties{
				Policy: &armstorage.BlobInventoryPolicySchema{
					Type:    armstorage.InventoryRuleType("Inventory").ToPtr(),
					Enabled: to.BoolPtr(true),
					Rules: []*armstorage.BlobInventoryPolicyRule{
						{
							Name: to.StringPtr("<name>"),
							Definition: &armstorage.BlobInventoryPolicyDefinition{
								Format: armstorage.Format("Csv").ToPtr(),
								Filters: &armstorage.BlobInventoryPolicyFilter{
									BlobTypes: []*string{
										to.StringPtr("blockBlob"),
										to.StringPtr("appendBlob"),
										to.StringPtr("pageBlob")},
									IncludeBlobVersions: to.BoolPtr(true),
									IncludeSnapshots:    to.BoolPtr(true),
									PrefixMatch: []*string{
										to.StringPtr("inventoryprefix1"),
										to.StringPtr("inventoryprefix2")},
								},
								ObjectType: armstorage.ObjectType("Blob").ToPtr(),
								Schedule:   armstorage.Schedule("Daily").ToPtr(),
								SchemaFields: []*string{
									to.StringPtr("Name"),
									to.StringPtr("Creation-Time"),
									to.StringPtr("Last-Modified"),
									to.StringPtr("Content-Length"),
									to.StringPtr("Content-MD5"),
									to.StringPtr("BlobType"),
									to.StringPtr("AccessTier"),
									to.StringPtr("AccessTierChangeTime"),
									to.StringPtr("Snapshot"),
									to.StringPtr("VersionId"),
									to.StringPtr("IsCurrentVersion"),
									to.StringPtr("Metadata")},
							},
							Destination: to.StringPtr("<destination>"),
							Enabled:     to.BoolPtr(true),
						},
						{
							Name: to.StringPtr("<name>"),
							Definition: &armstorage.BlobInventoryPolicyDefinition{
								Format:     armstorage.Format("Parquet").ToPtr(),
								ObjectType: armstorage.ObjectType("Container").ToPtr(),
								Schedule:   armstorage.Schedule("Weekly").ToPtr(),
								SchemaFields: []*string{
									to.StringPtr("Name"),
									to.StringPtr("Last-Modified"),
									to.StringPtr("Metadata"),
									to.StringPtr("LeaseStatus"),
									to.StringPtr("LeaseState"),
									to.StringPtr("LeaseDuration"),
									to.StringPtr("PublicAccess"),
									to.StringPtr("HasImmutabilityPolicy"),
									to.StringPtr("HasLegalHold")},
							},
							Destination: to.StringPtr("<destination>"),
							Enabled:     to.BoolPtr(true),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobInventoryPoliciesClientCreateOrUpdateResult)
}
```
