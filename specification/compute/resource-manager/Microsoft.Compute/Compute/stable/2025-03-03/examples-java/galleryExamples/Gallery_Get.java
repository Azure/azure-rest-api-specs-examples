
/**
 * Samples for Galleries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/Gallery_Get.json
     */
    /**
     * Sample code: Get a gallery.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAGallery(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleries().getByResourceGroupWithResponse("myResourceGroup", "myGalleryName", null,
            null, com.azure.core.util.Context.NONE);
    }
}
