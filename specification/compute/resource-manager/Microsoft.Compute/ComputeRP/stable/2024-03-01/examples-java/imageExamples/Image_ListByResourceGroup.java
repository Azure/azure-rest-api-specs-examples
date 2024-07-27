
/**
 * Samples for Images ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/imageExamples/
     * Image_ListByResourceGroup.json
     */
    /**
     * Sample code: List all virtual machine images in a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllVirtualMachineImagesInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getImages().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
