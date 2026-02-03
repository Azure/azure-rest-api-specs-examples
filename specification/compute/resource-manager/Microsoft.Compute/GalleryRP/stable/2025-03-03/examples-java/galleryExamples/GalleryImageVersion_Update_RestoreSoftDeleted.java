
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionUpdate;

/**
 * Samples for GalleryImageVersions Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GalleryImageVersion_Update_RestoreSoftDeleted.json
     */
    /**
     * Sample code: Restore a Soft Deleted Gallery Image Version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restoreASoftDeletedGalleryImageVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions().update("myResourceGroup",
            "myGalleryName", "myGalleryImageName", "1.0.0", new GalleryImageVersionUpdate()
                .withStorageProfile(new GalleryImageVersionStorageProfile()).withRestore(true),
            com.azure.core.util.Context.NONE);
    }
}
