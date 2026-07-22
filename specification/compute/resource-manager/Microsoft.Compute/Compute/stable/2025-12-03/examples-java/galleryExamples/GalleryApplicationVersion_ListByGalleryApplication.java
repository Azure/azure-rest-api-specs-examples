
/**
 * Samples for GalleryApplicationVersions ListByGalleryApplication.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryApplicationVersion_ListByGalleryApplication.json
     */
    /**
     * Sample code: List gallery Application Versions in a gallery Application Definition.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listGalleryApplicationVersionsInAGalleryApplicationDefinition(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryApplicationVersions().listByGalleryApplication("myResourceGroup",
            "myGalleryName", "myGalleryApplicationName", com.azure.core.util.Context.NONE);
    }
}
