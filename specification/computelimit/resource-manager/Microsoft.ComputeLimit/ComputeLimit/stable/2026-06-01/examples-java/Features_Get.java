
/**
 * Samples for Features Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Features_Get.json
     */
    /**
     * Sample code: Get feature.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void getFeature(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.features().getWithResponse("eastus", "VmCategoryQuota", com.azure.core.util.Context.NONE);
    }
}
