import com.azure.resourcemanager.storage.fluent.models.StorageAccountMigrationInner;
import com.azure.resourcemanager.storage.models.SkuName;

/** Samples for StorageAccounts CustomerInitiatedMigration. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountPostMigration.json
     */
    /**
     * Sample code: StorageAccountPostMigration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountPostMigration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .customerInitiatedMigration(
                "resource-group-name",
                "accountname",
                new StorageAccountMigrationInner().withTargetSkuName(SkuName.STANDARD_ZRS),
                com.azure.core.util.Context.NONE);
    }
}
