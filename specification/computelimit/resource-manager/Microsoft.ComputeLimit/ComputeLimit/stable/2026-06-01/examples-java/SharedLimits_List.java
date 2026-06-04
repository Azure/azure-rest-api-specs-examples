
/**
 * Samples for SharedLimits ListBySubscriptionLocationResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/SharedLimits_List.json
     */
    /**
     * Sample code: List all shared limits for a scope.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void
        listAllSharedLimitsForAScope(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.sharedLimits().listBySubscriptionLocationResource("eastus", com.azure.core.util.Context.NONE);
    }
}
