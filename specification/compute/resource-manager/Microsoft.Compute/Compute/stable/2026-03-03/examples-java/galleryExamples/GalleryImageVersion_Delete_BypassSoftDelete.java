
/**
 * Samples for GalleryImageVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-03/galleryExamples/GalleryImageVersion_Delete_BypassSoftDelete.json
     */
    /**
     * Sample code: Permanently delete a gallery image version by bypassing soft delete.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void permanentlyDeleteAGalleryImageVersionByBypassingSoftDelete(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().delete("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0", true, com.azure.core.util.Context.NONE);
    }
}
