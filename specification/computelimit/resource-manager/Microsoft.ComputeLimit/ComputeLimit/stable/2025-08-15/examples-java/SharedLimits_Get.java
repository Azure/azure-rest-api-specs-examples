
/**
 * Samples for SharedLimits Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-15/SharedLimits_Get.json
     */
    /**
     * Sample code: Get a shared limit.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void getASharedLimit(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.sharedLimits().getWithResponse("eastus", "StandardDSv3Family", com.azure.core.util.Context.NONE);
    }
}
