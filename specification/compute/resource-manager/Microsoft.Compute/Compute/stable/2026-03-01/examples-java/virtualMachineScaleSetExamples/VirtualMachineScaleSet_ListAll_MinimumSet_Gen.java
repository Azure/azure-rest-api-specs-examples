
/**
 * Samples for VirtualMachineScaleSets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ListAll_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ListAll_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetListAllMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().list(com.azure.core.util.Context.NONE);
    }
}
