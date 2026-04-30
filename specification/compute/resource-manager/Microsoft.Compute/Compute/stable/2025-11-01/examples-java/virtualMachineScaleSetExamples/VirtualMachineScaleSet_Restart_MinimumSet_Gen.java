
/**
 * Samples for VirtualMachineScaleSets Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Restart_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Restart_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetRestartMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().restart("rgcompute", "aaaa", null,
            com.azure.core.util.Context.NONE);
    }
}
