
/**
 * Samples for VmFamilies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/VmFamilies_Get.json
     */
    /**
     * Sample code: Get a VM family.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void getAVMFamily(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.vmFamilies().getWithResponse("eastus", "standardDSv2Family", com.azure.core.util.Context.NONE);
    }
}
