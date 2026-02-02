
/**
 * Samples for GalleryScriptVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * galleryScriptExamples/GalleryScriptVersion_Get_WithReplicationStatus.json
     */
    /**
     * Sample code: Get a gallery Script Version with replication status.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAGalleryScriptVersionWithReplicationStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryScriptVersions().getWithResponse(
            "myResourceGroupName", "myGalleryName", "myGalleryScriptName", "1.0.0", com.azure.core.util.Context.NONE);
    }
}
