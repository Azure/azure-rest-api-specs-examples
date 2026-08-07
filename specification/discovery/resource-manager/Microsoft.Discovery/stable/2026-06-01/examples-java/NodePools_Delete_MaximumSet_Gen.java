
/**
 * Samples for NodePools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/NodePools_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NodePools_Delete_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void nodePoolsDeleteMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.nodePools().delete("rgdiscovery", "5024d5e8fe2b743588", "d79a36a71cc10c19ce",
            com.azure.core.util.Context.NONE);
    }
}
