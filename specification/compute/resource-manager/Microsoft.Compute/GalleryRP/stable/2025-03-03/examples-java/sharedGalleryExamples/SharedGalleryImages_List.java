
/**
 * Samples for SharedGalleryImages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * sharedGalleryExamples/SharedGalleryImages_List.json
     */
    /**
     * Sample code: List shared gallery images.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSharedGalleryImages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSharedGalleryImages().list("myLocation",
            "galleryUniqueName", null, com.azure.core.util.Context.NONE);
    }
}
