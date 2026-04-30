
/**
 * Samples for GalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/GalleryImageVersion_Get_WithSnapshotsAsSource.json
     */
    /**
     * Sample code: Get a gallery image version with snapshots as a source.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAGalleryImageVersionWithSnapshotsAsASource(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().getWithResponse("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0", null, com.azure.core.util.Context.NONE);
    }
}
