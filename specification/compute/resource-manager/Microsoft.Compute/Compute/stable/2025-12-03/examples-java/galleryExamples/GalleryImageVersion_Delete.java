
/**
 * Samples for GalleryImageVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryImageVersion_Delete.json
     */
    /**
     * Sample code: Delete a gallery image version.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteAGalleryImageVersion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().delete("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0", com.azure.core.util.Context.NONE);
    }
}
