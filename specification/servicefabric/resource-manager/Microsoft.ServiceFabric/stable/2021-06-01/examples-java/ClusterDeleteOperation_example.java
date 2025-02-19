
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/
     * ClusterDeleteOperation_example.json
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
