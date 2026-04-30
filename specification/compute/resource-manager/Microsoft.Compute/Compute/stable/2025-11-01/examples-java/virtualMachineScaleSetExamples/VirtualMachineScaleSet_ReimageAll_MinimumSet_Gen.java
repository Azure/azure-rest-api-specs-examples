
/**
 * Samples for VirtualMachineScaleSets ReimageAll.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ReimageAll_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ReimageAll_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetReimageAllMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().reimageAll("rgcompute", "aaaaaaaaaaaaaaaaaaaaaa", null,
            com.azure.core.util.Context.NONE);
    }
}
