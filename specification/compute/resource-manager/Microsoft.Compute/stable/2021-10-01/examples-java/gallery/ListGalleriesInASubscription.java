import com.azure.core.util.Context;

/** Samples for Galleries List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/ListGalleriesInASubscription.json
     */
    /**
     * Sample code: List galleries in a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listGalleriesInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleries().list(Context.NONE);
    }
}
