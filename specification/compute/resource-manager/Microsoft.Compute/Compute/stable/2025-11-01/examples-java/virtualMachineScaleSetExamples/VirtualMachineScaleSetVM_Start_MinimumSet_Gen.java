
/**
 * Samples for VirtualMachineScaleSetVMs Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Start_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Start_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMStartMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().start("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
