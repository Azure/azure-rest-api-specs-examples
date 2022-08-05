import com.azure.core.util.Context;

/** Samples for GalleryApplications ListByGallery. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/ListGalleryApplicationsInAGallery.json
     */
    /**
     * Sample code: List gallery Applications in a gallery.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listGalleryApplicationsInAGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryApplications()
            .listByGallery("myResourceGroup", "myGalleryName", Context.NONE);
    }
}
