
/**
 * Samples for SharedGalleryImages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/sharedGalleryExamples/SharedGalleryImages_List.json
     */
    /**
     * Sample code: List shared gallery images.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listSharedGalleryImages(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSharedGalleryImages().list("myLocation", "galleryUniqueName", null,
            com.azure.core.util.Context.NONE);
    }
}
