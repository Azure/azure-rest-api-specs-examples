
/**
 * Samples for Galleries Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-08-03/examples/galleryExamples/
     * Gallery_Delete.json
     */
    /**
     * Sample code: Delete a gallery.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleries().delete("myResourceGroup", "myGalleryName",
            com.azure.core.util.Context.NONE);
    }
}
