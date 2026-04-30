
/**
 * Samples for VirtualMachineScaleSets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetListMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().listByResourceGroup("rgcompute",
            com.azure.core.util.Context.NONE);
    }
}
