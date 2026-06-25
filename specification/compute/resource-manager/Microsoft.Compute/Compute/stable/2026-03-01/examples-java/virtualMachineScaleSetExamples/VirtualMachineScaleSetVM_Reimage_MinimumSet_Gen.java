
/**
 * Samples for VirtualMachineScaleSetVMs Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Reimage_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Reimage_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMReimageMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().reimage("rgcompute", "aaaaaaa", "aaaaaaaaaaaaa", null,
            com.azure.core.util.Context.NONE);
    }
}
