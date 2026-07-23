
/**
 * Samples for SoftDeletedResource ListByArtifactName.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/GallerySoftDeletedResource_ListByArtifactName.json
     */
    /**
     * Sample code: List soft-deleted resources of an artifact in the gallery.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listSoftDeletedResourcesOfAnArtifactInTheGallery(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSoftDeletedResources().listByArtifactName("myResourceGroup", "myGalleryName",
            "images", "myGalleryImageName", com.azure.core.util.Context.NONE);
    }
}
