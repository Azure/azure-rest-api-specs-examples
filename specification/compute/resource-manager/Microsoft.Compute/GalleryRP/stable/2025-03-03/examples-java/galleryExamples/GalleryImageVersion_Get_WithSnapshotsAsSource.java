
/**
 * Samples for GalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
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
