
/**
 * Samples for VirtualMachineScaleSets Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Reimage_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Reimage_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetReimageMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().reimage("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaa", null,
            com.azure.core.util.Context.NONE);
    }
}
