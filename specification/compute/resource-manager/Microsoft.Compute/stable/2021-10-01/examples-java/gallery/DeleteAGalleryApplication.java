import com.azure.core.util.Context;

/** Samples for GalleryApplications Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/DeleteAGalleryApplication.json
     */
    /**
     * Sample code: Delete a gallery Application.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAGalleryApplication(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryApplications()
            .delete("myResourceGroup", "myGalleryName", "myGalleryApplicationName", Context.NONE);
    }
}
