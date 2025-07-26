
/**
 * Samples for VirtualMachineImages ListWithProperties.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineImageExamples/VirtualMachineImages_ListWithProperties_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImages_ListWithProperties_MaximumSet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineImagesListWithPropertiesMaximumSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineImages().listWithPropertiesWithResponse(
            "eastus", "MicrosoftWindowsServer", "WindowsServer", "2022-datacenter-azure-edition", "Properties", 4, "aa",
            com.azure.core.util.Context.NONE);
    }
}
