
/**
 * Samples for VirtualMachineScaleSetVMs Deallocate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Deallocate_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Deallocate_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMDeallocateMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().deallocate("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaa",
            "aaaaaa", com.azure.core.util.Context.NONE);
    }
}
