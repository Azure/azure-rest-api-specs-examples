import com.azure.core.util.Context;

/** Samples for SharedGalleries List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/sharedGalleryExamples/SharedGallery_List.json
     */
    /**
     * Sample code: List shared galleries.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSharedGalleries(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSharedGalleries().list("myLocation", null, Context.NONE);
    }
}
