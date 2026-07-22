
/**
 * Samples for GalleryImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryImage_Get.json
     */
    /**
     * Sample code: Get a gallery image.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAGalleryImage(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImages().getWithResponse("myResourceGroup", "myGalleryName",
            "myGalleryImageName", com.azure.core.util.Context.NONE);
    }
}
