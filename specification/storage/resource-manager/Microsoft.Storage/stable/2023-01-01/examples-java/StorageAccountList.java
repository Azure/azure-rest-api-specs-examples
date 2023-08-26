/** Samples for StorageAccounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountList.json
     */
    /**
     * Sample code: StorageAccountList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts().list(com.azure.core.util.Context.NONE);
    }
}
