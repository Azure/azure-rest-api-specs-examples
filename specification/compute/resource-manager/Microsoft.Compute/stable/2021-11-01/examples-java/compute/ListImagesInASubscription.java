import com.azure.core.util.Context;

/** Samples for Images List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/ListImagesInASubscription.json
     */
    /**
     * Sample code: List all virtual machine images in a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllVirtualMachineImagesInASubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getImages().list(Context.NONE);
    }
}
