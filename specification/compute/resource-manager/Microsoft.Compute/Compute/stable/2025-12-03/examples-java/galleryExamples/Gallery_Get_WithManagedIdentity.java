
/**
 * Samples for Galleries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/Gallery_Get_WithManagedIdentity.json
     */
    /**
     * Sample code: Get a gallery with system-assigned and user-assigned managed identities.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAGalleryWithSystemAssignedAndUserAssignedManagedIdentities(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleries().getByResourceGroupWithResponse("myResourceGroup", "myGalleryName", null,
            null, com.azure.core.util.Context.NONE);
    }
}
