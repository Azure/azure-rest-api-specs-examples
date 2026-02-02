
import com.azure.resourcemanager.compute.models.GalleryUpdate;

/**
 * Samples for Galleries Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * Gallery_Update.json
     */
    /**
     * Sample code: Update a simple gallery.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateASimpleGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleries().update("myResourceGroup", "myGalleryName",
            new GalleryUpdate().withDescription("This is the gallery description."), com.azure.core.util.Context.NONE);
    }
}
