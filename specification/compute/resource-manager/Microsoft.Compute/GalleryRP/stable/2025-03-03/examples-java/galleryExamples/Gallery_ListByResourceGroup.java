
/**
 * Samples for Galleries ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * Gallery_ListByResourceGroup.json
     */
    /**
     * Sample code: List galleries in a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listGalleriesInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleries().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
