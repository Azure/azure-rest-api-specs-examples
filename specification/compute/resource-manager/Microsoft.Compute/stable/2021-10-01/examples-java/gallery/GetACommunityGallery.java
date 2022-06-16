import com.azure.core.util.Context;

/** Samples for Galleries GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/GetACommunityGallery.json
     */
    /**
     * Sample code: Get a community gallery.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getACommunityGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleries()
            .getByResourceGroupWithResponse("myResourceGroup", "myGalleryName", null, null, Context.NONE);
    }
}
