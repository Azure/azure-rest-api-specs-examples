
import com.azure.resourcemanager.compute.models.ReplicationStatusTypes;

/**
 * Samples for GalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-03-03/galleryExamples/GalleryImageVersion_Get_WithValidationProfileAndReplicationStatus.json
     */
    /**
     * Sample code: Get a gallery image version with validation profile and replication status.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAGalleryImageVersionWithValidationProfileAndReplicationStatus(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().getWithResponse("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0", ReplicationStatusTypes.fromString("ValidationProfile,ReplicationStatus"),
            com.azure.core.util.Context.NONE);
    }
}
