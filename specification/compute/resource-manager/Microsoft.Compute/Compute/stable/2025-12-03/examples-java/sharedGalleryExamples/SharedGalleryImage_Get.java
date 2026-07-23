
/**
 * Samples for SharedGalleryImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/sharedGalleryExamples/SharedGalleryImage_Get.json
     */
    /**
     * Sample code: Get a shared gallery image.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getASharedGalleryImage(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSharedGalleryImages().getWithResponse("myLocation", "galleryUniqueName",
            "myGalleryImageName", com.azure.core.util.Context.NONE);
    }
}
