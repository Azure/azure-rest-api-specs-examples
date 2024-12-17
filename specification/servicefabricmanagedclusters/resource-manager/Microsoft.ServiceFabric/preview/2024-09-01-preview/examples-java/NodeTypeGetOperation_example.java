
/**
 * Samples for NodeTypes Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/
     * examples/NodeTypeGetOperation_example.json
     */
    /**
     * Sample code: Get a node type.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getANodeType(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().getWithResponse("resRg", "myCluster", "FE", com.azure.core.util.Context.NONE);
    }
}
