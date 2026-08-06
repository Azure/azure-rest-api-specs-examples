
/**
 * Samples for Tools GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Tools_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Tools_Get_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void toolsGetMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.tools().getByResourceGroupWithResponse("rgdiscovery", "1ba7068ab4d3671156",
            com.azure.core.util.Context.NONE);
    }
}
