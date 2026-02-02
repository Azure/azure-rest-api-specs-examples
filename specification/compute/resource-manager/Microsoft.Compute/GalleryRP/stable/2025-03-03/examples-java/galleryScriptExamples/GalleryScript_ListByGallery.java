
/**
 * Samples for GalleryScripts ListByGallery.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * galleryScriptExamples/GalleryScript_ListByGallery.json
     */
    /**
     * Sample code: List gallery Scripts in a gallery.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listGalleryScriptsInAGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryScripts().listByGallery("myResourceGroup",
            "myGalleryName", com.azure.core.util.Context.NONE);
    }
}
