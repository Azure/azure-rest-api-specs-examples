import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.SelectPermissions;

/** Samples for Galleries GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/GetAGalleryWithSelectPermissions.json
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
                "myResourceGroup", "myGalleryName", SelectPermissions.PERMISSIONS, Context.NONE);
    }
}
