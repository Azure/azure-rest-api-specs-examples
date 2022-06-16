import com.azure.core.util.Context;

/** Samples for StorageAccounts Failover. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/StorageAccountFailover.json
     */
    /**
     * Sample code: StorageAccountFailover.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountFailover(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .failover("res4228", "sto2434", Context.NONE);
    }
}
