
/**
 * Samples for GalleryScriptVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryScriptExamples/GalleryScriptVersion_Get_WithReplicationStatus.json
     */
    /**
     * Sample code: Get a gallery Script Version with replication status.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAGalleryScriptVersionWithReplicationStatus(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryScriptVersions().getWithResponse("myResourceGroupName", "myGalleryName",
            "myGalleryScriptName", "1.0.0", com.azure.core.util.Context.NONE);
    }
}
