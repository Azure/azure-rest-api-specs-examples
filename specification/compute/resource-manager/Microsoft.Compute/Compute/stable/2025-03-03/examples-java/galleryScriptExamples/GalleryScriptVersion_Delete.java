
/**
 * Samples for GalleryScriptVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryScriptExamples/GalleryScriptVersion_Delete.json
     */
    /**
     * Sample code: Delete a gallery Script Version.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteAGalleryScriptVersion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryScriptVersions().delete("myResourceGroupName", "myGalleryName",
            "myGalleryScriptName", "1.0.0", com.azure.core.util.Context.NONE);
    }
}
