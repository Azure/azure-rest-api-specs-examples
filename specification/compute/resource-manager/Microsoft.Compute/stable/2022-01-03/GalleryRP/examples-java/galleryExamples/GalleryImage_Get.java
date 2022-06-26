import com.azure.core.util.Context;

/** Samples for GalleryImages Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/GalleryImage_Get.json
     */
    /**
     * Sample code: Get a gallery image.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryImages()
            .getWithResponse("myResourceGroup", "myGalleryName", "myGalleryImageName", Context.NONE);
    }
}
