import com.azure.core.util.Context;

/** Samples for SharedGalleries Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/sharedGalleryExamples/SharedGallery_Get.json
     */
    /**
     * Sample code: Get a shared gallery.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASharedGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSharedGalleries()
            .getWithResponse("myLocation", "galleryUniqueName", Context.NONE);
    }
}
