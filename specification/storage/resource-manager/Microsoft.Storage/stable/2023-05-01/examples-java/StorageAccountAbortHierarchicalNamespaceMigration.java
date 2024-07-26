
/**
 * Samples for StorageAccounts AbortHierarchicalNamespaceMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * StorageAccountAbortHierarchicalNamespaceMigration.json
     */
    /**
     * Sample code: StorageAccountAbortHierarchicalNamespaceMigration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        storageAccountAbortHierarchicalNamespaceMigration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts()
            .abortHierarchicalNamespaceMigration("res4228", "sto2434", com.azure.core.util.Context.NONE);
    }
}
