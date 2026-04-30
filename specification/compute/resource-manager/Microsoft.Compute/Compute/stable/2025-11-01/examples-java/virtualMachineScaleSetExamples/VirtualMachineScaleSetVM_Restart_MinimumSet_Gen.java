
/**
 * Samples for VirtualMachineScaleSetVMs Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Restart_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Restart_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMRestartMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().restart("rgcompute", "aaaaaaaaaaaa", "aaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
