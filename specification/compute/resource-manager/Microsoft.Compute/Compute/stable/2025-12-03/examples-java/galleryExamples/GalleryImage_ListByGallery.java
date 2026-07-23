
/**
 * Samples for GalleryImages ListByGallery.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryImage_ListByGallery.json
     */
    /**
     * Sample code: List gallery images in a gallery.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listGalleryImagesInAGallery(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImages().listByGallery("myResourceGroup", "myGalleryName",
            com.azure.core.util.Context.NONE);
    }
}
