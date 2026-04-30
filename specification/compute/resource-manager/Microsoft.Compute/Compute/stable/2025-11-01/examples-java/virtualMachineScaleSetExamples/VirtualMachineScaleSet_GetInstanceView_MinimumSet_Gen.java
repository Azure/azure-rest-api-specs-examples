
/**
 * Samples for VirtualMachineScaleSets GetInstanceView.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_GetInstanceView_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_GetInstanceView_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetGetInstanceViewMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().getInstanceViewWithResponse("rgcompute", "aaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
