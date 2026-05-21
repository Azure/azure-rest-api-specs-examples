
import com.azure.resourcemanager.storage.fluent.models.StorageAccountMigrationInner;
import com.azure.resourcemanager.storage.models.SkuName;

/**
 * Samples for StorageAccounts CustomerInitiatedMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountPostMigration.json
     */
    /**
     * Sample code: StorageAccountPostMigration.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountPostMigration(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().customerInitiatedMigration("resource-group-name", "accountname",
            new StorageAccountMigrationInner().withTargetSkuName(SkuName.STANDARD_ZRS),
            com.azure.core.util.Context.NONE);
    }
}
