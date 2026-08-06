
/**
 * Samples for Tools ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Tools_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Tools_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        toolsListByResourceGroupMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.tools().listByResourceGroup("rgdiscovery", com.azure.core.util.Context.NONE);
    }
}
