
/**
 * Samples for VmApplications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/application/VMApplications_Delete.json
     */
    /**
     * Sample code: Delete a VM application.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void deleteAVMApplication(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.vmApplications().delete("myResourceGroup", "myMachine", "myApplication",
            com.azure.core.util.Context.NONE);
    }
}
