
/**
 * Samples for NodePools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/NodePools_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NodePools_Get_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void nodePoolsGetMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.nodePools().getWithResponse("rgdiscovery", "68ccaea8f927d3c9d7", "f86825f20c4fb1d2fc",
            com.azure.core.util.Context.NONE);
    }
}
