
import com.azure.resourcemanager.compute.models.SelectPermissions;

/**
 * Samples for Galleries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/Gallery_Get_WithSelectPermissions.json
     */
    /**
     * Sample code: Get a gallery with select permissions.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAGalleryWithSelectPermissions(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleries().getByResourceGroupWithResponse("myResourceGroup", "myGalleryName",
            SelectPermissions.PERMISSIONS, null, com.azure.core.util.Context.NONE);
    }
}
