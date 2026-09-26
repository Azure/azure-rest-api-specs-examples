
/**
 * Samples for GalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-03/galleryExamples/GalleryImageVersion_Get_WithImageMetadataProfiles.json
     */
    /**
     * Sample code: Get a gallery image version with image metadata profiles.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAGalleryImageVersionWithImageMetadataProfiles(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().getWithResponse("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0", null, com.azure.core.util.Context.NONE);
    }
}
