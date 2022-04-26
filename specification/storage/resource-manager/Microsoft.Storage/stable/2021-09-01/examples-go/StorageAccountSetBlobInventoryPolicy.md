Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.6.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

```go
package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountSetBlobInventoryPolicy.json
func ExampleBlobInventoryPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobInventoryPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.BlobInventoryPolicyNameDefault,
		armstorage.BlobInventoryPolicy{
			Properties: &armstorage.BlobInventoryPolicyProperties{
				Policy: &armstorage.BlobInventoryPolicySchema{
					Type:    to.Ptr(armstorage.InventoryRuleTypeInventory),
					Enabled: to.Ptr(true),
					Rules: []*armstorage.BlobInventoryPolicyRule{
						{
							Name: to.Ptr("<name>"),
							Definition: &armstorage.BlobInventoryPolicyDefinition{
								Format: to.Ptr(armstorage.FormatCSV),
								Filters: &armstorage.BlobInventoryPolicyFilter{
									BlobTypes: []*string{
										to.Ptr("blockBlob"),
										to.Ptr("appendBlob"),
										to.Ptr("pageBlob")},
									ExcludePrefix: []*string{
										to.Ptr("excludeprefix1"),
										to.Ptr("excludeprefix2")},
									IncludeBlobVersions: to.Ptr(true),
									IncludeSnapshots:    to.Ptr(true),
									PrefixMatch: []*string{
										to.Ptr("inventoryprefix1"),
										to.Ptr("inventoryprefix2")},
								},
								ObjectType: to.Ptr(armstorage.ObjectTypeBlob),
								Schedule:   to.Ptr(armstorage.ScheduleDaily),
								SchemaFields: []*string{
									to.Ptr("Name"),
									to.Ptr("Creation-Time"),
									to.Ptr("Last-Modified"),
									to.Ptr("Content-Length"),
									to.Ptr("Content-MD5"),
									to.Ptr("BlobType"),
									to.Ptr("AccessTier"),
									to.Ptr("AccessTierChangeTime"),
									to.Ptr("Snapshot"),
									to.Ptr("VersionId"),
									to.Ptr("IsCurrentVersion"),
									to.Ptr("Metadata")},
							},
							Destination: to.Ptr("<destination>"),
							Enabled:     to.Ptr(true),
						},
						{
							Name: to.Ptr("<name>"),
							Definition: &armstorage.BlobInventoryPolicyDefinition{
								Format:     to.Ptr(armstorage.FormatParquet),
								ObjectType: to.Ptr(armstorage.ObjectTypeContainer),
								Schedule:   to.Ptr(armstorage.ScheduleWeekly),
								SchemaFields: []*string{
									to.Ptr("Name"),
									to.Ptr("Last-Modified"),
									to.Ptr("Metadata"),
									to.Ptr("LeaseStatus"),
									to.Ptr("LeaseState"),
									to.Ptr("LeaseDuration"),
									to.Ptr("PublicAccess"),
									to.Ptr("HasImmutabilityPolicy"),
									to.Ptr("HasLegalHold")},
							},
							Destination: to.Ptr("<destination>"),
							Enabled:     to.Ptr(true),
						}},
				},
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
