
/**
 * Samples for VirtualMachineScaleSetVMs PowerOff.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_PowerOff_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_PowerOff_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMPowerOffMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().powerOff("rgcompute", "aaaaaa", "aaaaaaaaa", true,
            com.azure.core.util.Context.NONE);
    }
}
