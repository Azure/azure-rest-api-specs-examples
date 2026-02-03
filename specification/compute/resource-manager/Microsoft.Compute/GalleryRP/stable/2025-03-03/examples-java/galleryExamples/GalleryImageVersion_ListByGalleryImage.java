
/**
 * Samples for GalleryImageVersions ListByGalleryImage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GalleryImageVersion_ListByGalleryImage.json
     */
    /**
     * Sample code: List gallery image versions in a gallery image definition.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listGalleryImageVersionsInAGalleryImageDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions().listByGalleryImage(
            "myResourceGroup", "myGalleryName", "myGalleryImageName", com.azure.core.util.Context.NONE);
    }
}
