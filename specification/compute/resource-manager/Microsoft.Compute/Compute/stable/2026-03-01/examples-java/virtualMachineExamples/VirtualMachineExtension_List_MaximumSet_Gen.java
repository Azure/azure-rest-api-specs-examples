
/**
 * Samples for VirtualMachineExtensions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachineExtension_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtension_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineExtensionListMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineExtensions().listWithResponse("rgcompute", "aaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
