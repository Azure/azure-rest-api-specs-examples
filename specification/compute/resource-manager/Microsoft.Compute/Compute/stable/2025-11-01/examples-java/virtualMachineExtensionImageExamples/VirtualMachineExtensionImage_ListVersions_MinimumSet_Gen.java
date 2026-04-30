
/**
 * Samples for VirtualMachineExtensionImages ListVersions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_ListVersions_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImage_ListVersions_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineExtensionImageListVersionsMinimumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineExtensionImages().listVersionsWithResponse("aaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaa", null, null, null, com.azure.core.util.Context.NONE);
    }
}
