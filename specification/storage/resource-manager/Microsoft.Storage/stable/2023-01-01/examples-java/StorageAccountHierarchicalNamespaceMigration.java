
/** Samples for StorageAccounts HierarchicalNamespaceMigration. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/
     * StorageAccountHierarchicalNamespaceMigration.json
     */
    /**
     * Sample code: StorageAccountHierarchicalNamespaceMigration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        storageAccountHierarchicalNamespaceMigration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts().hierarchicalNamespaceMigration("res4228",
            "sto2434", "HnsOnValidationRequest", com.azure.core.util.Context.NONE);
    }
}
