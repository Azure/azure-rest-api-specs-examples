
/**
 * Samples for SharedGalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/sharedGalleryExamples/SharedGalleryImageVersion_Get.json
     */
    /**
     * Sample code: Get a shared gallery image version.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getASharedGalleryImageVersion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSharedGalleryImageVersions().getWithResponse("myLocation", "galleryUniqueName",
            "myGalleryImageName", "myGalleryImageVersionName", com.azure.core.util.Context.NONE);
    }
}
