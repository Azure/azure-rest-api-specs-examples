
/**
 * Samples for GalleryApplicationVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryApplicationVersion_Get.json
     */
    /**
     * Sample code: Get a gallery Application Version.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAGalleryApplicationVersion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryApplicationVersions().getWithResponse("myResourceGroup", "myGalleryName",
            "myGalleryApplicationName", "1.0.0", null, com.azure.core.util.Context.NONE);
    }
}
