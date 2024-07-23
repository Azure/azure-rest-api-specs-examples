
/**
 * Samples for Machines AssessPatches.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/machine/
     * Machine_AssessPatches.json
     */
    /**
     * Sample code: Assess patch state of a machine.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        assessPatchStateOfAMachine(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machines().assessPatches("myResourceGroupName", "myMachineName", com.azure.core.util.Context.NONE);
    }
}
