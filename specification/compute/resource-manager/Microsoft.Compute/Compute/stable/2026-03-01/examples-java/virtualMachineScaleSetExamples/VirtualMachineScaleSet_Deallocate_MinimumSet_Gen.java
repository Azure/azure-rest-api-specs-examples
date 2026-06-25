
/**
 * Samples for VirtualMachineScaleSets Deallocate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Deallocate_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Deallocate_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetDeallocateMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().deallocate("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaa", null,
            null, com.azure.core.util.Context.NONE);
    }
}
