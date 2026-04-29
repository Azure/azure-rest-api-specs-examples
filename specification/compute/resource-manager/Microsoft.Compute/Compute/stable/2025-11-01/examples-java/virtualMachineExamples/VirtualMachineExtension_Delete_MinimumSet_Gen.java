
/**
 * Samples for VirtualMachineExtensions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachineExtension_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtension_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineExtensionDeleteMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineExtensions().delete("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaa", "aa",
            com.azure.core.util.Context.NONE);
    }
}
