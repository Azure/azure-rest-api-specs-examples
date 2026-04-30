
/**
 * Samples for GalleryApplications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/GalleryApplication_Get.json
     */
    /**
     * Sample code: Get a gallery Application.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAGalleryApplication(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryApplications().getWithResponse("myResourceGroup", "myGalleryName",
            "myGalleryApplicationName", com.azure.core.util.Context.NONE);
    }
}
