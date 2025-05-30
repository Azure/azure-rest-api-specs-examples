
/**
 * Samples for GalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/
     * GalleryImageVersion_Get_WithVhdAsSource.json
     */
    /**
     * Sample code: Get a gallery image version with vhd as a source.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryImageVersionWithVhdAsASource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions().getWithResponse("myResourceGroup",
            "myGalleryName", "myGalleryImageName", "1.0.0", null, com.azure.core.util.Context.NONE);
    }
}
