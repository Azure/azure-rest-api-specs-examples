
/**
 * Samples for GalleryImages Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GalleryImage_Delete.json
     */
    /**
     * Sample code: Delete a gallery image.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAGalleryImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImages().delete("myResourceGroup", "myGalleryName",
            "myGalleryImageName", com.azure.core.util.Context.NONE);
    }
}
