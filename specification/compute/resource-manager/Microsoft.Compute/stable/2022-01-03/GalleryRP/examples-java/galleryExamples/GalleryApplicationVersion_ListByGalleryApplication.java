import com.azure.core.util.Context;

/** Samples for GalleryApplicationVersions ListByGalleryApplication. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/GalleryApplicationVersion_ListByGalleryApplication.json
     */
    /**
     * Sample code: List gallery Application Versions in a gallery Application Definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listGalleryApplicationVersionsInAGalleryApplicationDefinition(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryApplicationVersions()
            .listByGalleryApplication("myResourceGroup", "myGalleryName", "myGalleryApplicationName", Context.NONE);
    }
}
