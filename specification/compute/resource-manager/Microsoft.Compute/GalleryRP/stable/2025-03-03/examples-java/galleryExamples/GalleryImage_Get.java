
/**
 * Samples for GalleryImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GalleryImage_Get.json
     */
    /**
     * Sample code: Get a gallery image.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImages().getWithResponse("myResourceGroup",
            "myGalleryName", "myGalleryImageName", com.azure.core.util.Context.NONE);
    }
}
