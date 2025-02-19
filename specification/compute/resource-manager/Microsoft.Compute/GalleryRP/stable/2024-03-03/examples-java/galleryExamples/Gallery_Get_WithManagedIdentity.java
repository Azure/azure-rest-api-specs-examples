
/**
 * Samples for Galleries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/
     * Gallery_Get_WithManagedIdentity.json
     */
    /**
     * Sample code: Get a gallery with system-assigned and user-assigned managed identities.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryWithSystemAssignedAndUserAssignedManagedIdentities(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleries().getByResourceGroupWithResponse(
            "myResourceGroup", "myGalleryName", null, null, com.azure.core.util.Context.NONE);
    }
}
