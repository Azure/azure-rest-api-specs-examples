import com.azure.core.util.Context;

/** Samples for SharedGalleryImageVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/sharedGalleryExamples/SharedGalleryImageVersions_List.json
     */
    /**
     * Sample code: List shared gallery image versions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSharedGalleryImageVersions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSharedGalleryImageVersions()
            .list("myLocation", "galleryUniqueName", "myGalleryImageName", null, Context.NONE);
    }
}
