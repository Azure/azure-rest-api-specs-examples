import com.azure.core.util.Context;

/** Samples for GalleryImages ListByGallery. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/ListGalleryImagesInAGallery.json
     */
    /**
     * Sample code: List gallery images in a gallery.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listGalleryImagesInAGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryImages()
            .listByGallery("myResourceGroup", "myGalleryName", Context.NONE);
    }
}
