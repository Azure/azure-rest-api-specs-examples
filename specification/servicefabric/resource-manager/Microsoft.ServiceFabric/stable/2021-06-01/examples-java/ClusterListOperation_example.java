/** Samples for Clusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterListOperation_example.json
     */
    /**
     * Sample code: List clusters.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void listClusters(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.clusters().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
