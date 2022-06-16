import com.azure.core.util.Context;

/** Samples for StorageAccounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountDelete.json
     */
    /**
     * Sample code: StorageAccountDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .deleteWithResponse("res4228", "sto2434", Context.NONE);
    }
}
