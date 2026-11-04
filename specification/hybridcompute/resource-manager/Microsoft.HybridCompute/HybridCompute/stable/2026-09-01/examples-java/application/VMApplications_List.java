
/**
 * Samples for VmApplications List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/application/VMApplications_List.json
     */
    /**
     * Sample code: List VM applications.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void listVMApplications(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.vmApplications().list("myResourceGroup", "myMachine", 100, null, com.azure.core.util.Context.NONE);
    }
}
