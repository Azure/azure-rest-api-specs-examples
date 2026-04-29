
/**
 * Samples for VirtualMachineScaleSets GetInstanceView.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_GetInstanceView_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_GetInstanceView_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetGetInstanceViewMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().getInstanceViewWithResponse("rgcompute", "aaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
