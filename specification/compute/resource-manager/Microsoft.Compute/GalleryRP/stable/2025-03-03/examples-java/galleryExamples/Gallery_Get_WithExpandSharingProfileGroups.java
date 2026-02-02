
import com.azure.resourcemanager.compute.models.GalleryExpandParams;

/**
 * Samples for Galleries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * Gallery_Get_WithExpandSharingProfileGroups.json
     */
    /**
     * Sample code: Get a gallery with expand sharingProfile groups.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryWithExpandSharingProfileGroups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleries().getByResourceGroupWithResponse(
            "myResourceGroup", "myGalleryName", null, GalleryExpandParams.SHARING_PROFILE_GROUPS,
            com.azure.core.util.Context.NONE);
    }
}
