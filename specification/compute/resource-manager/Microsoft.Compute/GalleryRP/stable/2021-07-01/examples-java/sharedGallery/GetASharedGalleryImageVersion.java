import com.azure.core.util.Context;

/** Samples for SharedGalleryImageVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/sharedGallery/GetASharedGalleryImageVersion.json
     */
    /**
     * Sample code: Get a gallery.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSharedGalleryImageVersions()
            .getWithResponse(
                "myLocation", "galleryUniqueName", "myGalleryImageName", "myGalleryImageVersionName", Context.NONE);
    }
}
