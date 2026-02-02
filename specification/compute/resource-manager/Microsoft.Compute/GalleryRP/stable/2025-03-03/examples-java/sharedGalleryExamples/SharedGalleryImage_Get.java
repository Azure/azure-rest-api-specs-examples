
/**
 * Samples for SharedGalleryImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * sharedGalleryExamples/SharedGalleryImage_Get.json
     */
    /**
     * Sample code: Get a shared gallery image.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASharedGalleryImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSharedGalleryImages().getWithResponse("myLocation",
            "galleryUniqueName", "myGalleryImageName", com.azure.core.util.Context.NONE);
    }
}
