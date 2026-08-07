
/**
 * Samples for Tools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Tools_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Tools_Delete_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void toolsDeleteMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.tools().delete("rgdiscovery", "f127c90bef940264e3", com.azure.core.util.Context.NONE);
    }
}
