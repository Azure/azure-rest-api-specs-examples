
/**
 * Samples for Galleries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/CommunityGallery_Get.json
     */
    /**
     * Sample code: Get a community gallery.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getACommunityGallery(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleries().getByResourceGroupWithResponse("myResourceGroup", "myGalleryName", null,
            null, com.azure.core.util.Context.NONE);
    }
}
