
/**
 * Samples for VmApplications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/application/VMApplications_Get.json
     */
    /**
     * Sample code: Get a VM application.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void getAVMApplication(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.vmApplications().getWithResponse("myResourceGroup", "myMachine", "myApplication",
            com.azure.core.util.Context.NONE);
    }
}
