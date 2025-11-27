
/**
 * Samples for ManagedClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ManagedClusterGetOperation_example.json
     */
    /**
     * Sample code: Get a cluster.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getACluster(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedClusters().getByResourceGroupWithResponse("resRg", "myCluster",
            com.azure.core.util.Context.NONE);
    }
}
