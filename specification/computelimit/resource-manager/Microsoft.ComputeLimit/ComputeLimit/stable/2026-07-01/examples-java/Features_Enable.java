
import com.azure.resourcemanager.computelimit.models.FeatureEnableRequest;

/**
 * Samples for Features Enable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Features_Enable.json
     */
    /**
     * Sample code: Enable feature.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void enableFeature(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.features().enable("eastus", "VmCategoryQuota",
            new FeatureEnableRequest().withServiceTreeId("a1b2c3d4-5678-90ab-cdef-1234567890ab"),
            com.azure.core.util.Context.NONE);
    }
}
