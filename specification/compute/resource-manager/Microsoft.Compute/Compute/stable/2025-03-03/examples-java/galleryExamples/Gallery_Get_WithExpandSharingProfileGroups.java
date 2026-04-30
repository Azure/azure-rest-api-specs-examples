
import com.azure.resourcemanager.compute.models.GalleryExpandParams;

/**
 * Samples for Galleries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/Gallery_Get_WithExpandSharingProfileGroups.json
     */
    /**
     * Sample code: Get a gallery with expand sharingProfile groups.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAGalleryWithExpandSharingProfileGroups(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleries().getByResourceGroupWithResponse("myResourceGroup", "myGalleryName", null,
            GalleryExpandParams.SHARING_PROFILE_GROUPS, com.azure.core.util.Context.NONE);
    }
}
