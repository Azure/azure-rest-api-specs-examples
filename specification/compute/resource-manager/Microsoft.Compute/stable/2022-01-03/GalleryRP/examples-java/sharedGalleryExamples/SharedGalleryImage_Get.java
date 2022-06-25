import com.azure.core.util.Context;

/** Samples for SharedGalleryImages Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/sharedGalleryExamples/SharedGalleryImage_Get.json
     */
    /**
     * Sample code: Get a shared gallery image.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASharedGalleryImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSharedGalleryImages()
            .getWithResponse("myLocation", "galleryUniqueName", "myGalleryImageName", Context.NONE);
    }
}
