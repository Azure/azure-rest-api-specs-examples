
/**
 * Samples for GalleryImageVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GalleryImageVersion_Delete.json
     */
    /**
     * Sample code: Delete a gallery image version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAGalleryImageVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions().delete("myResourceGroup",
            "myGalleryName", "myGalleryImageName", "1.0.0", com.azure.core.util.Context.NONE);
    }
}
