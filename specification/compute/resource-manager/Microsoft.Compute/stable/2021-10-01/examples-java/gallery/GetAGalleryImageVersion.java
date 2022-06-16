import com.azure.core.util.Context;

/** Samples for GalleryImageVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/GetAGalleryImageVersion.json
     */
    /**
     * Sample code: Get a gallery image version.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryImageVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryImageVersions()
            .getWithResponse("myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0", null, Context.NONE);
    }
}
