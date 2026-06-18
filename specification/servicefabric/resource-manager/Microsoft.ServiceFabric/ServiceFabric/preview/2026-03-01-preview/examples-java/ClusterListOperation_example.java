
/**
 * Samples for Clusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ClusterListOperation_example.json
     */
    /**
     * Sample code: List clusters.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void listClusters(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.clusters().list(com.azure.core.util.Context.NONE);
    }
}
