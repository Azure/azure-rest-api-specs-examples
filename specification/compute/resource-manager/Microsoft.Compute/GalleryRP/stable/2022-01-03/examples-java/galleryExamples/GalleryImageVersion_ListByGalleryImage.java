import com.azure.core.util.Context;

/** Samples for GalleryImageVersions ListByGalleryImage. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/GalleryImageVersion_ListByGalleryImage.json
     */
    /**
     * Sample code: List gallery image versions in a gallery image definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listGalleryImageVersionsInAGalleryImageDefinition(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryImageVersions()
            .listByGalleryImage("myResourceGroup", "myGalleryName", "myGalleryImageName", Context.NONE);
    }
}
