
/**
 * Samples for VmFamilies ListBySubscriptionLocationResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/VmFamilies_List.json
     */
    /**
     * Sample code: List VM families.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void listVMFamilies(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.vmFamilies().listBySubscriptionLocationResource("eastus", null, com.azure.core.util.Context.NONE);
    }
}
