
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/
     * Operations_example.json
     */
    /**
     * Sample code: List available operations.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void listAvailableOperations(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
