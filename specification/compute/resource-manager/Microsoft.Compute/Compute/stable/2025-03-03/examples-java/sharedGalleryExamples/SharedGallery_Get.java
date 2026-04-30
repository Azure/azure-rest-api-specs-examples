
/**
 * Samples for SharedGalleries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/sharedGalleryExamples/SharedGallery_Get.json
     */
    /**
     * Sample code: Get a shared gallery.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getASharedGallery(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSharedGalleries().getWithResponse("myLocation", "galleryUniqueName",
            com.azure.core.util.Context.NONE);
    }
}
