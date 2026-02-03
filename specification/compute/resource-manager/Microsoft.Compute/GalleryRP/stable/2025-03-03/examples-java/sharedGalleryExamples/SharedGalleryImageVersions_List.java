
/**
 * Samples for SharedGalleryImageVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * sharedGalleryExamples/SharedGalleryImageVersions_List.json
     */
    /**
     * Sample code: List shared gallery image versions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSharedGalleryImageVersions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSharedGalleryImageVersions().list("myLocation",
            "galleryUniqueName", "myGalleryImageName", null, com.azure.core.util.Context.NONE);
    }
}
