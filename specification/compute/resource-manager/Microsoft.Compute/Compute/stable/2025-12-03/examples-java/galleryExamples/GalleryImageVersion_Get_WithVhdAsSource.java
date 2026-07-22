
/**
 * Samples for GalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryImageVersion_Get_WithVhdAsSource.json
     */
    /**
     * Sample code: Get a gallery image version with vhd as a source.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAGalleryImageVersionWithVhdAsASource(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().getWithResponse("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0", null, com.azure.core.util.Context.NONE);
    }
}
