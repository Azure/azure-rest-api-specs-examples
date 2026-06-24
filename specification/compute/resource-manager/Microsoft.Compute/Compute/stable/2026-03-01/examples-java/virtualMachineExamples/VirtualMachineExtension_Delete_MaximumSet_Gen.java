
/**
 * Samples for VirtualMachineExtensions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachineExtension_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtension_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineExtensionDeleteMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineExtensions().delete("rgcompute", "aaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
