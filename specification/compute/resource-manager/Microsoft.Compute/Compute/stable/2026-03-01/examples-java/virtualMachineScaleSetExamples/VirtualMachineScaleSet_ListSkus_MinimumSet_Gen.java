
/**
 * Samples for VirtualMachineScaleSets ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ListSkus_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ListSkus_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetListSkusMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().listSkus("rgcompute", "aaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
