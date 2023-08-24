package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountSetBlobInventoryPolicyIncludeDeleteAndNewSchemaForHnsAccount.json
func ExampleBlobInventoryPoliciesClient_CreateOrUpdate_storageAccountSetBlobInventoryPolicyIncludeDeleteAndNewSchemaForHnsAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlobInventoryPoliciesClient().CreateOrUpdate(ctx, "res7687", "sto9699", armstorage.BlobInventoryPolicyNameDefault, armstorage.BlobInventoryPolicy{
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
								IncludeDeleted:      to.Ptr(true),
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
								to.Ptr("ContentType"),
								to.Ptr("ContentEncoding"),
								to.Ptr("ContentLanguage"),
								to.Ptr("ContentCRC64"),
								to.Ptr("CacheControl"),
								to.Ptr("Metadata"),
								to.Ptr("DeletionId"),
								to.Ptr("Deleted"),
								to.Ptr("DeletedTime"),
								to.Ptr("RemainingRetentionDays")},
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
								to.Ptr("HasLegalHold"),
								to.Ptr("Etag"),
								to.Ptr("DefaultEncryptionScope"),
								to.Ptr("DenyEncryptionScopeOverride"),
								to.Ptr("ImmutableStorageWithVersioningEnabled"),
								to.Ptr("Deleted"),
								to.Ptr("Version"),
								to.Ptr("DeletedTime"),
								to.Ptr("RemainingRetentionDays")},
						},
						Destination: to.Ptr("container2"),
						Enabled:     to.Ptr(true),
					}},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BlobInventoryPolicy = armstorage.BlobInventoryPolicy{
	// 	Name: to.Ptr("DefaultInventoryPolicy"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/inventoryPolicies"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7687/providers/Microsoft.Storage/storageAccounts/sto9699/inventoryPolicies/default"),
	// 	Properties: &armstorage.BlobInventoryPolicyProperties{
	// 		Policy: &armstorage.BlobInventoryPolicySchema{
	// 			Type: to.Ptr(armstorage.InventoryRuleTypeInventory),
	// 			Enabled: to.Ptr(true),
	// 			Rules: []*armstorage.BlobInventoryPolicyRule{
	// 				{
	// 					Name: to.Ptr("inventoryPolicyRule1"),
	// 					Definition: &armstorage.BlobInventoryPolicyDefinition{
	// 						Format: to.Ptr(armstorage.FormatCSV),
	// 						Filters: &armstorage.BlobInventoryPolicyFilter{
	// 							BlobTypes: []*string{
	// 								to.Ptr("blockBlob"),
	// 								to.Ptr("appendBlob"),
	// 								to.Ptr("pageBlob")},
	// 								ExcludePrefix: []*string{
	// 									to.Ptr("excludeprefix1"),
	// 									to.Ptr("excludeprefix2")},
	// 									IncludeBlobVersions: to.Ptr(true),
	// 									IncludeDeleted: to.Ptr(true),
	// 									IncludeSnapshots: to.Ptr(true),
	// 									PrefixMatch: []*string{
	// 										to.Ptr("inventoryprefix1"),
	// 										to.Ptr("inventoryprefix2")},
	// 									},
	// 									ObjectType: to.Ptr(armstorage.ObjectTypeBlob),
	// 									Schedule: to.Ptr(armstorage.ScheduleDaily),
	// 									SchemaFields: []*string{
	// 										to.Ptr("Name"),
	// 										to.Ptr("Creation-Time"),
	// 										to.Ptr("Last-Modified"),
	// 										to.Ptr("Content-Length"),
	// 										to.Ptr("Content-MD5"),
	// 										to.Ptr("BlobType"),
	// 										to.Ptr("AccessTier"),
	// 										to.Ptr("AccessTierChangeTime"),
	// 										to.Ptr("Snapshot"),
	// 										to.Ptr("VersionId"),
	// 										to.Ptr("IsCurrentVersion"),
	// 										to.Ptr("ContentType"),
	// 										to.Ptr("ContentEncoding"),
	// 										to.Ptr("ContentLanguage"),
	// 										to.Ptr("ContentCRC64"),
	// 										to.Ptr("CacheControl"),
	// 										to.Ptr("Metadata"),
	// 										to.Ptr("DeletionId"),
	// 										to.Ptr("Deleted"),
	// 										to.Ptr("DeletedTime"),
	// 										to.Ptr("RemainingRetentionDays")},
	// 									},
	// 									Destination: to.Ptr("container1"),
	// 									Enabled: to.Ptr(true),
	// 								},
	// 								{
	// 									Name: to.Ptr("inventoryPolicyRule2"),
	// 									Definition: &armstorage.BlobInventoryPolicyDefinition{
	// 										Format: to.Ptr(armstorage.FormatParquet),
	// 										ObjectType: to.Ptr(armstorage.ObjectTypeContainer),
	// 										Schedule: to.Ptr(armstorage.ScheduleWeekly),
	// 										SchemaFields: []*string{
	// 											to.Ptr("Name"),
	// 											to.Ptr("Last-Modified"),
	// 											to.Ptr("Metadata"),
	// 											to.Ptr("LeaseStatus"),
	// 											to.Ptr("LeaseState"),
	// 											to.Ptr("LeaseDuration"),
	// 											to.Ptr("PublicAccess"),
	// 											to.Ptr("HasImmutabilityPolicy"),
	// 											to.Ptr("HasLegalHold"),
	// 											to.Ptr("Etag"),
	// 											to.Ptr("DefaultEncryptionScope"),
	// 											to.Ptr("DenyEncryptionScopeOverride"),
	// 											to.Ptr("ImmutableStorageWithVersioningEnabled"),
	// 											to.Ptr("Deleted"),
	// 											to.Ptr("Version"),
	// 											to.Ptr("DeletedTime"),
	// 											to.Ptr("RemainingRetentionDays")},
	// 										},
	// 										Destination: to.Ptr("container2"),
	// 										Enabled: to.Ptr(true),
	// 								}},
	// 							},
	// 						},
	// 					}
}
