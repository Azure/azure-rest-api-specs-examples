
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionUpdate;

/**
 * Samples for GalleryImageVersions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/GalleryImageVersion_Update_RestoreSoftDeleted.json
     */
    /**
     * Sample code: Restore a Soft Deleted Gallery Image Version.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        restoreASoftDeletedGalleryImageVersion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().update("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0", new GalleryImageVersionUpdate()
                .withStorageProfile(new GalleryImageVersionStorageProfile()).withRestore(true),
            com.azure.core.util.Context.NONE);
    }
}
