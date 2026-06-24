
/**
 * Samples for VirtualMachineScaleSetVMs Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Start_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Start_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMStartMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().start("rgcompute", "aaaaaaaaaaaaaa", "aaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
