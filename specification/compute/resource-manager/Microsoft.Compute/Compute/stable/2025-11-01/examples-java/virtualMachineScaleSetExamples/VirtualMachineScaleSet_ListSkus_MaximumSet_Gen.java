
/**
 * Samples for VirtualMachineScaleSets ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ListSkus_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ListSkus_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetListSkusMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().listSkus("rgcompute", "aaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
