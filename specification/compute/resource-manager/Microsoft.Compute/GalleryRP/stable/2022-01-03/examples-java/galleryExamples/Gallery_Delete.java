import com.azure.core.util.Context;

/** Samples for Galleries Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/Gallery_Delete.json
     */
    /**
     * Sample code: Delete a gallery.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleries()
            .delete("myResourceGroup", "myGalleryName", Context.NONE);
    }
}
