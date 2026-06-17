
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ClusterDeleteOperation_example.json
     */
    /**
     * Sample code: Delete a cluster.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void deleteACluster(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.clusters().deleteByResourceGroupWithResponse("resRg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
