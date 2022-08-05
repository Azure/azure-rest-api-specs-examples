import com.azure.core.util.Context;

/** Samples for GalleryImages Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/DeleteAGalleryImage.json
     */
    /**
     * Sample code: Delete a gallery image.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAGalleryImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryImages()
            .delete("myResourceGroup", "myGalleryName", "myGalleryImageName", Context.NONE);
    }
}
