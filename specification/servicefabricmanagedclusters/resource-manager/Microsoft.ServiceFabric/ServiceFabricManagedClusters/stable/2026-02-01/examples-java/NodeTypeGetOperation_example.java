
/**
 * Samples for NodeTypes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/NodeTypeGetOperation_example.json
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
