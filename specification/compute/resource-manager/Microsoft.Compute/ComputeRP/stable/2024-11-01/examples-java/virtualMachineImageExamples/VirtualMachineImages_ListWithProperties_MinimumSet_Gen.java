
import com.azure.resourcemanager.compute.models.Expand;

/**
 * Samples for VirtualMachineImages ListWithProperties.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineImageExamples/VirtualMachineImages_ListWithProperties_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImages_ListWithProperties_MinimumSet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineImagesListWithPropertiesMinimumSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineImages().listWithProperties("eastus",
            "MicrosoftWindowsServer", "WindowsServer", "2022-datacenter-azure-edition", Expand.PROPERTIES, null, null,
            com.azure.core.util.Context.NONE);
    }
}
