
/**
 * Samples for NodePools ListBySupercomputer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/NodePools_ListBySupercomputer_MaximumSet_Gen.json
     */
    /**
     * Sample code: NodePools_ListBySupercomputer_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        nodePoolsListBySupercomputerMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.nodePools().listBySupercomputer("rgdiscovery", "a4d55e3b47501e6fe1", com.azure.core.util.Context.NONE);
    }
}
