
/**
 * Samples for NodeTypes ListByManagedClusters.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/
     * examples/NodeTypeListOperation_example.json
     */
    /**
     * Sample code: List node type of the specified managed cluster.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void listNodeTypeOfTheSpecifiedManagedCluster(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().listByManagedClusters("resRg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
