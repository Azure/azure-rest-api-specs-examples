
/**
 * Samples for NetworkProfile Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-16-preview/NetworkProfile_Get.json
     */
    /**
     * Sample code: GET Network Profile.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void gETNetworkProfile(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.networkProfiles().getWithResponse("myResourceGroup", "myMachine", com.azure.core.util.Context.NONE);
    }
}
