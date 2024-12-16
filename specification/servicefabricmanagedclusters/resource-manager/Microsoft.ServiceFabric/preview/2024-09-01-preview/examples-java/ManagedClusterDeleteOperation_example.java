
/**
 * Samples for ManagedClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/
     * examples/ManagedClusterDeleteOperation_example.json
     */
    /**
     * Sample code: Delete a cluster.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void deleteACluster(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedClusters().delete("resRg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
