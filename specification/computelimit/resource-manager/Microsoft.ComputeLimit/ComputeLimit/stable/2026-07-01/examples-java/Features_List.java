
/**
 * Samples for Features ListBySubscriptionLocationResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Features_List.json
     */
    /**
     * Sample code: List features.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void listFeatures(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.features().listBySubscriptionLocationResource("eastus", com.azure.core.util.Context.NONE);
    }
}
