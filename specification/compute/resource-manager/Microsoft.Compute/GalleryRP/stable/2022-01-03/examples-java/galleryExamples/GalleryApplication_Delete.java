import com.azure.core.util.Context;

/** Samples for GalleryApplications Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/GalleryApplication_Delete.json
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
