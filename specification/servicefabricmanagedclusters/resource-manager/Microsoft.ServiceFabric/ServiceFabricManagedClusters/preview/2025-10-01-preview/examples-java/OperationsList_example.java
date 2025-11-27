
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/OperationsList_example.json
     */
    /**
     * Sample code: List the operations for the provider.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void listTheOperationsForTheProvider(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
