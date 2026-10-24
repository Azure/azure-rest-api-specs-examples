
/**
 * Samples for Nodes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Nodes_List.json
     */
    /**
     * Sample code: List the nodes allocated to a HorizonDB pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listTheNodesAllocatedToAHorizonDBPool(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.nodes().list("example-resource-group", "example-cluster", "example-pool-01",
            com.azure.core.util.Context.NONE);
    }
}
