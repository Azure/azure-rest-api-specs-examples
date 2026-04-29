
/**
 * Samples for VirtualMachineScaleSetVMs Deallocate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Deallocate_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Deallocate_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMDeallocateMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().deallocate("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
