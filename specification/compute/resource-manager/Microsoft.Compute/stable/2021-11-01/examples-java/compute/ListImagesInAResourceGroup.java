import com.azure.core.util.Context;

/** Samples for Images ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/ListImagesInAResourceGroup.json
     */
    /**
     * Sample code: List all virtual machine images in a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllVirtualMachineImagesInAResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getImages()
            .listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
