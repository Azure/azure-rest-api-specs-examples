
/**
 * Samples for NodeTypes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/
     * examples/NodeTypeDeleteOperation_example.json
     */
    /**
     * Sample code: Delete a node type.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void deleteANodeType(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().delete("resRg", "myCluster", "BE", com.azure.core.util.Context.NONE);
    }
}
