
/**
 * Samples for VirtualMachineScaleSetVMs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMListMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().list("rgcompute", "aaaaaaaaaaaaaa", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
