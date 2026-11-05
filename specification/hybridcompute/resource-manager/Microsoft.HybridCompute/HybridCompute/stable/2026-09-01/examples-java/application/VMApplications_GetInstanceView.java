
/**
 * Samples for VmApplications GetInstanceView.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/application/VMApplications_GetInstanceView.json
     */
    /**
     * Sample code: Get the instance view of a VM application.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        getTheInstanceViewOfAVMApplication(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.vmApplications().getInstanceViewWithResponse("myResourceGroup", "myMachine", "myApplication",
            com.azure.core.util.Context.NONE);
    }
}
