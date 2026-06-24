
/**
 * Samples for VirtualMachineScaleSetVMs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMListMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().list("rgcompute", "aaaaaaaaaaaaaaaaaaaaaa",
            "aaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
