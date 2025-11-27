
/**
 * Samples for NodeTypes ListByManagedClusters.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/NodeTypeListOperation_example.json
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
