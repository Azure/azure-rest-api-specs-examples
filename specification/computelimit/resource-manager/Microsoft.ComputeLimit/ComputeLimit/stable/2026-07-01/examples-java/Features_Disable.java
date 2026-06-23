
/**
 * Samples for Features Disable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Features_Disable.json
     */
    /**
     * Sample code: Disable feature.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void disableFeature(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.features().disable("eastus", "VmCategoryQuota", com.azure.core.util.Context.NONE);
    }
}
