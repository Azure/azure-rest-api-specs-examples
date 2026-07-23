
/**
 * Samples for GalleryApplications ListByGallery.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryApplication_ListByGallery.json
     */
    /**
     * Sample code: List gallery Applications in a gallery.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listGalleryApplicationsInAGallery(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryApplications().listByGallery("myResourceGroup", "myGalleryName",
            com.azure.core.util.Context.NONE);
    }
}
