
/**
 * Samples for SharedGalleries Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * sharedGalleryExamples/SharedGallery_Get.json
     */
    /**
     * Sample code: Get a shared gallery.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASharedGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSharedGalleries().getWithResponse("myLocation",
            "galleryUniqueName", com.azure.core.util.Context.NONE);
    }
}
