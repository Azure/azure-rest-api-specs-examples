
import com.azure.resourcemanager.storage.models.MigrationName;

/**
 * Samples for StorageAccounts GetCustomerInitiatedMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountGetMigrationInProgress.json
     */
    /**
     * Sample code: StorageAccountGetMigrationInProgress.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountGetMigrationInProgress(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().getCustomerInitiatedMigrationWithResponse("resource-group-name",
            "accountname", MigrationName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
