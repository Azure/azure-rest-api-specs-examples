
/**
 * Samples for SharedLimits Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SharedLimits_Delete.json
     */
    /**
     * Sample code: Delete a shared limit.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void deleteASharedLimit(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.sharedLimits().deleteByResourceGroupWithResponse("eastus", "StandardDSv3Family",
            com.azure.core.util.Context.NONE);
    }
}
