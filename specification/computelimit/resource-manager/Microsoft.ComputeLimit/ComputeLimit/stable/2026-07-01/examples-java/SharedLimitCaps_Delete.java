
/**
 * Samples for SharedLimitCaps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SharedLimitCaps_Delete.json
     */
    /**
     * Sample code: Delete the shared limit cap for a VM family.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void
        deleteTheSharedLimitCapForAVMFamily(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.sharedLimitCaps().deleteByResourceGroupWithResponse("eastus", "StandardDSv3Family",
            com.azure.core.util.Context.NONE);
    }
}
