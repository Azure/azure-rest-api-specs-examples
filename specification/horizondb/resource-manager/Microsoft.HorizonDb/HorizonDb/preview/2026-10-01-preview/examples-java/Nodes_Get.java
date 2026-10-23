
/**
 * Samples for Nodes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Nodes_Get.json
     */
    /**
     * Sample code: Get a HorizonDB node.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAHorizonDBNode(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.nodes().getWithResponse("example-resource-group", "example-cluster", "example-pool-01", "node-01",
            com.azure.core.util.Context.NONE);
    }
}
