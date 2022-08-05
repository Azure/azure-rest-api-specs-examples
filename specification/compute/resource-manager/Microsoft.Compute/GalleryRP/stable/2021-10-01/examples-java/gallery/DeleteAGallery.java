import com.azure.core.util.Context;

/** Samples for Galleries Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/DeleteAGallery.json
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
