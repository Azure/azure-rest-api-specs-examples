
/**
 * Samples for VirtualMachineExtensionImages ListVersions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_ListVersions_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImage_ListVersions_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineExtensionImageListVersionsMaximumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineExtensionImages().listVersionsWithResponse(
            "aaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaa", 22,
            "a", com.azure.core.util.Context.NONE);
    }
}
