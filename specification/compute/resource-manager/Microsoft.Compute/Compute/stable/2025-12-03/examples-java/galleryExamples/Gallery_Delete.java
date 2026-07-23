
/**
 * Samples for Galleries Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/Gallery_Delete.json
     */
    /**
     * Sample code: Delete a gallery.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteAGallery(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleries().delete("myResourceGroup", "myGalleryName",
            com.azure.core.util.Context.NONE);
    }
}
