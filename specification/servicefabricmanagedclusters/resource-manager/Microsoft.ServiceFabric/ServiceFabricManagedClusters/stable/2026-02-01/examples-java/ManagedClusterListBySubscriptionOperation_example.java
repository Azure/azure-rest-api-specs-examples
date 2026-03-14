
/**
 * Samples for ManagedClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedClusterListBySubscriptionOperation_example.json
     */
    /**
     * Sample code: List managed clusters.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void listManagedClusters(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedClusters().list(com.azure.core.util.Context.NONE);
    }
}
