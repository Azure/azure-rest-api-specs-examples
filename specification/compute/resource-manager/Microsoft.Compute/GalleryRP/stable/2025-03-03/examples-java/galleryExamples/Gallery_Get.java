
/**
 * Samples for Galleries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * Gallery_Get.json
     */
    /**
     * Sample code: Get a gallery.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleries().getByResourceGroupWithResponse(
            "myResourceGroup", "myGalleryName", null, null, com.azure.core.util.Context.NONE);
    }
}
