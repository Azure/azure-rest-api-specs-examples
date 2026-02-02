
/**
 * Samples for SoftDeletedResource ListByArtifactName.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GallerySoftDeletedResource_ListByArtifactName.json
     */
    /**
     * Sample code: List soft-deleted resources of an artifact in the gallery.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listSoftDeletedResourcesOfAnArtifactInTheGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSoftDeletedResources().listByArtifactName(
            "myResourceGroup", "myGalleryName", "images", "myGalleryImageName", com.azure.core.util.Context.NONE);
    }
}
