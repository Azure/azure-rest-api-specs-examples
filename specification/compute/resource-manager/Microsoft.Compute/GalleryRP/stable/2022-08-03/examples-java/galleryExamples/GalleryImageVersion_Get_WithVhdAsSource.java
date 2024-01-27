
import com.azure.resourcemanager.compute.models.ReplicationStatusTypes;

/**
 * Samples for GalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-08-03/examples/galleryExamples/
     * GalleryImageVersion_Get_WithVhdAsSource.json
     */
    /**
     * Sample code: Get a gallery image version with vhd as a source.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryImageVersionWithVhdAsASource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions().getWithResponse("myResourceGroup",
            "myGalleryName", "myGalleryImageName", "1.0.0", null, com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-08-03/examples/galleryExamples/
     * GalleryImageVersion_Get_WithReplicationStatus.json
     */
    /**
     * Sample code: Get a gallery image version with replication status.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAGalleryImageVersionWithReplicationStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions().getWithResponse("myResourceGroup",
            "myGalleryName", "myGalleryImageName", "1.0.0", ReplicationStatusTypes.REPLICATION_STATUS,
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-08-03/examples/galleryExamples/
     * GalleryImageVersion_Get_WithSnapshotsAsSource.json
     */
    /**
     * Sample code: Get a gallery image version with snapshots as a source.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAGalleryImageVersionWithSnapshotsAsASource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions().getWithResponse("myResourceGroup",
            "myGalleryName", "myGalleryImageName", "1.0.0", null, com.azure.core.util.Context.NONE);
    }
}
