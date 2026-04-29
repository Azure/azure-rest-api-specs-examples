
/**
 * Samples for VirtualMachineImages ListWithProperties.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineImageExamples/VirtualMachineImages_ListWithProperties_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImages_ListWithProperties_MinimumSet.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineImagesListWithPropertiesMinimumSet(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().listWithPropertiesWithResponse("eastus",
            "MicrosoftWindowsServer", "WindowsServer", "2022-datacenter-azure-edition", "Properties", null, null,
            com.azure.core.util.Context.NONE);
    }
}
