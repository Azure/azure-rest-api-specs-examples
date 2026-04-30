
/**
 * Samples for GalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/GalleryImageVersion_Get.json
     */
    /**
     * Sample code: Get a gallery image version.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAGalleryImageVersion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().getWithResponse("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0", null, com.azure.core.util.Context.NONE);
    }
}
