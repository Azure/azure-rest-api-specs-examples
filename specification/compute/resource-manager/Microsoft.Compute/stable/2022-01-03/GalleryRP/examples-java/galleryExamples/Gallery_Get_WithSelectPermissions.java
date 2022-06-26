import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.SelectPermissions;

/** Samples for Galleries GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/Gallery_Get_WithSelectPermissions.json
     */
    /**
     * Sample code: Get a gallery with select permissions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryWithSelectPermissions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleries()
            .getByResourceGroupWithResponse(
                "myResourceGroup", "myGalleryName", SelectPermissions.PERMISSIONS, null, Context.NONE);
    }
}
