
import com.azure.resourcemanager.compute.models.GalleryUpdate;

/**
 * Samples for Galleries Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/Gallery_Update.json
     */
    /**
     * Sample code: Update a simple gallery.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void updateASimpleGallery(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleries().update("myResourceGroup", "myGalleryName",
            new GalleryUpdate().withDescription("This is the gallery description."), com.azure.core.util.Context.NONE);
    }
}
