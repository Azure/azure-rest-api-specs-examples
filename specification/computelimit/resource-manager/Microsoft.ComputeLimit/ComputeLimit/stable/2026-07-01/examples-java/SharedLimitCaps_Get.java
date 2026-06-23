
/**
 * Samples for SharedLimitCaps Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SharedLimitCaps_Get.json
     */
    /**
     * Sample code: Get a shared limit cap for a VM family.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void
        getASharedLimitCapForAVMFamily(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.sharedLimitCaps().getWithResponse("eastus", "StandardDSv3Family", com.azure.core.util.Context.NONE);
    }
}
