
/**
 * Samples for Features Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Features_Get_SharedLimit.json
     */
    /**
     * Sample code: Get SharedLimit feature.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void getSharedLimitFeature(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.features().getWithResponse("eastus", "SharedLimit", com.azure.core.util.Context.NONE);
    }
}
