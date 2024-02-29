
import com.azure.resourcemanager.compute.models.ReplicationStatusTypes;

/**
 * Samples for GalleryApplicationVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/
     * GalleryApplicationVersion_Get.json
     */
    /**
     * Sample code: Get a gallery Application Version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryApplicationVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryApplicationVersions().getWithResponse(
            "myResourceGroup", "myGalleryName", "myGalleryApplicationName", "1.0.0", null,
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/
     * GalleryApplicationVersion_Get_WithReplicationStatus.json
     */
    /**
     * Sample code: Get a gallery Application Version with replication status.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAGalleryApplicationVersionWithReplicationStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryApplicationVersions().getWithResponse(
            "myResourceGroup", "myGalleryName", "myGalleryApplicationName", "1.0.0",
            ReplicationStatusTypes.REPLICATION_STATUS, com.azure.core.util.Context.NONE);
    }
}
