import com.azure.core.util.Context;

/** Samples for GalleryApplications Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/GetAGalleryApplication.json
     */
    /**
     * Sample code: Get a gallery Application.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryApplication(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryApplications()
            .getWithResponse("myResourceGroup", "myGalleryName", "myGalleryApplicationName", Context.NONE);
    }
}
