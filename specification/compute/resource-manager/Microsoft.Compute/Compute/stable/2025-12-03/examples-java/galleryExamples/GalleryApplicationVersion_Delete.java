
/**
 * Samples for GalleryApplicationVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryApplicationVersion_Delete.json
     */
    /**
     * Sample code: Delete a gallery Application Version.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteAGalleryApplicationVersion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryApplicationVersions().delete("myResourceGroup", "myGalleryName",
            "myGalleryApplicationName", "1.0.0", com.azure.core.util.Context.NONE);
    }
}
