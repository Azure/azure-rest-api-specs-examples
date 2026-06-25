
/**
 * Samples for VirtualMachineScaleSets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetListMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().listByResourceGroup("rgcompute",
            com.azure.core.util.Context.NONE);
    }
}
