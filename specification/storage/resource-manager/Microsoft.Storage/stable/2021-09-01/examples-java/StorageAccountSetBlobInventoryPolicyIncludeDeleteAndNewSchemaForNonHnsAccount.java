import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.BlobInventoryPolicyInner;
import com.azure.resourcemanager.storage.models.BlobInventoryPolicyDefinition;
import com.azure.resourcemanager.storage.models.BlobInventoryPolicyFilter;
import com.azure.resourcemanager.storage.models.BlobInventoryPolicyName;
import com.azure.resourcemanager.storage.models.BlobInventoryPolicyRule;
import com.azure.resourcemanager.storage.models.BlobInventoryPolicySchema;
import com.azure.resourcemanager.storage.models.Format;
import com.azure.resourcemanager.storage.models.InventoryRuleType;
import com.azure.resourcemanager.storage.models.ObjectType;
import com.azure.resourcemanager.storage.models.Schedule;
import java.util.Arrays;

/** Samples for BlobInventoryPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountSetBlobInventoryPolicyIncludeDeleteAndNewSchemaForNonHnsAccount.json
     */
    /**
     * Sample code: StorageAccountSetBlobInventoryPolicyIncludeDeleteAndNewSchemaForNonHnsAccount.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountSetBlobInventoryPolicyIncludeDeleteAndNewSchemaForNonHnsAccount(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobInventoryPolicies()
            .createOrUpdateWithResponse(
                "res7687",
                "sto9699",
                BlobInventoryPolicyName.DEFAULT,
                new BlobInventoryPolicyInner()
                    .withPolicy(
                        new BlobInventoryPolicySchema()
                            .withEnabled(true)
                            .withType(InventoryRuleType.INVENTORY)
                            .withRules(
                                Arrays
                                    .asList(
                                        new BlobInventoryPolicyRule()
                                            .withEnabled(true)
                                            .withName("inventoryPolicyRule1")
                                            .withDestination("container1")
                                            .withDefinition(
                                                new BlobInventoryPolicyDefinition()
                                                    .withFilters(
                                                        new BlobInventoryPolicyFilter()
                                                            .withPrefixMatch(
                                                                Arrays.asList("inventoryprefix1", "inventoryprefix2"))
                                                            .withExcludePrefix(
                                                                Arrays.asList("excludeprefix1", "excludeprefix2"))
                                                            .withBlobTypes(
                                                                Arrays.asList("blockBlob", "appendBlob", "pageBlob"))
                                                            .withIncludeBlobVersions(true)
                                                            .withIncludeSnapshots(true)
                                                            .withIncludeDeleted(true))
                                                    .withFormat(Format.CSV)
                                                    .withSchedule(Schedule.DAILY)
                                                    .withObjectType(ObjectType.BLOB)
                                                    .withSchemaFields(
                                                        Arrays
                                                            .asList(
                                                                "Name",
                                                                "Creation-Time",
                                                                "Last-Modified",
                                                                "Content-Length",
                                                                "Content-MD5",
                                                                "BlobType",
                                                                "AccessTier",
                                                                "AccessTierChangeTime",
                                                                "Snapshot",
                                                                "VersionId",
                                                                "IsCurrentVersion",
                                                                "Tags",
                                                                "ContentType",
                                                                "ContentEncoding",
                                                                "ContentLanguage",
                                                                "ContentCRC64",
                                                                "CacheControl",
                                                                "Metadata",
                                                                "Deleted",
                                                                "RemainingRetentionDays"))),
                                        new BlobInventoryPolicyRule()
                                            .withEnabled(true)
                                            .withName("inventoryPolicyRule2")
                                            .withDestination("container2")
                                            .withDefinition(
                                                new BlobInventoryPolicyDefinition()
                                                    .withFormat(Format.PARQUET)
                                                    .withSchedule(Schedule.WEEKLY)
                                                    .withObjectType(ObjectType.CONTAINER)
                                                    .withSchemaFields(
                                                        Arrays
                                                            .asList(
                                                                "Name",
                                                                "Last-Modified",
                                                                "Metadata",
                                                                "LeaseStatus",
                                                                "LeaseState",
                                                                "LeaseDuration",
                                                                "PublicAccess",
                                                                "HasImmutabilityPolicy",
                                                                "HasLegalHold",
                                                                "Etag",
                                                                "DefaultEncryptionScope",
                                                                "DenyEncryptionScopeOverride",
                                                                "ImmutableStorageWithVersioningEnabled",
                                                                "Deleted",
                                                                "Version",
                                                                "DeletedTime",
                                                                "RemainingRetentionDays")))))),
                Context.NONE);
    }
}
