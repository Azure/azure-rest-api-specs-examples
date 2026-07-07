
/**
 * Samples for StorageAccounts AbortHierarchicalNamespaceMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountAbortHierarchicalNamespaceMigration.json
     */
    /**
     * Sample code: StorageAccountAbortHierarchicalNamespaceMigration.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountAbortHierarchicalNamespaceMigration(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().abortHierarchicalNamespaceMigration("res4228", "sto2434",
            com.azure.core.util.Context.NONE);
    }
}
