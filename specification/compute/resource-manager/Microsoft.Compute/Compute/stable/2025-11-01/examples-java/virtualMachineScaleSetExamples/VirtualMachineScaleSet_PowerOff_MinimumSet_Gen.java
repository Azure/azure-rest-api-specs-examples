
/**
 * Samples for VirtualMachineScaleSets PowerOff.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_PowerOff_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_PowerOff_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetPowerOffMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().powerOff("rgcompute", "a", null, null,
            com.azure.core.util.Context.NONE);
    }
}
