
import com.azure.resourcemanager.storage.models.MigrationName;

/** Samples for StorageAccounts GetCustomerInitiatedMigration. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/
     * StorageAccountGetMigrationInProgress.json
     */
    /**
     * Sample code: StorageAccountGetMigrationInProgress.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountGetMigrationInProgress(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts()
            .getCustomerInitiatedMigrationWithResponse("resource-group-name", "accountname", MigrationName.DEFAULT,
                com.azure.core.util.Context.NONE);
    }
}
