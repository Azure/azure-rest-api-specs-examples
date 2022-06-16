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
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobInventoryPoliciesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"res7687",
		"sto9699",
		armstorage.BlobInventoryPolicyNameDefault,
		armstorage.BlobInventoryPolicy{
			Properties: &armstorage.BlobInventoryPolicyProperties{
				Policy: &armstorage.BlobInventoryPolicySchema{
					Type:    to.Ptr(armstorage.InventoryRuleTypeInventory),
					Enabled: to.Ptr(true),
					Rules: []*armstorage.BlobInventoryPolicyRule{
						{
							Name: to.Ptr("inventoryPolicyRule1"),
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
							Destination: to.Ptr("container1"),
							Enabled:     to.Ptr(true),
						},
						{
							Name: to.Ptr("inventoryPolicyRule2"),
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
							Destination: to.Ptr("container2"),
							Enabled:     to.Ptr(true),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
