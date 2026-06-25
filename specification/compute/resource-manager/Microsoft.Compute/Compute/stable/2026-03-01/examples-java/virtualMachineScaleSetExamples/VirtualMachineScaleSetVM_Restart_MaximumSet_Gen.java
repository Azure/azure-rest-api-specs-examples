
/**
 * Samples for VirtualMachineScaleSetVMs Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Restart_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Restart_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMRestartMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().restart("rgcompute", "aa", "aaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
