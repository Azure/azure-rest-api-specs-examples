
/**
 * Samples for Supercomputers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Supercomputers_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Supercomputers_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        supercomputersListByResourceGroupMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.supercomputers().listByResourceGroup("rgdiscovery", com.azure.core.util.Context.NONE);
    }
}
