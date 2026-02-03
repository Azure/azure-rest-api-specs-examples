
/**
 * Samples for GalleryApplications ListByGallery.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GalleryApplication_ListByGallery.json
     */
    /**
     * Sample code: List gallery Applications in a gallery.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listGalleryApplicationsInAGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryApplications().listByGallery("myResourceGroup",
            "myGalleryName", com.azure.core.util.Context.NONE);
    }
}
