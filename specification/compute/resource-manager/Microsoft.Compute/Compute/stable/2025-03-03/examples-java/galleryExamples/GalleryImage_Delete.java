
/**
 * Samples for GalleryImages Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/GalleryImage_Delete.json
     */
    /**
     * Sample code: Delete a gallery image.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteAGalleryImage(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImages().delete("myResourceGroup", "myGalleryName", "myGalleryImageName",
            com.azure.core.util.Context.NONE);
    }
}
