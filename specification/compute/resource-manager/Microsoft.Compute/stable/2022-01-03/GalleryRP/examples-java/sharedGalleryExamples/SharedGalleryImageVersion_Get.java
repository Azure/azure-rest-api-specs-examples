import com.azure.core.util.Context;

/** Samples for SharedGalleryImageVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/sharedGalleryExamples/SharedGalleryImageVersion_Get.json
     */
    /**
     * Sample code: Get a shared gallery image version.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASharedGalleryImageVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSharedGalleryImageVersions()
            .getWithResponse(
                "myLocation", "galleryUniqueName", "myGalleryImageName", "myGalleryImageVersionName", Context.NONE);
    }
}
