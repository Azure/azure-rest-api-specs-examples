
/**
 * Samples for Projects Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Projects_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_Delete_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void projectsDeleteMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.projects().delete("rgdiscovery", "97520ee2d6a76d232e", "e9f886be682c26b909",
            com.azure.core.util.Context.NONE);
    }
}
