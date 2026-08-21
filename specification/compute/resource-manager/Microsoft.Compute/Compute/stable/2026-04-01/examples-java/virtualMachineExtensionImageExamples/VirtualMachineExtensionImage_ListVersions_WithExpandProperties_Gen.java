
import com.azure.resourcemanager.compute.models.ListVersionsExpandOptions;

/**
 * Samples for VirtualMachineExtensionImages ListVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/virtualMachineExtensionImageExamples/
     * VirtualMachineExtensionImage_ListVersions_WithExpandProperties_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImage_ListVersions_WithExpandProperties_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineExtensionImageListVersionsWithExpandPropertiesGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineExtensionImages().listVersionsWithResponse(
            "aaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaa", 22,
            "a", ListVersionsExpandOptions.PROPERTIES, com.azure.core.util.Context.NONE);
    }
}
