import com.azure.core.util.Context;

/** Samples for GalleryApplicationVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/DeleteAGalleryApplicationVersion.json
     */
    /**
     * Sample code: Delete a gallery Application Version.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAGalleryApplicationVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryApplicationVersions()
            .delete("myResourceGroup", "myGalleryName", "myGalleryApplicationName", "1.0.0", Context.NONE);
    }
}
