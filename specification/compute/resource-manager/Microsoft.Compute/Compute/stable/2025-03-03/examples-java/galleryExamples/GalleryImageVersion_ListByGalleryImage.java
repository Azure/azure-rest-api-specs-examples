
/**
 * Samples for GalleryImageVersions ListByGalleryImage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/GalleryImageVersion_ListByGalleryImage.json
     */
    /**
     * Sample code: List gallery image versions in a gallery image definition.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listGalleryImageVersionsInAGalleryImageDefinition(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().listByGalleryImage("myResourceGroup", "myGalleryName",
            "myGalleryImageName", com.azure.core.util.Context.NONE);
    }
}
