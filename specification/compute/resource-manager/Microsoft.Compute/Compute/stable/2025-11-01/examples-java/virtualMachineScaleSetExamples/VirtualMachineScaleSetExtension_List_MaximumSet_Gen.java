
/**
 * Samples for VirtualMachineScaleSetExtensions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetExtension_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetExtensionListMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetExtensions().list("rgcompute", "aaaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
