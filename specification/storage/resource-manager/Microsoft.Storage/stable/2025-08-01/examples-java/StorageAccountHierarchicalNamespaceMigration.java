
/**
 * Samples for StorageAccounts HierarchicalNamespaceMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountHierarchicalNamespaceMigration.json
     */
    /**
     * Sample code: StorageAccountHierarchicalNamespaceMigration.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountHierarchicalNamespaceMigration(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().hierarchicalNamespaceMigration("res4228", "sto2434",
            "HnsOnValidationRequest", com.azure.core.util.Context.NONE);
    }
}
