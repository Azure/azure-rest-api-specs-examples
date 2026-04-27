
/**
 * Samples for Features Enable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/Features_Enable.json
     */
    /**
     * Sample code: Enable feature.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void enableFeature(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.features().enable("eastus", "VmCategoryQuota", com.azure.core.util.Context.NONE);
    }
}
