
/**
 * Samples for Projects Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Projects_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_Get_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void projectsGetMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.projects().getWithResponse("rgdiscovery", "80895d77522bf22889", "b8f0217d144f00d223",
            com.azure.core.util.Context.NONE);
    }
}
