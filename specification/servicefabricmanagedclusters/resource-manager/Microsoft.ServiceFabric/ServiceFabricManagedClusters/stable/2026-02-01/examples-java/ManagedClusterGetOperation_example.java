
/**
 * Samples for ManagedClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedClusterGetOperation_example.json
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
