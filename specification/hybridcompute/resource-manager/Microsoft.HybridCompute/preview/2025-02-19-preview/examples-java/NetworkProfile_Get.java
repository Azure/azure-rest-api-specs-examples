
/**
 * Samples for NetworkProfile Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/
     * NetworkProfile_Get.json
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
