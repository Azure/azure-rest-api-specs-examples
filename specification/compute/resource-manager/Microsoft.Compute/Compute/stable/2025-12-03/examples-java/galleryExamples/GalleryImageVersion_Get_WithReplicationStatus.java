
import com.azure.resourcemanager.compute.models.ReplicationStatusTypes;

/**
 * Samples for GalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryImageVersion_Get_WithReplicationStatus.json
     */
    /**
     * Sample code: Get a gallery image version with replication status.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAGalleryImageVersionWithReplicationStatus(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().getWithResponse("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0", ReplicationStatusTypes.REPLICATION_STATUS, com.azure.core.util.Context.NONE);
    }
}
