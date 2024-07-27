
/**
 * Samples for Images List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/imageExamples/
     * Image_ListBySubscription.json
     */
    /**
     * Sample code: List all virtual machine images in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllVirtualMachineImagesInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getImages().list(com.azure.core.util.Context.NONE);
    }
}
