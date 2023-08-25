import com.azure.resourcemanager.storage.models.FailoverType;

/** Samples for StorageAccounts Failover. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountFailoverPlanned.json
     */
    /**
     * Sample code: StorageAccountFailoverPlanned.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountFailoverPlanned(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .failover("res4228", "sto2434", FailoverType.PLANNED, com.azure.core.util.Context.NONE);
    }
}
